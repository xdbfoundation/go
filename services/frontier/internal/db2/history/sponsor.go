package history

import (
	"github.com/guregu/null"
	"github.com/digitalbits/go/xdr"
)

func ledgerEntrySponsorToNullString(entry xdr.LedgerEntry) null.String {
	sponsoringID := entry.SponsoringID()

	var sponsor null.String
	if sponsoringID != nil {
		sponsor.SetValid((*sponsoringID).Address())
	}

	return sponsor
}
