// Package meta provides helpers for processing the metadata that is produced by
// digitalbits-core while processing transactions.
package meta

import "github.com/xdbfoundation/go/xdr"

// Bundle represents all of the metadata emitted from the application of a single
// digitalbits transaction; Both fee meta and result meta is included.
type Bundle struct {
	FeeMeta         xdr.LedgerEntryChanges
	TransactionMeta xdr.TransactionMeta
}
