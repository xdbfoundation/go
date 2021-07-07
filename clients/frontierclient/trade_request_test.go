package frontierclient

import (
	"context"
	"testing"

	hProtocol "github.com/xdbfoundation/go/protocols/frontier"
	"github.com/xdbfoundation/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTradeRequestBuildUrl(t *testing.T) {
	tr := TradeRequest{}
	endpoint, err := tr.BuildURL()

	// It should return valid all trades endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "trades", endpoint)

	tr = TradeRequest{ForAccount: "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU"}
	endpoint, err = tr.BuildURL()

	// It should return valid account trades endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU/trades", endpoint)

	tr = TradeRequest{ForOfferID: "123"}
	endpoint, err = tr.BuildURL()

	// It should return valid offer trades endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "offers/123/trades", endpoint)

	tr = TradeRequest{Cursor: "123"}
	endpoint, err = tr.BuildURL()

	// It should return valid trades endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "trades?cursor=123", endpoint)

	tr = TradeRequest{ForOfferID: "123", ForAccount: "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU"}
	_, err = tr.BuildURL()

	// error case: too many parameters for building any operation endpoint
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "invalid request: too many parameters")
	}

	tr = TradeRequest{Cursor: "123456", Limit: 30, Order: OrderAsc}
	endpoint, err = tr.BuildURL()
	// It should return valid all trades endpoint with query params and no errors
	require.NoError(t, err)
	assert.Equal(t, "trades?cursor=123456&limit=30&order=asc", endpoint)

}

func TestTradesRequest(t *testing.T) {
	hmock := httptest.NewClient()
	client := &Client{
		FrontierURL: "https://localhost/",
		HTTP:       hmock,
	}

	tradeRequest := TradeRequest{}

	// all trades
	hmock.On(
		"GET",
		"https://localhost/trades",
	).ReturnString(200, tradesResponse)

	trades, err := client.Trades(tradeRequest)
	if assert.NoError(t, err) {
		assert.IsType(t, trades, hProtocol.TradesPage{})
		links := trades.Links
		assert.Equal(t, links.Self.Href, "https://frontier-testnet.digitalbits.org/trades?cursor=&limit=2&order=desc")

		assert.Equal(t, links.Next.Href, "https://frontier-testnet.digitalbits.org/trades?cursor=2099298409914407-0&limit=2&order=desc")

		assert.Equal(t, links.Prev.Href, "https://frontier-testnet.digitalbits.org/trades?cursor=2099319884746791-0&limit=2&order=asc")

		trade := trades.Embedded.Records[0]
		assert.IsType(t, trade, hProtocol.Trade{})
		assert.Equal(t, trade.ID, "2099319884746791-0")
		assert.Equal(t, trade.BaseAmount, "2.4104452")
		assert.Equal(t, trade.CounterAmount, "0.0973412")
		assert.Equal(t, trade.OfferID, "3698823")
		assert.Equal(t, trade.BaseIsSeller, false)
	}

	tradeRequest = TradeRequest{ForAccount: "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU"}
	hmock.On(
		"GET",
		"https://localhost/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU/trades",
	).ReturnString(200, tradesResponse)

	trades, err = client.Trades(tradeRequest)
	if assert.NoError(t, err) {
		assert.IsType(t, trades, hProtocol.TradesPage{})
	}

	// too many parameters
	tradeRequest = TradeRequest{ForAccount: "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU", ForOfferID: "123"}
	hmock.On(
		"GET",
		"https://localhost/trades",
	).ReturnString(200, "")

	_, err = client.Trades(tradeRequest)
	// error case
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "too many parameters")
	}
}

func TestNextTradesPage(t *testing.T) {
	hmock := httptest.NewClient()
	client := &Client{
		FrontierURL: "https://localhost/",
		HTTP:       hmock,
	}

	tradeRequest := TradeRequest{ForAccount: "GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG", Limit: 2}

	hmock.On(
		"GET",
		"https://localhost/accounts/GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG/trades?limit=2",
	).ReturnString(200, firstTradesPage)

	trades, err := client.Trades(tradeRequest)

	if assert.NoError(t, err) {
		assert.Equal(t, len(trades.Embedded.Records), 2)
	}

	hmock.On(
		"GET",
		"https://frontier-testnet.digitalbits.org/accounts/GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG/trades?cursor=45122926424065-0&limit=2&order=asc",
	).ReturnString(200, emptyTradesPage)

	nextPage, err := client.NextTradesPage(trades)
	if assert.NoError(t, err) {
		assert.Equal(t, len(nextPage.Embedded.Records), 0)
	}
}

