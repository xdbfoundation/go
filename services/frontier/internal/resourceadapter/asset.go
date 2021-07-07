package resourceadapter

import (
	"context"

	protocol "github.com/xdbfoundation/go/protocols/frontier"
	"github.com/xdbfoundation/go/xdr"
)

func PopulateAsset(ctx context.Context, dest *protocol.Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
