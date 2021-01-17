package tickerdb_test

import (
	"context"
	"fmt"
	"math"
	"testing"
	"time"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/digitalbits/go/services/ticker/internal/tickerdb"
	"github.com/digitalbits/go/services/ticker/internal/tickerdb/tickerdbtest"
	"github.com/digitalbits/go/support/db/dbtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRetrieveMarketData(t *testing.T) {
	db := dbtest.Postgres(t)
	defer db.Close()

	var session tickerdb.TickerSession
	session.DB = db.Open()
	session.Ctx = context.Background()
	defer session.DB.Close()

	// Run migrations to make sure the tests are run
	// on the most updated schema version
	migrations := &migrate.FileMigrationSource{
		Dir: "../migrations",
	}
	_, err := migrate.Exec(session.DB.DB, "postgres", migrations, migrate.Up)
	require.NoError(t, err)

	// Adding a seed issuer to be used later:
	tbl := session.GetTable("issuers")
	_, err = tbl.Insert(tickerdb.Issuer{
		PublicKey: "GCF3TQXKZJNFJK7HCMNE2O2CUNKCJH2Y2ROISTBPLC7C5EIA5NNG2XZB",
		Name:      "FOO BAR",
	}).IgnoreCols("id").Exec()
	require.NoError(t, err)
	var issuer tickerdb.Issuer
	err = session.GetRaw(&issuer, `
		SELECT *
		FROM issuers
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)

	// Adding a seed asset to be used later:
	err = session.InsertOrUpdateAsset(&tickerdb.Asset{
		Code:     "XDB",
		IssuerID: issuer.ID,
		IsValid:  true,
	}, []string{"code", "issuer_id"})
	require.NoError(t, err)
	var xdbAsset tickerdb.Asset
	err = session.GetRaw(&xdbAsset, `
		SELECT *
		FROM assets
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)

	// Adding another asset to be used later:
	err = session.InsertOrUpdateAsset(&tickerdb.Asset{
		Code:     "BTC",
		IssuerID: issuer.ID,
		IsValid:  true,
	}, []string{"code", "issuer_id"})
	require.NoError(t, err)
	var btcAsset tickerdb.Asset
	err = session.GetRaw(&btcAsset, `
		SELECT *
		FROM assets
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)

	// Adding a third asset:
	err = session.InsertOrUpdateAsset(&tickerdb.Asset{
		Code:     "ETH",
		IssuerID: issuer.ID,
		IsValid:  true,
	}, []string{"code", "issuer_id"})
	require.NoError(t, err)
	var ethAsset tickerdb.Asset
	err = session.GetRaw(&ethAsset, `
		SELECT *
		FROM assets
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)

	// Verify that we actually have three assets:
	assert.NotEqual(t, xdbAsset.ID, btcAsset.ID)
	assert.NotEqual(t, btcAsset.ID, ethAsset.ID)
	assert.NotEqual(t, xdbAsset.ID, ethAsset.ID)

	// A few times to be used:
	now := time.Now()
	oneHourAgo := now.Add(-1 * time.Hour)
	threeDaysAgo := now.AddDate(0, 0, -3)
	oneMonthAgo := now.AddDate(0, -1, 0)

	// Now let's create the trades:
	trades := []tickerdb.Trade{
		tickerdb.Trade{ // XDB_BTC trade
			FrontierID:       "hrzid1",
			BaseAssetID:     xdbAsset.ID,
			BaseAmount:      100.0,
			CounterAssetID:  btcAsset.ID,
			CounterAmount:   10.0,
			Price:           0.1,
			LedgerCloseTime: now,
		},
		tickerdb.Trade{ // XDB_ETH trade
			FrontierID:       "hrzid3",
			BaseAssetID:     xdbAsset.ID,
			BaseAmount:      24.0,
			CounterAssetID:  ethAsset.ID,
			CounterAmount:   26.0,
			Price:           0.92,
			LedgerCloseTime: oneHourAgo,
		},
		tickerdb.Trade{ // XDB_ETH trade
			FrontierID:       "hrzid2",
			BaseAssetID:     xdbAsset.ID,
			BaseAmount:      50.0,
			CounterAssetID:  ethAsset.ID,
			CounterAmount:   50.0,
			Price:           1.0,
			LedgerCloseTime: now,
		},
		tickerdb.Trade{ // XDB_BTC trade
			FrontierID:       "hrzid4",
			BaseAssetID:     xdbAsset.ID,
			BaseAmount:      50.0,
			CounterAssetID:  btcAsset.ID,
			CounterAmount:   6.0,
			Price:           0.12,
			LedgerCloseTime: threeDaysAgo,
		},
		tickerdb.Trade{ // XDB_ETH trade
			FrontierID:       "hrzid5",
			BaseAssetID:     xdbAsset.ID,
			BaseAmount:      24.0,
			CounterAssetID:  ethAsset.ID,
			CounterAmount:   28.0,
			Price:           1.10,
			LedgerCloseTime: oneMonthAgo,
		},
	}
	err = session.BulkInsertTrades(trades)
	require.NoError(t, err)

	// Adding some orderbook stats:
	obTime := time.Now()
	orderbookStats := tickerdb.OrderbookStats{
		BaseAssetID:    xdbAsset.ID,
		CounterAssetID: ethAsset.ID,
		NumBids:        15,
		BidVolume:      0.15,
		HighestBid:     200.0,
		NumAsks:        17,
		AskVolume:      30.0,
		LowestAsk:      0.1,
		Spread:         0.93,
		SpreadMidPoint: 0.35,
		UpdatedAt:      obTime,
	}
	err = session.InsertOrUpdateOrderbookStats(
		&orderbookStats,
		[]string{"base_asset_id", "counter_asset_id"},
	)
	require.NoError(t, err)

	var obBTCETH1 tickerdb.OrderbookStats
	err = session.GetRaw(&obBTCETH1, `
		SELECT *
		FROM orderbook_stats
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)

	orderbookStats = tickerdb.OrderbookStats{
		BaseAssetID:    xdbAsset.ID,
		CounterAssetID: btcAsset.ID,
		NumBids:        1,
		BidVolume:      0.1,
		HighestBid:     20.0,
		NumAsks:        1,
		AskVolume:      15.0,
		LowestAsk:      0.2,
		Spread:         0.96,
		SpreadMidPoint: 0.36,
		UpdatedAt:      obTime,
	}
	err = session.InsertOrUpdateOrderbookStats(
		&orderbookStats,
		[]string{"base_asset_id", "counter_asset_id"},
	)
	require.NoError(t, err)

	var obBTCETH2 tickerdb.OrderbookStats
	err = session.GetRaw(&obBTCETH2, `
		SELECT *
		FROM orderbook_stats
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)
	assert.NotEqual(t, obBTCETH1.ID, obBTCETH2.ID)

	markets, err := session.RetrieveMarketData()
	require.NoError(t, err)
	assert.Equal(t, 2, len(markets))

	// Mapping the retrieved markets:
	var xdbbtcMkt, xdbethMkt tickerdb.Market
	for _, mkt := range markets {
		if mkt.TradePair == "XDB_BTC" {
			xdbbtcMkt = mkt
		}

		if mkt.TradePair == "XDB_ETH" {
			xdbethMkt = mkt
		}
	}
	assert.NotEqual(t, "", xdbbtcMkt.TradePair)
	assert.NotEqual(t, "", xdbethMkt.TradePair)

	// Validating the aggregated data
	assert.Equal(t, 100.0, xdbbtcMkt.BaseVolume24h)
	assert.Equal(t, 10.0, xdbbtcMkt.CounterVolume24h)
	assert.Equal(t, int64(1), xdbbtcMkt.TradeCount24h)
	assert.Equal(t, 0.1, xdbbtcMkt.OpenPrice24h)
	assert.Equal(t, 0.1, xdbbtcMkt.LowestPrice24h)
	assert.Equal(t, 0.1, xdbbtcMkt.HighestPrice24h)

	assert.Equal(t, 150.0, xdbbtcMkt.BaseVolume7d)
	assert.Equal(t, 16.0, xdbbtcMkt.CounterVolume7d)
	assert.Equal(t, int64(2), xdbbtcMkt.TradeCount7d)
	assert.Equal(t, 0.12, xdbbtcMkt.OpenPrice7d)
	assert.Equal(t, 0.1, xdbbtcMkt.LowestPrice7d)
	assert.Equal(t, 0.12, xdbbtcMkt.HighestPrice7d)

	assert.Equal(t, 0.1, xdbbtcMkt.LastPrice)
	assert.WithinDuration(t, now.Local(), xdbbtcMkt.LastPriceCloseTime.Local(), 10*time.Millisecond)

	assert.Equal(t, 0.0, xdbbtcMkt.PriceChange24h)
	// There might be some floating point rounding issues, so this test
	// needs to be a bit more flexible. Since the change is 0.02, an error
	// around 0.0000000000001 is acceptable:
	priceChange7dDiff := math.Abs(-0.02 - xdbbtcMkt.PriceChange7d)
	assert.True(t, priceChange7dDiff < 0.0000000000001)

	assert.Equal(t, 74.0, xdbethMkt.BaseVolume24h)
	assert.Equal(t, 76.0, xdbethMkt.CounterVolume24h)
	assert.Equal(t, int64(2), xdbethMkt.TradeCount24h)
	assert.Equal(t, 0.92, xdbethMkt.OpenPrice24h)
	assert.Equal(t, 0.92, xdbethMkt.LowestPrice24h)
	assert.Equal(t, 1.0, xdbethMkt.HighestPrice24h)

	assert.Equal(t, 74.0, xdbethMkt.BaseVolume7d)
	assert.Equal(t, 76.0, xdbethMkt.CounterVolume7d)
	assert.Equal(t, int64(2), xdbethMkt.TradeCount7d)
	assert.Equal(t, 0.92, xdbethMkt.OpenPrice7d)
	assert.Equal(t, 0.92, xdbethMkt.LowestPrice7d)
	assert.Equal(t, 1.0, xdbethMkt.HighestPrice7d)

	assert.Equal(t, 1.0, xdbethMkt.LastPrice)
	assert.WithinDuration(t, now.Local(), xdbbtcMkt.LastPriceCloseTime.Local(), 10*time.Millisecond)

	// There might be some floating point rounding issues, so this test
	// needs to be a bit more flexible. Since the change is 0.08, an error
	// around 0.0000000000001 is acceptable:
	priceChange24hDiff := math.Abs(0.08 - xdbethMkt.PriceChange24h)
	assert.True(t, priceChange24hDiff < 0.0000000000001)

	priceChange7dDiff = math.Abs(0.08 - xdbethMkt.PriceChange7d)
	assert.True(t, priceChange7dDiff < 0.0000000000001)

	assert.Equal(t, priceChange24hDiff, priceChange7dDiff)

	// Analysing aggregated orderbook data:
	assert.Equal(t, 15, xdbethMkt.NumBids)
	assert.Equal(t, 0.15, xdbethMkt.BidVolume)
	assert.Equal(t, 200.0, xdbethMkt.HighestBid)
	assert.Equal(t, 17, xdbethMkt.NumAsks)
	assert.Equal(t, 30.0, xdbethMkt.AskVolume)
	assert.Equal(t, 0.1, xdbethMkt.LowestAsk)

	assert.Equal(t, 1, xdbbtcMkt.NumBids)
	assert.Equal(t, 0.1, xdbbtcMkt.BidVolume)
	assert.Equal(t, 20.0, xdbbtcMkt.HighestBid)
	assert.Equal(t, 1, xdbbtcMkt.NumAsks)
	assert.Equal(t, 15.0, xdbbtcMkt.AskVolume)
	assert.Equal(t, 0.2, xdbbtcMkt.LowestAsk)
}