func TestTradeRequestStreamTrades(t *testing.T) {

	hmock := httptest.NewClient()
	client := &Client{
		FrontierURL: "https://localhost/",
		HTTP:       hmock,
	}

	// all trades
	trRequest := TradeRequest{}
	ctx, cancel := context.WithCancel(context.Background())

	hmock.On(
		"GET",
		"https://localhost/trades?cursor=now",
	).ReturnString(200, tradeStreamResponse)

	trades := make([]hProtocol.Trade, 1)
	err := client.StreamTrades(ctx, trRequest, func(tr hProtocol.Trade) {
		trades[0] = tr
		cancel()
	})

	if assert.NoError(t, err) {
		assert.Equal(t, trades[0].ID, "76909979385857-0")
		assert.Equal(t, trades[0].OfferID, "494")
	}

	// trades for accounts
	trRequest = TradeRequest{ForAccount: "GCRHQBHX7JNBZE4HHPLNOAAYDRDVAGBJKJ4KPGHIID3CBGVALXBD6TVQ"}
	ctx, cancel = context.WithCancel(context.Background())

	hmock.On(
		"GET",
		"https://localhost/accounts/GCRHQBHX7JNBZE4HHPLNOAAYDRDVAGBJKJ4KPGHIID3CBGVALXBD6TVQ/trades?cursor=now",
	).ReturnString(200, tradeStreamResponse)

	trades = make([]hProtocol.Trade, 1)
	err = client.StreamTrades(ctx, trRequest, func(tr hProtocol.Trade) {
		trades[0] = tr
		cancel()
	})

	if assert.NoError(t, err) {
		assert.Equal(t, trades[0].ID, "76909979385857-0")
		assert.Equal(t, trades[0].OfferID, "494")
	}

	// trades for offers
	trRequest = TradeRequest{ForOfferID: "494"}
	ctx, cancel = context.WithCancel(context.Background())

	hmock.On(
		"GET",
		"https://localhost/offers/494/trades?cursor=now",
	).ReturnString(200, tradeStreamResponse)

	trades = make([]hProtocol.Trade, 1)
	err = client.StreamTrades(ctx, trRequest, func(tr hProtocol.Trade) {
		trades[0] = tr
		cancel()
	})

	if assert.NoError(t, err) {
		assert.Equal(t, trades[0].ID, "76909979385857-0")
		assert.Equal(t, trades[0].OfferID, "494")
	}

	// test error
	trRequest = TradeRequest{}
	ctx, cancel = context.WithCancel(context.Background())

	hmock.On(
		"GET",
		"https://localhost/trades?cursor=now",
	).ReturnString(500, tradeStreamResponse)

	trades = make([]hProtocol.Trade, 1)
	err = client.StreamTrades(ctx, trRequest, func(tr hProtocol.Trade) {
		cancel()
	})

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "got bad HTTP status code 500")
	}
}

