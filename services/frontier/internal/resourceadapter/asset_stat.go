package resourceadapter

import (
	"context"
	"strings"

	"github.com/digitalbits/go/amount"
	protocol "github.com/digitalbits/go/protocols/frontier"
	"github.com/digitalbits/go/services/frontier/internal/db2/history"
	"github.com/digitalbits/go/support/errors"
	"github.com/digitalbits/go/support/render/hal"
	"github.com/digitalbits/go/xdr"
)

// PopulateAssetStat populates an AssetStat using asset stats and account entries
// generated from the ingestion system.
func PopulateAssetStat(
	ctx context.Context,
	res *protocol.AssetStat,
	row history.ExpAssetStat,
	issuer history.AccountEntry,
) (err error) {
	res.Asset.Type = xdr.AssetTypeToString[row.AssetType]
	res.Asset.Code = row.AssetCode
	res.Asset.Issuer = row.AssetIssuer
	res.Amount, err = amount.IntStringToAmount(row.Amount)
	if err != nil {
		return errors.Wrap(err, "Invalid amount in PopulateAssetStat")
	}
	res.NumAccounts = row.NumAccounts
	flags := int8(issuer.Flags)
	res.Flags = protocol.AccountFlags{
		(flags & int8(xdr.AccountFlagsAuthRequiredFlag)) != 0,
		(flags & int8(xdr.AccountFlagsAuthRevocableFlag)) != 0,
		(flags & int8(xdr.AccountFlagsAuthImmutableFlag)) != 0,
	}
	res.PT = row.PagingToken()

	trimmed := strings.TrimSpace(issuer.HomeDomain)
	var toml string
	if trimmed != "" {
		toml = "https://" + issuer.HomeDomain + "/.well-known/digitalbits.toml"
	}
	res.Links.Toml = hal.NewLink(toml)
	return
}
