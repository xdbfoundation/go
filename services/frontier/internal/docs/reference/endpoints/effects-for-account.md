This endpoint represents all [effects](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/effect) that changed a given
[account](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/account). It will return relevant effects from the creation of the
account to the current ledger.

This endpoint can also be used in [streaming](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/streaming) mode so it is possible to use it to
listen for new effects as transactions happen in the DigitalBits network.
If called in streaming mode Frontier will start at the earliest known effect unless a `cursor` is
set. In that case it will start from the `cursor`. You can also set `cursor` value to `now` to only
stream effects created since your request time.

## Request

```
GET /accounts/{account}/effects{?cursor,limit,order}
```

## Arguments

| name | notes | description | example |
| ---- | ----- | ----------- | ------- |
| `account` | required, string | Account ID | `GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY` |
| `?cursor` | optional, default _null_ | A paging token, specifying where to start returning records from. When streaming this can be set to `now` to stream object created since your request time. | `12884905984` |
| `?order`  | optional, string, default `asc` | The order in which to return rows, "asc" or "desc". | `asc` |
| `?limit`  | optional, number, default `10` | Maximum number of records to return. | `200` |

### curl Example Request

```sh
curl "https://frontier.testnet.digitalbits.io/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/effects?limit=1"
```

### JavaScript Example Request

```javascript
var DigitalBitsSdk = require('digitalbits-sdk');
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

server.effects()
  .forAccount("GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY")
  .call()
  .then(function (effectResults) {
    // page 1
    console.log(JSON.strigify(effectResults.records))
  })
  .catch(function (err) {
    console.log(err)
  })
```

### JavaScript Streaming Example

```javascript
var DigitalBitsSdk = require('digitalbits-sdk')
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

var effectHandler = function (effectResponse) {
  console.log(effectResponse);
};

var es = server.effects()
  .forAccount("GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY")
  .cursor('now')
  .stream({
    onmessage: effectHandler
  })
```

## Response

The list of effects.

### Example Response

```json
[
  {
    "_links": {
      "operation": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113023891410945"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4113023891410945-1"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4113023891410945-1"
      }
    },
    "id": "0004113023891410945-0000000001",
    "paging_token": "4113023891410945-1",
    "account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "account_created",
    "type_i": 0,
    "created_at": "2021-06-16T08:55:24Z",
    "starting_balance": "10000.0000000"
  },
  {
    "_links": {
      "operation": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113023891410945"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4113023891410945-3"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4113023891410945-3"
      }
    },
    "id": "0004113023891410945-0000000003",
    "paging_token": "4113023891410945-3",
    "account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "signer_created",
    "type_i": 10,
    "created_at": "2021-06-16T08:55:24Z",
    "weight": 1,
    "public_key": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "key": ""
  },
  {
    "_links": {
      "operation": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113161330364417"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4113161330364417-1"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4113161330364417-1"
      }
    },
    "id": "0004113161330364417-0000000001",
    "paging_token": "4113161330364417-1",
    "account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "trustline_created",
    "type_i": 20,
    "created_at": "2021-06-16T08:58:29Z",
    "asset_type": "credit_alphanum4",
    "asset_code": "USD",
    "asset_issuer": "GB4RZUSF3HZGCAKB3VBM2S7QOHHC5KTV3LLZXGBYR5ZO4B26CKHFZTSZ",
    "limit": "1000.0000000"
  },
  {
    "_links": {
      "operation": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113165625331713"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4113165625331713-1"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4113165625331713-1"
      }
    },
    "id": "0004113165625331713-0000000001",
    "paging_token": "4113165625331713-1",
    "account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "account_credited",
    "type_i": 2,
    "created_at": "2021-06-16T08:58:35Z",
    "asset_type": "credit_alphanum4",
    "asset_code": "USD",
    "asset_issuer": "GB4RZUSF3HZGCAKB3VBM2S7QOHHC5KTV3LLZXGBYR5ZO4B26CKHFZTSZ",
    "amount": "10.0000000"
  },
  {
    "_links": {
      "operation": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113603711995905"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4113603711995905-1"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4113603711995905-1"
      }
    },
    "id": "0004113603711995905-0000000001",
    "paging_token": "4113603711995905-1",
    "account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "trustline_created",
    "type_i": 20,
    "created_at": "2021-06-16T09:08:23Z",
    "asset_type": "credit_alphanum4",
    "asset_code": "EUR",
    "asset_issuer": "GDCIQQY2UKVNLLWGIX74DMTEAFCMQKAKYUWPBO7PLTHIHRKSFZN7V2FC",
    "limit": "1000.0000000"
  },
  {
    "_links": {
      "operation": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113608006963201"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4113608006963201-1"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4113608006963201-1"
      }
    },
    "id": "0004113608006963201-0000000001",
    "paging_token": "4113608006963201-1",
    "account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "account_credited",
    "type_i": 2,
    "created_at": "2021-06-16T09:08:28Z",
    "asset_type": "credit_alphanum4",
    "asset_code": "EUR",
    "asset_issuer": "GDCIQQY2UKVNLLWGIX74DMTEAFCMQKAKYUWPBO7PLTHIHRKSFZN7V2FC",
    "amount": "5.0000000"
  },
  {
    "_links": {
      "operation": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4114797712904193"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4114797712904193-1"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4114797712904193-1"
      }
    },
    "id": "0004114797712904193-0000000001",
    "paging_token": "4114797712904193-1",
    "account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "trustline_created",
    "type_i": 20,
    "created_at": "2021-06-16T09:34:52Z",
    "asset_type": "credit_alphanum4",
    "asset_code": "HUF",
    "asset_issuer": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "limit": "1000.0000000"
  },
  {
    "_links": {
      "operation": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4114849252511745"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4114849252511745-1"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4114849252511745-1"
      }
    },
    "id": "0004114849252511745-0000000001",
    "paging_token": "4114849252511745-1",
    "account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "trustline_updated",
    "type_i": 22,
    "created_at": "2021-06-16T09:35:59Z",
    "asset_type": "credit_alphanum4",
    "asset_code": "HUF",
    "asset_issuer": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "limit": "922337203685.4775807"
  },
  {
    "_links": {
      "operation": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4114853547479041"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4114853547479041-1"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4114853547479041-1"
      }
    },
    "id": "0004114853547479041-0000000001",
    "paging_token": "4114853547479041-1",
    "account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "account_credited",
    "type_i": 2,
    "created_at": "2021-06-16T09:36:04Z",
    "asset_type": "credit_alphanum4",
    "asset_code": "HUF",
    "asset_issuer": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "amount": "50000.0000000"
  },
  {
    "_links": {
      "operation": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4115467727802369"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4115467727802369-1"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4115467727802369-1"
      }
    },
    "id": "0004115467727802369-0000000001",
    "paging_token": "4115467727802369-1",
    "account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "data_created",
    "type_i": 40,
    "created_at": "2021-06-16T09:49:53Z",
    "name": "user-id",
    "value": "WERCRm91bmRhdGlvbg=="
  }
]

```

## Possible Errors

- The [standard errors](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors#standard-errors).
- [not_found](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors/not-found): A `not_found` error will be returned if there are no effects for the given account.