var tradesResponse = `{
  "_links": {
    "self": {
      "href": "https://frontier-testnet.digitalbits.org/trades?cursor=&limit=2&order=desc"
    },
    "next": {
      "href": "https://frontier-testnet.digitalbits.org/trades?cursor=2099298409914407-0&limit=2&order=desc"
    },
    "prev": {
      "href": "https://frontier-testnet.digitalbits.org/trades?cursor=2099319884746791-0&limit=2&order=asc"
    }
  },
  "_embedded": {
    "records": [
      {
        "_links": {
          "self": {
            "href": ""
          },
          "base": {
            "href": "https://frontier-testnet.digitalbits.org/accounts/GAQHWQYBBW272OOXNQMMLCA5WY2XAZPODGB7Q3S5OKKIXVESKO55ZQ7C"
          },
          "counter": {
            "href": "https://frontier-testnet.digitalbits.org/accounts/GCYN7MI6VXVRP74KR6MKBAW2ELLCXL6QCY5H4YQ62HVWZWMCE6Y232UC"
          },
          "operation": {
            "href": "https://frontier-testnet.digitalbits.org/operations/2099319884746791"
          }
        },
        "id": "2099319884746791-0",
        "paging_token": "2099319884746791-0",
        "ledger_close_time": "2019-03-28T10:45:28Z",
        "offer_id": "3698823",
        "base_offer_id": "4613785338312134695",
        "base_account": "GAQHWQYBBW272OOXNQMMLCA5WY2XAZPODGB7Q3S5OKKIXVESKO55ZQ7C",
        "base_amount": "2.4104452",
        "base_asset_type": "credit_alphanum4",
        "base_asset_code": "DSQ",
        "base_asset_issuer": "GBDQPTQJDATT7Z7EO4COS4IMYXH44RDLLI6N6WIL5BZABGMUOVMLWMQF",
        "counter_offer_id": "3698823",
        "counter_account": "GCYN7MI6VXVRP74KR6MKBAW2ELLCXL6QCY5H4YQ62HVWZWMCE6Y232UC",
        "counter_amount": "0.0973412",
        "counter_asset_type": "credit_alphanum4",
        "counter_asset_code": "USD",
        "counter_asset_issuer": "GAA4MFNZGUPJAVLWWG6G5XZJFZDHLKQNG3Q6KB24BAD6JHNNVXDCF4XG",
        "base_is_seller": false,
        "price": {
          "n": 2000000,
          "d": 49525693
        }
      },
      {
        "_links": {
          "self": {
            "href": ""
          },
          "base": {
            "href": "https://frontier-testnet.digitalbits.org/accounts/GAQHWQYBBW272OOXNQMMLCA5WY2XAZPODGB7Q3S5OKKIXVESKO55ZQ7C"
          },
          "counter": {
            "href": "https://frontier-testnet.digitalbits.org/accounts/GCYN7MI6VXVRP74KR6MKBAW2ELLCXL6QCY5H4YQ62HVWZWMCE6Y232UC"
          },
          "operation": {
            "href": "https://frontier-testnet.digitalbits.org/operations/2099298409914407"
          }
        },
        "id": "2099298409914407-0",
        "paging_token": "2099298409914407-0",
        "ledger_close_time": "2019-03-28T10:45:02Z",
        "offer_id": "3698823",
        "base_offer_id": "4613785316837302311",
        "base_account": "GAQHWQYBBW272OOXNQMMLCA5WY2XAZPODGB7Q3S5OKKIXVESKO55ZQ7C",
        "base_amount": "89.3535843",
        "base_asset_type": "credit_alphanum4",
        "base_asset_code": "DSQ",
        "base_asset_issuer": "GBDQPTQJDATT7Z7EO4COS4IMYXH44RDLLI6N6WIL5BZABGMUOVMLWMQF",
        "counter_offer_id": "3698823",
        "counter_account": "GCYN7MI6VXVRP74KR6MKBAW2ELLCXL6QCY5H4YQ62HVWZWMCE6Y232UC",
        "counter_amount": "3.6083729",
        "counter_asset_type": "credit_alphanum4",
        "counter_asset_code": "USD",
        "counter_asset_issuer": "GAA4MFNZGUPJAVLWWG6G5XZJFZDHLKQNG3Q6KB24BAD6JHNNVXDCF4XG",
        "base_is_seller": false,
        "price": {
          "n": 2000000,
          "d": 49525693
        }
      }
    ]
  }
}`

var tradeStreamResponse = `data: {"_links":{"self":{"href":""},"base":{"href":"https://frontier-testnet.digitalbits.org/accounts/GCRHQBHX7JNBZE4HHPLNOAAYDRDVAGBJKJ4KPGHIID3CBGVALXBD6TVQ"},"counter":{"href":"https://frontier-testnet.digitalbits.org/accounts/GAEETTPUI5CO3CSYXXM5CRX4FHLDWJ3KD6XRRJ3GJISWQSCYF5ALN6JC"},"operation":{"href":"https://frontier-testnet.digitalbits.org/operations/76909979385857"}},"id":"76909979385857-0","paging_token":"76909979385857-0","ledger_close_time":"2019-02-28T11:29:40Z","offer_id":"494","base_offer_id":"4611762928406773761","base_account":"GCRHQBHX7JNBZE4HHPLNOAAYDRDVAGBJKJ4KPGHIID3CBGVALXBD6TVQ","base_amount":"0.0000001","base_asset_type":"native","counter_offer_id":"494","counter_account":"GAEETTPUI5CO3CSYXXM5CRX4FHLDWJ3KD6XRRJ3GJISWQSCYF5ALN6JC","counter_amount":"0.0001000","counter_asset_type":"credit_alphanum4","counter_asset_code":"WTF","counter_asset_issuer":"GAQZKAGUAHCN4OHAMQVQ3PNA5DUHCQ3CEVOSOTPUAXHG3UHTRSSUFHUL","base_is_seller":false,"price":{"n":1000,"d":1}}
`

