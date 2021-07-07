package frontier

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/xdbfoundation/go/protocols/frontier"
	"github.com/xdbfoundation/go/services/frontier/internal/db2/history"
	"github.com/xdbfoundation/go/services/frontier/internal/ingest"
	"github.com/xdbfoundation/go/xdr"
)

func TestOfferActions_Show(t *testing.T) {
	ht := StartHTTPTestWithoutScenario(t)
	defer ht.Finish()
	q := &history.Q{ht.FrontierSession()}

	err := q.UpdateLastLedgerExpIngest(100)
	ht.Assert.NoError(err)
	err = q.UpdateExpIngestVersion(ingest.CurrentVersion)
	ht.Assert.NoError(err)

	ledgerCloseTime := time.Now().Unix()
	_, err = q.InsertLedger(xdr.LedgerHeaderHistoryEntry{
		Header: xdr.LedgerHeader{
			LedgerSeq: 100,
			ScpValue: xdr.DigitalBitsValue{
				CloseTime: xdr.TimePoint(ledgerCloseTime),
			},
		},
	}, 0, 0, 0, 0, 0)
	ht.Assert.NoError(err)

	issuer := xdr.MustAddress("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H")
	nativeAsset := xdr.MustNewNativeAsset()
	usdAsset := xdr.MustNewCreditAsset("USD", issuer.Address())
	eurAsset := xdr.MustNewCreditAsset("EUR", issuer.Address())

	eurOffer := history.Offer{
		SellerID: issuer.Address(),
		OfferID:  int64(4),

		BuyingAsset:  eurAsset,
		SellingAsset: nativeAsset,

		Amount:             int64(500),
		Pricen:             int32(1),
		Priced:             int32(1),
		Price:              float64(1),
		Flags:              1,
		LastModifiedLedger: uint32(3),
	}
	usdOffer := history.Offer{
		SellerID: issuer.Address(),
		OfferID:  int64(6),

		BuyingAsset:  usdAsset,
		SellingAsset: eurAsset,

		Amount:             int64(500),
		Pricen:             int32(1),
		Priced:             int32(1),
		Price:              float64(1),
		Flags:              1,
		LastModifiedLedger: uint32(4),
	}

	batch := q.NewOffersBatchInsertBuilder(3)
	err = batch.Add(eurOffer)
	ht.Assert.NoError(err)
	err = batch.Add(usdOffer)
	ht.Assert.NoError(err)
	ht.Assert.NoError(batch.Exec())

	w := ht.Get("/offers")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(2, w.Body)
	}

	w = ht.Get("/offers/4")
	if ht.Assert.Equal(200, w.Code) {
		var response frontier.Offer
		err = json.Unmarshal(w.Body.Bytes(), &response)
		ht.Assert.NoError(err)
		ht.Assert.Equal(int64(4), response.ID)
	}
}
