package main

import (
	"testing"

	hProtocol "github.com/digitalbits/go/protocols/frontier"
	"github.com/stretchr/testify/assert"
)

var badAmtOrders = []hProtocol.PriceLevel{hProtocol.PriceLevel{
	PriceR: hProtocol.Price{
		N: 4,
		D: 2,
	},
	Price:  "2.0",
	Amount: "amount",
}}

func TestConvertBids(t *testing.T) {
	usdXdbPrice := 0.10
	basePrice := 0.10
	bids, err := convertBids(badAmtOrders, usdXdbPrice, basePrice)
	assert.Error(t, err)
	assert.Equal(t, 0, len(bids))

	highBid := usdOrder{
		xdbAmount:  100.0,
		usdAmount:  10.0,
		baseAmount: 1.0,
		usdPrice:   0.25,
	}

	lowBid := usdOrder{
		xdbAmount:  50.0,
		usdAmount:  5.0,
		baseAmount: 0.5,
		usdPrice:   0.2,
	}

	bids, err = convertBids(hOrders, usdXdbPrice, basePrice)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, bids[0].usdPrice, bids[1].usdPrice)
	assert.Equal(t, highBid, bids[0])
	assert.Equal(t, lowBid, bids[1])
}

func TestConvertAsks(t *testing.T) {
	usdXdbPrice := 0.10
	basePrice := 0.10
	asks, err := convertAsks(badAmtOrders, usdXdbPrice, basePrice)
	assert.Error(t, err)
	assert.Equal(t, 0, len(asks))

	lowAsk := usdOrder{
		xdbAmount:  50,
		usdPrice:   0.2,
		usdAmount:  5,
		baseAmount: 0.5,
	}

	highAsk := usdOrder{
		xdbAmount:  100,
		usdPrice:   0.25,
		usdAmount:  10,
		baseAmount: 1,
	}

	orders := []hProtocol.PriceLevel{hHighOrder, hLowOrder}
	asks, err = convertAsks(orders, usdXdbPrice, basePrice)
	assert.NoError(t, err)
	assert.LessOrEqual(t, asks[0].usdPrice, asks[1].usdPrice)
	assert.Equal(t, lowAsk, asks[0])
	assert.Equal(t, highAsk, asks[1])
}