func TestRetrievePartialMarkets(t *testing.T) {
	session := tickerdbtest.SetupTickerTestSession(t, "../migrations")
	defer session.DB.Close()

	issuer1PK := "GCF3TQXKZJNFJK7HCMNE2O2CUNKCJH2Y2ROISTBPLC7C5EIA5NNG2XZB"
	issuer2PK := "ABF3TQXKZJNFJK7HCMNE2O2CUNKCJH2Y2ROISTBPLC7C5EIA5NNG2XZB"
	now := time.Now()
	tenMinutesAgo := now.Add(-10 * time.Minute)
	oneHourAgo := now.Add(-1 * time.Hour)

	partialMkts, err := session.RetrievePartialMarkets(
		nil, nil, nil, nil, 12,
	)
	require.NoError(t, err)
	assert.Equal(t, 3, len(partialMkts))

	// Mapping the retrieved markets:
	var btceth1Mkt, btceth2Mkt tickerdb.PartialMarket
	for _, mkt := range partialMkts {
		if mkt.CounterAssetIssuer == issuer1PK && mkt.CounterAssetCode == "ETH" {
			btceth1Mkt = mkt
		}

		if mkt.CounterAssetIssuer == issuer2PK && mkt.CounterAssetCode == "ETH" {
			btceth2Mkt = mkt
		}
	}
	tradePair1 := fmt.Sprintf("BTC:%s / ETH:%s", issuer1PK, issuer1PK)
	tradePair2 := fmt.Sprintf("BTC:%s / ETH:%s", issuer1PK, issuer2PK)

	assert.Equal(t, tradePair1, btceth1Mkt.TradePairName)
	assert.Equal(t, tradePair2, btceth2Mkt.TradePairName)

	// Validating the aggregated data
	assert.Equal(t, 150.0, btceth1Mkt.BaseVolume)
	assert.Equal(t, 60.0, btceth1Mkt.CounterVolume)
	assert.Equal(t, int32(2), btceth1Mkt.TradeCount)
	assert.Equal(t, 1.0, btceth1Mkt.Open)
	assert.Equal(t, 0.1, btceth1Mkt.Close)
	assert.Equal(t, -0.9, btceth1Mkt.Change)
	assert.Equal(t, 1.0, btceth1Mkt.High)
	assert.Equal(t, 0.1, btceth1Mkt.Low)
	assert.WithinDuration(t, oneHourAgo.Local(), btceth1Mkt.FirstLedgerCloseTime.Local(), 100*time.Millisecond)
	assert.WithinDuration(t, tenMinutesAgo.Local(), btceth1Mkt.LastLedgerCloseTime.Local(), 100*time.Millisecond)
	assert.Equal(t, 24.0, btceth2Mkt.BaseVolume)
	assert.Equal(t, 26.0, btceth2Mkt.CounterVolume)
	assert.Equal(t, int32(1), btceth2Mkt.TradeCount)
	assert.Equal(t, 0.92, btceth2Mkt.Open)
	assert.Equal(t, 0.92, btceth2Mkt.Close)
	assert.Equal(t, 0.0, btceth2Mkt.Change)
	assert.Equal(t, 0.92, btceth2Mkt.High)
	assert.Equal(t, 0.92, btceth2Mkt.Low)
	assert.WithinDuration(t, now.Local(), btceth2Mkt.FirstLedgerCloseTime.Local(), 100*time.Millisecond)
	assert.WithinDuration(t, now.Local(), btceth2Mkt.LastLedgerCloseTime.Local(), 100*time.Millisecond)

	// Analyzing non-aggregated orderbook data
	assert.Equal(t, 15, btceth1Mkt.NumBids)
	assert.Equal(t, 0.15, btceth1Mkt.BidVolume)
	assert.Equal(t, 200.0, btceth1Mkt.HighestBid)
	assert.Equal(t, 17, btceth1Mkt.NumAsks)
	assert.Equal(t, 30.0, btceth1Mkt.AskVolume)
	assert.Equal(t, 0.1, btceth1Mkt.LowestAsk)
	assert.Equal(t, 10, btceth1Mkt.NumBidsReverse)
	assert.Equal(t, 100.0, btceth1Mkt.HighestBidReverse)
	assert.Equal(t, 12, btceth1Mkt.NumAsksReverse)
	assert.Equal(t, 0.2, btceth1Mkt.LowestAskReverse)

	assert.Equal(t, 1, btceth2Mkt.NumBids)
	assert.Equal(t, 0.1, btceth2Mkt.BidVolume)
	assert.Equal(t, 20.0, btceth2Mkt.HighestBid)
	assert.Equal(t, 1, btceth2Mkt.NumAsks)
	assert.Equal(t, 15.0, btceth2Mkt.AskVolume)
	assert.Equal(t, 0.2, btceth2Mkt.LowestAsk)
	assert.Equal(t, 20, btceth2Mkt.NumBidsReverse)
	assert.Equal(t, 0.60, btceth2Mkt.BidVolumeReverse)
	assert.Equal(t, 300.0, btceth2Mkt.HighestBidReverse)
	assert.Equal(t, 20, btceth2Mkt.NumAsksReverse)
	assert.Equal(t, 256.0, btceth2Mkt.AskVolumeReverse)
	assert.Equal(t, 0.70, btceth2Mkt.LowestAskReverse)

	// Now let's use the same data, but aggregating by asset pair
	partialAggMkts, err := session.RetrievePartialAggMarkets(nil, nil, 12)
	require.NoError(t, err)
	assert.Equal(t, 2, len(partialAggMkts))

	partialAggMkt := partialAggMkts[0]

	btcEthStr := "BTC_ETH"
	assert.Equal(t, btcEthStr, partialAggMkt.TradePairName)
	assert.Equal(t, 174.0, partialAggMkt.BaseVolume)
	assert.Equal(t, 86.0, partialAggMkt.CounterVolume)
	assert.Equal(t, int32(3), partialAggMkt.TradeCount)
	assert.Equal(t, 1.0, partialAggMkt.Open)
	assert.Equal(t, 0.92, partialAggMkt.Close)
	assert.Equal(t, 1.0, partialAggMkt.High)
	assert.Equal(t, 0.1, partialAggMkt.Low)
	assert.WithinDuration(t, oneHourAgo.Local(), partialAggMkt.FirstLedgerCloseTime.Local(), 100*time.Millisecond)
	assert.WithinDuration(t, now.Local(), partialAggMkt.LastLedgerCloseTime.Local(), 100*time.Millisecond)

	// There might be some floating point rounding issues, so this test
	// needs to be a bit more flexible. Since the change is 0.08, an error
	// around 0.0000000000001 is acceptable:
	priceDiff := math.Abs(-0.08 - partialAggMkt.Change)
	assert.True(t, priceDiff < 0.0000000000001)

	// Validate the pair name parsing:
	pairNames := []*string{&btcEthStr}
	partialAggMkts, err = session.RetrievePartialAggMarkets(nil, &pairNames, 12)
	require.NoError(t, err)
	assert.Equal(t, 1, len(partialAggMkts))
	assert.Equal(t, int32(3), partialAggMkts[0].TradeCount)

	// Analyzing aggregated orderbook data:
	assert.Equal(t, 16, partialAggMkt.NumBids)
	assert.Equal(t, 0.25, partialAggMkt.BidVolume)
	assert.Equal(t, 200.0, partialAggMkt.HighestBid)
	assert.Equal(t, 18, partialAggMkt.NumAsks)
	assert.Equal(t, 45.0, partialAggMkt.AskVolume)
	assert.Equal(t, 0.1, partialAggMkt.LowestAsk)
	assert.Equal(t, 30, partialAggMkt.NumBidsReverse)
	assert.Equal(t, 1.50, partialAggMkt.BidVolumeReverse)
	assert.Equal(t, 300.0, partialAggMkt.HighestBidReverse)
	assert.Equal(t, 32, partialAggMkt.NumAsksReverse)
	assert.Equal(t, 281.0, partialAggMkt.AskVolumeReverse)
	assert.Equal(t, 0.2, partialAggMkt.LowestAskReverse)

	// Validate that both markets are parsed.
	btcXdbStr := "BTC_XDB"
	pairNames = []*string{&btcEthStr, &btcXdbStr}
	partialAggMkts, err = session.RetrievePartialAggMarkets(nil, &pairNames, 12)
	require.NoError(t, err)
	assert.Equal(t, 2, len(partialAggMkts))
	assert.Equal(t, int32(3), partialAggMkts[0].TradeCount)
	assert.Equal(t, int32(2), partialAggMkts[1].TradeCount)

	// Validate that passing a code works.
	btcStr := "BTC"
	partialAggMkts, err = session.RetrievePartialAggMarkets(&btcStr, nil, 12)
	require.NoError(t, err)
	assert.Equal(t, 2, len(partialAggMkts))

	// Make sure there's an error with a non-nil code and non-nil pair names.
	partialAggMkts, err = session.RetrievePartialAggMarkets(&btcStr, &pairNames, 12)
	require.Error(t, err)
}

