package frontier

import (
	"testing"

	"github.com/digitalbits/go/services/frontier/internal/db2/history"
	"github.com/digitalbits/go/services/frontier/internal/ingest"
	"github.com/digitalbits/go/xdr"
)

func TestAccountActions_InvalidID(t *testing.T) {
	ht := StartHTTPTestWithoutScenario(t)
	defer ht.Finish()

	// Makes StateMiddleware happy
	q := history.Q{ht.FrontierSession()}
	err := q.UpdateLastLedgerExpIngest(100)
	ht.Assert.NoError(err)
	err = q.UpdateExpIngestVersion(ingest.CurrentVersion)
	ht.Assert.NoError(err)
	_, err = q.InsertLedger(xdr.LedgerHeaderHistoryEntry{
		Header: xdr.LedgerHeader{
			LedgerSeq: 100,
		},
	}, 0, 0, 0, 0, 0)
	ht.Assert.NoError(err)

	// existing account
	w := ht.Get(
		"/accounts/=cr%FF%98%CB%F3%AF%E72%D85%FE%28%15y%8Fz%C4Ng%CE%98h%02%2A:%B6%FF%B9%CF%92%88O%91%10d&S%7C%9Bi%D4%CFI%28%CFo",
	)
	ht.Assert.Equal(400, w.Code)
}
