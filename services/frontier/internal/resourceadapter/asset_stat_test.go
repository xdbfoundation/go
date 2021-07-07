package resourceadapter

import (
	"context"
	"testing"

	"github.com/xdbfoundation/go/protocols/frontier"
	protocol "github.com/xdbfoundation/go/protocols/frontier"
	"github.com/xdbfoundation/go/services/frontier/internal/db2/history"
	"github.com/xdbfoundation/go/xdr"
	"github.com/stretchr/testify/assert"
)

func TestPopulateExpAssetStat(t *testing.T) {
	row := history.ExpAssetStat{
		AssetType:   xdr.AssetTypeAssetTypeCreditAlphanum4,
		AssetCode:   "XIM",
		AssetIssuer: "GBZ35ZJRIKJGYH5PBKLKOZ5L6EXCNTO7BKIL7DAVVDFQ2ODJEEHHJXIM",
		Amount:      "100000000000000000000", // 10T
		NumAccounts: 429,
	}
	issuer := history.AccountEntry{
		AccountID:  "GBZ35ZJRIKJGYH5PBKLKOZ5L6EXCNTO7BKIL7DAVVDFQ2ODJEEHHJXIM",
		Flags:      0,
		HomeDomain: "xim.com",
	}

	var res protocol.AssetStat
	err := PopulateAssetStat(context.Background(), &res, row, issuer)
	assert.NoError(t, err)

	assert.Equal(t, "credit_alphanum4", res.Type)
	assert.Equal(t, "XIM", res.Code)
	assert.Equal(t, "GBZ35ZJRIKJGYH5PBKLKOZ5L6EXCNTO7BKIL7DAVVDFQ2ODJEEHHJXIM", res.Issuer)
	assert.Equal(t, "10000000000000.0000000", res.Amount)
	assert.Equal(t, int32(429), res.NumAccounts)
	assert.Equal(t, frontier.AccountFlags{}, res.Flags)
	assert.Equal(t, "https://xim.com/.well-known/digitalbits.toml", res.Links.Toml.Href)
	assert.Equal(t, row.PagingToken(), res.PagingToken())

	issuer.HomeDomain = ""
	issuer.Flags = uint32(xdr.AccountFlagsAuthRequiredFlag) |
		uint32(xdr.AccountFlagsAuthImmutableFlag)

	err = PopulateAssetStat(context.Background(), &res, row, issuer)
	assert.NoError(t, err)

	assert.Equal(t, "credit_alphanum4", res.Type)
	assert.Equal(t, "XIM", res.Code)
	assert.Equal(t, "GBZ35ZJRIKJGYH5PBKLKOZ5L6EXCNTO7BKIL7DAVVDFQ2ODJEEHHJXIM", res.Issuer)
	assert.Equal(t, "10000000000000.0000000", res.Amount)
	assert.Equal(t, int32(429), res.NumAccounts)
	assert.Equal(
		t,
		frontier.AccountFlags{
			AuthRequired:  true,
			AuthImmutable: true,
		},
		res.Flags,
	)
	assert.Equal(t, "", res.Links.Toml.Href)
	assert.Equal(t, row.PagingToken(), res.PagingToken())
}