func Test24hStatsFallback(t *testing.T) {
	db := dbtest.Postgres(t)
	defer db.Close()

	var session tickerdb.TickerSession
	session.DB = db.Open()
	session.Ctx = context.Background()
	defer session.DB.Close()

	// Run migrations to make sure the tests are run
	// on the most updated schema version
	migrations := &migrate.FileMigrationSource{
		Dir: "../migrations",
	}
	_, err := migrate.Exec(session.DB.DB, "postgres", migrations, migrate.Up)
	require.NoError(t, err)

	// Adding a seed issuer to be used later:
	tbl := session.GetTable("issuers")
	_, err = tbl.Insert(tickerdb.Issuer{
		PublicKey: "GCF3TQXKZJNFJK7HCMNE2O2CUNKCJH2Y2ROISTBPLC7C5EIA5NNG2XZB",
		Name:      "FOO BAR",
	}).IgnoreCols("id").Exec()
	require.NoError(t, err)
	var issuer tickerdb.Issuer
	err = session.GetRaw(&issuer, `
		SELECT *
		FROM issuers
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)

	// Adding a seed asset to be used later:
	err = session.InsertOrUpdateAsset(&tickerdb.Asset{
		Code:     "XDB",
		IssuerID: issuer.ID,
		IsValid:  true,
	}, []string{"code", "issuer_id"})
	require.NoError(t, err)
	var xdbAsset tickerdb.Asset
	err = session.GetRaw(&xdbAsset, `
		SELECT *
		FROM assets
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)

	// Adding another asset to be used later:
	err = session.InsertOrUpdateAsset(&tickerdb.Asset{
		Code:     "BTC",
		IssuerID: issuer.ID,
		IsValid:  true,
	}, []string{"code", "issuer_id"})
	require.NoError(t, err)
	var btcAsset tickerdb.Asset
	err = session.GetRaw(&btcAsset, `
		SELECT *
		FROM assets
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)

	// A few times to be used:
	now := time.Now()
	twoDaysAgo := now.AddDate(0, 0, -3)
	threeDaysAgo := now.AddDate(0, 0, -3)

	// Now let's create the trades:
	trades := []tickerdb.Trade{
		tickerdb.Trade{
			FrontierID:       "hrzid1",
			BaseAssetID:     xdbAsset.ID,
			BaseAmount:      1.0,
			CounterAssetID:  btcAsset.ID,
			CounterAmount:   1.0,
			Price:           0.5, // close price & lowest price
			LedgerCloseTime: twoDaysAgo,
		},
		tickerdb.Trade{ // BTC_ETH trade (ETH is from issuer 2)
			FrontierID:       "hrzid2",
			BaseAssetID:     xdbAsset.ID,
			BaseAmount:      1.0,
			CounterAssetID:  btcAsset.ID,
			CounterAmount:   1.0,
			Price:           1.0, // open price & highest price
			LedgerCloseTime: threeDaysAgo,
		},
	}
	err = session.BulkInsertTrades(trades)
	require.NoError(t, err)

	markets, err := session.RetrieveMarketData()
	require.NoError(t, err)
	assert.Equal(t, 1, len(markets))
	mkt := markets[0]

	// When there are no 24h data, 24h OHLC should fallback to the 7d close value
	assert.Equal(t, 0.5, mkt.LastPrice)
	assert.Equal(t, 0.5, mkt.LowestPrice24h)
	assert.Equal(t, 0.5, mkt.OpenPrice24h)
	assert.Equal(t, 0.5, mkt.HighestPrice24h)
}

func TestPreferAnchorAssetCode(t *testing.T) {
	db := dbtest.Postgres(t)
	defer db.Close()

	var session tickerdb.TickerSession
	session.DB = db.Open()
	session.Ctx = context.Background()
	defer session.DB.Close()

	// Run migrations to make sure the tests are run
	// on the most updated schema version
	migrations := &migrate.FileMigrationSource{
		Dir: "../migrations",
	}
	_, err := migrate.Exec(session.DB.DB, "postgres", migrations, migrate.Up)
	require.NoError(t, err)

	// Adding a seed issuer to be used later:
	tbl := session.GetTable("issuers")
	_, err = tbl.Insert(tickerdb.Issuer{
		PublicKey: "GCF3TQXKZJNFJK7HCMNE2O2CUNKCJH2Y2ROISTBPLC7C5EIA5NNG2XZB",
		Name:      "FOO BAR",
	}).IgnoreCols("id").Exec()
	require.NoError(t, err)
	var issuer tickerdb.Issuer
	err = session.GetRaw(&issuer, `
		SELECT *
		FROM issuers
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)

	// Adding a seed asset to be used later:
	err = session.InsertOrUpdateAsset(&tickerdb.Asset{
		Code:     "XDB",
		IssuerID: issuer.ID,
		IsValid:  true,
	}, []string{"code", "issuer_id"})
	require.NoError(t, err)
	var xdbAsset tickerdb.Asset
	err = session.GetRaw(&xdbAsset, `
		SELECT *
		FROM assets
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)

	// Adding another asset to be used later:
	err = session.InsertOrUpdateAsset(&tickerdb.Asset{
		Code:            "EURT",
		IssuerID:        issuer.ID,
		IsValid:         true,
		AnchorAssetCode: "EUR",
	}, []string{"code", "issuer_id"})
	require.NoError(t, err)
	var btcAsset tickerdb.Asset
	err = session.GetRaw(&btcAsset, `
		SELECT *
		FROM assets
		ORDER BY id DESC
		LIMIT 1`,
	)
	require.NoError(t, err)

	// A few times to be used:
	now := time.Now()
	twoDaysAgo := now.AddDate(0, 0, -3)
	threeDaysAgo := now.AddDate(0, 0, -3)

	// Now let's create the trades:
	trades := []tickerdb.Trade{
		tickerdb.Trade{
			FrontierID:       "hrzid1",
			BaseAssetID:     xdbAsset.ID,
			BaseAmount:      1.0,
			CounterAssetID:  btcAsset.ID,
			CounterAmount:   1.0,
			Price:           0.5, // close price & lowest price
			LedgerCloseTime: twoDaysAgo,
		},
		tickerdb.Trade{ // BTC_ETH trade (ETH is from issuer 2)
			FrontierID:       "hrzid2",
			BaseAssetID:     xdbAsset.ID,
			BaseAmount:      1.0,
			CounterAssetID:  btcAsset.ID,
			CounterAmount:   1.0,
			Price:           1.0, // open price & highest price
			LedgerCloseTime: threeDaysAgo,
		},
	}
	err = session.BulkInsertTrades(trades)
	require.NoError(t, err)

	markets, err := session.RetrieveMarketData()
	require.NoError(t, err)
	require.Equal(t, 1, len(markets))
	for _, mkt := range markets {
		require.Equal(t, "XDB_EUR", mkt.TradePair)
	}

	partialAggMkts, err := session.RetrievePartialAggMarkets(nil, nil, 168)
	require.NoError(t, err)
	assert.Equal(t, 1, len(partialAggMkts))
	for _, aggMkt := range partialAggMkts {
		require.Equal(t, "XDB_EUR", aggMkt.TradePairName)
	}
}
