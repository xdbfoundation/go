// Package db provides helpers to connect to test databases.  It has no
// internal dependencies on frontier and so should be able to be imported by
// any frontier package.
package db

import (
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	// pq enables postgres support
	_ "github.com/lib/pq"
	db "github.com/xdbfoundation/go/support/db/dbtest"
)

var (
	coreDB     *sqlx.DB
	coreUrl    *string
	frontierDB  *sqlx.DB
	frontierUrl *string
)

// Frontier returns a connection to the frontier test database
func Frontier(t *testing.T) *sqlx.DB {
	if frontierDB != nil {
		return frontierDB
	}
	postgres := db.Postgres(t)
	frontierUrl = &postgres.DSN
	frontierDB = postgres.Open()

	return frontierDB
}

// FrontierURL returns the database connection the url any test
// use when connecting to the history/frontier database
func FrontierURL() string {
	if frontierUrl == nil {
		log.Panic(fmt.Errorf("Frontier not initialized"))
	}
	return *frontierUrl
}

// DigitalBitsCore returns a connection to the digitalbits core test database
func DigitalBitsCore(t *testing.T) *sqlx.DB {
	if coreDB != nil {
		return coreDB
	}
	postgres := db.Postgres(t)
	coreUrl = &postgres.DSN
	coreDB = postgres.Open()
	return coreDB
}

// DigitalBitsCoreURL returns the database connection the url any test
// use when connecting to the digitalbits-core database
func DigitalBitsCoreURL() string {
	if coreUrl == nil {
		log.Panic(fmt.Errorf("DigitalBitsCore not initialized"))
	}
	return *coreUrl
}