var firstTradesPage = `{
  "_links": {
    "self": {
      "href": "https://frontier-testnet.digitalbits.org/accounts/GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG/trades?cursor=&limit=2&order=asc"
    },
    "next": {
      "href": "https://frontier-testnet.digitalbits.org/accounts/GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG/trades?cursor=45122926424065-0&limit=2&order=asc"
    },
    "prev": {
      "href": "https://frontier-testnet.digitalbits.org/accounts/GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG/trades?cursor=45097156620289-0&limit=2&order=desc"
    }
  },
  "_embedded": {
    "records": [
      {
        "_links": {
          "self": {
            "href": ""
          },
          "base": {
            "href": "https://frontier-testnet.digitalbits.org/accounts/GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG"
          },
          "counter": {
            "href": "https://frontier-testnet.digitalbits.org/accounts/GBH77NK3ZP7RT52YZWGIU5Y6VTIJ52VXUSXDMQ7Z7VAAQO4U4QGGIROV"
          },
          "operation": {
            "href": "https://frontier-testnet.digitalbits.org/operations/45097156620289"
          }
        },
        "id": "45097156620289-0",
        "paging_token": "45097156620289-0",
        "ledger_close_time": "2019-04-25T02:29:20Z",
        "offer_id": "1219",
        "base_offer_id": "928",
        "base_account": "GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG",
        "base_amount": "2.7922715",
        "base_asset_type": "credit_alphanum4",
        "base_asset_code": "HT",
        "base_asset_issuer": "GCNSGHUCG5VMGLT5RIYYZSO7VQULQKAJ62QA33DBC5PPBSO57LFWVV6P",
        "counter_offer_id": "1219",
        "counter_account": "GBH77NK3ZP7RT52YZWGIU5Y6VTIJ52VXUSXDMQ7Z7VAAQO4U4QGGIROV",
        "counter_amount": "0.0012000",
        "counter_asset_type": "credit_alphanum4",
        "counter_asset_code": "BTC",
        "counter_asset_issuer": "GCNSGHUCG5VMGLT5RIYYZSO7VQULQKAJ62QA33DBC5PPBSO57LFWVV6P",
        "base_is_seller": false,
        "price": {
          "n": 383,
          "d": 891200
        }
      },
      {
        "_links": {
          "self": {
            "href": ""
          },
          "base": {
            "href": "https://frontier-testnet.digitalbits.org/accounts/GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG"
          },
          "counter": {
            "href": "https://frontier-testnet.digitalbits.org/accounts/GBH77NK3ZP7RT52YZWGIU5Y6VTIJ52VXUSXDMQ7Z7VAAQO4U4QGGIROV"
          },
          "operation": {
            "href": "https://frontier-testnet.digitalbits.org/operations/45122926424065"
          }
        },
        "id": "45122926424065-0",
        "paging_token": "45122926424065-0",
        "ledger_close_time": "2019-04-25T02:29:49Z",
        "offer_id": "928",
        "base_offer_id": "928",
        "base_account": "GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG",
        "base_amount": "2.7956854",
        "base_asset_type": "credit_alphanum4",
        "base_asset_code": "HT",
        "base_asset_issuer": "GCNSGHUCG5VMGLT5RIYYZSO7VQULQKAJ62QA33DBC5PPBSO57LFWVV6P",
        "counter_offer_id": "4611731141353811969",
        "counter_account": "GBH77NK3ZP7RT52YZWGIU5Y6VTIJ52VXUSXDMQ7Z7VAAQO4U4QGGIROV",
        "counter_amount": "0.0012000",
        "counter_asset_type": "credit_alphanum4",
        "counter_asset_code": "BTC",
        "counter_asset_issuer": "GCNSGHUCG5VMGLT5RIYYZSO7VQULQKAJ62QA33DBC5PPBSO57LFWVV6P",
        "base_is_seller": true,
        "price": {
          "n": 7973,
          "d": 18575000
        }
      }
    ]
  }
}`

var emptyTradesPage = `{
  "_links": {
    "self": {
      "href": "https://frontier-testnet.digitalbits.org/accounts/GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG/trades?cursor=45122926424065-0&limit=2&order=asc"
    },
    "next": {
      "href": "https://frontier-testnet.digitalbits.org/accounts/GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG/trades?cursor=59889023983617-0&limit=2&order=asc"
    },
    "prev": {
      "href": "https://frontier-testnet.digitalbits.org/accounts/GBZ5OD56VRTRQKMNADD6VUZUG3FCILMAMYQY5ZSC3AW3GBXNEPIK76IG/trades?cursor=45810121191425-0&limit=2&order=desc"
    }
  },
  "_embedded": {
    "records": []
  }
}`
