package test

import (
	"io"
	"testing"

	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/digitalbits/go/services/frontier/internal/db2/schema"
	"github.com/digitalbits/go/services/frontier/internal/ledger"
	"github.com/digitalbits/go/services/frontier/internal/operationfeestats"
	"github.com/digitalbits/go/services/frontier/internal/test/scenarios"
	"github.com/digitalbits/go/support/db"
	"github.com/digitalbits/go/support/render/hal"
)

// CoreSession returns a db.Session instance pointing at the digitalbits core test database
func (t *T) CoreSession() *db.Session {
	return &db.Session{
		DB:  t.CoreDB,
		Ctx: t.Ctx,
	}
}

// Finish finishes the test, logging any accumulated frontier logs to the logs
// output
func (t *T) Finish() {
	RestoreLogger()
	operationfeestats.ResetState()

	if t.LogBuffer.Len() > 0 {
		t.T.Log("\n" + t.LogBuffer.String())
	}
}

// FrontierSession returns a db.Session instance pointing at the frontier test
// database
func (t *T) FrontierSession() *db.Session {
	return &db.Session{
		DB:  t.FrontierDB,
		Ctx: t.Ctx,
	}
}

func (t *T) loadScenario(scenarioName string, includeFrontier bool) {
	digitalbitsCorePath := scenarioName + "-core.sql"

	scenarios.Load(DigitalBitsCoreDatabaseURL(), digitalbitsCorePath)

	if includeFrontier {
		frontierPath := scenarioName + "-frontier.sql"
		scenarios.Load(DatabaseURL(), frontierPath)
	}
}

// Scenario loads the named sql scenario into the database
func (t *T) Scenario(name string) ledger.Status {
	clearFrontierDB(t.T, t.FrontierDB)
	t.loadScenario(name, true)
	return t.LoadLedgerStatus()
}

// ScenarioWithoutFrontier loads the named sql scenario into the database
func (t *T) ScenarioWithoutFrontier(name string) ledger.Status {
	t.loadScenario(name, false)
	ResetFrontierDB(t.T, t.FrontierDB)
	return t.LoadLedgerStatus()
}

// ResetFrontierDB sets up a new frontier database with empty tables
func ResetFrontierDB(t *testing.T, db *sqlx.DB) {
	clearFrontierDB(t, db)
	_, err := schema.Migrate(db.DB, schema.MigrateUp, 0)
	if err != nil {
		t.Fatalf("could not run migrations up on test db: %v", err)
	}
}

func clearFrontierDB(t *testing.T, db *sqlx.DB) {
	_, err := schema.Migrate(db.DB, schema.MigrateDown, 0)
	if err != nil {
		t.Fatalf("could not run migrations down on test db: %v", err)
	}
}

// UnmarshalPage populates dest with the records contained in the json-encoded page in r
func (t *T) UnmarshalPage(r io.Reader, dest interface{}) hal.Links {
	var env struct {
		Embedded struct {
			Records json.RawMessage `json:"records"`
		} `json:"_embedded"`
		Links struct {
			Self hal.Link `json:"self"`
			Next hal.Link `json:"next"`
			Prev hal.Link `json:"prev"`
		} `json:"_links"`
	}

	err := json.NewDecoder(r).Decode(&env)
	t.Require.NoError(err, "failed to decode page")

	err = json.Unmarshal(env.Embedded.Records, dest)
	t.Require.NoError(err, "failed to decode records")

	return env.Links
}

// UnmarshalNext extracts and returns the next link
func (t *T) UnmarshalNext(r io.Reader) string {
	var env struct {
		Links struct {
			Next struct {
				Href string `json:"href"`
			} `json:"next"`
		} `json:"_links"`
	}

	err := json.NewDecoder(r).Decode(&env)
	t.Require.NoError(err, "failed to decode page")
	return env.Links.Next.Href
}

// UnmarshalExtras extracts and returns extras content
func (t *T) UnmarshalExtras(r io.Reader) map[string]string {
	var resp struct {
		Extras map[string]string `json:"extras"`
	}

	err := json.NewDecoder(r).Decode(&resp)
	t.Require.NoError(err, "failed to decode page")

	return resp.Extras
}

// LoadLedgerStatus loads ledger state from the core db(or panicing on failure).
func (t *T) LoadLedgerStatus() ledger.Status {
	var next ledger.Status

	err := t.CoreSession().GetRaw(&next, `
		SELECT
			COALESCE(MAX(ledgerseq), 0) as core_latest
		FROM ledgerheaders
	`)

	if err != nil {
		panic(err)
	}

	err = t.FrontierSession().GetRaw(&next, `
			SELECT
				COALESCE(MIN(sequence), 0) as history_elder,
				COALESCE(MAX(sequence), 0) as history_latest
			FROM history_ledgers
		`)

	if err != nil {
		panic(err)
	}

	return next
}
