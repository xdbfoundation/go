This endpoint represents successful [operations](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/operation) that were included in valid [transactions](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/transaction) that affected a particular [account](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/account).

This endpoint can also be used in [streaming](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/streaming) mode so it is possible to use it to listen for new operations that affect a given account as they happen.
If called in streaming mode Frontier will start at the earliest known operation unless a `cursor` is set. In that case it will start from the `cursor`. You can also set `cursor` value to `now` to only stream operations created since your request time.

## Request

```
GET /accounts/{account}/operations{?cursor,limit,order,include_failed}
```

### Arguments

| name     | notes                          | description                                                      | example                                                   |
| ------   | -------                        | -----------                                                      | -------                                                   |
| `account`| required, string               | Account ID                                                  | `GA2HGBJIJKI6O4XEM7CZWY5PS6GKSXL6D34ERAJYQSPYA6X6AI7HYW36`|
| `?cursor`| optional, default _null_       | A paging token, specifying where to start returning records from.  When streaming this can be set to `now` to stream object created since your request time. | `12884905984`                                             |
| `?order` | optional, string, default `asc`| The order in which to return rows, "asc" or "desc".              | `asc`                                                     |
| `?limit` | optional, number, default `10` | Maximum number of records to return.                             | `200`
| `?include_failed` | optional, bool, default: `false` | Set to `true` to include operations of failed transactions in results. | `true` |                                                     |
| `?join` | optional, string, default: _null_ | Set to `transactions` to include the transactions which created each of the operations in the response. | `transactions` |

### curl Example Request

```sh
curl "https://frontier.testnet.digitalbits.io/accounts/GA2HGBJIJKI6O4XEM7CZWY5PS6GKSXL6D34ERAJYQSPYA6X6AI7HYW36/operations"
```

### JavaScript Example Request

```js
var DigitalBitsSdk = require('digitalbits-sdk');
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

server.operations()
  .forAccount("GAKLBGHNHFQ3BMUYG5KU4BEWO6EYQHZHAXEWC33W34PH2RBHZDSQBD75")
  .call()
  .then(function (operationsResult) {
    console.log(operationsResult.records)
  })
  .catch(function (err) {
    console.log(err)
  })
```

### JavaScript Streaming Example

```javascript
var DigitalBitsSdk = require('digitalbits-sdk')
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

var operationHandler = function (operationResponse) {
    console.log(operationResponse);
};

var es = server.operations()
    .forAccount("GAKLBGHNHFQ3BMUYG5KU4BEWO6EYQHZHAXEWC33W34PH2RBHZDSQBD75")
    .cursor('now')
    .stream({
        onmessage: operationHandler
    })
```

## Response

This endpoint responds with a list of operations that affected the given account. See [operation resource](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/operation) for reference.

### Example Response

```json
{
  "_embedded": {
    "records": [
      {
        "_links": {
          "effects": {
            "href": "/operations/46316927324160/effects/{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "/operations?cursor=46316927324160&order=asc"
          },
          "self": {
            "href": "/operations/46316927324160"
          },
          "succeeds": {
            "href": "/operations?cursor=46316927324160&order=desc"
          },
          "transactions": {
            "href": "/transactions/46316927324160"
          }
        },
        "account": "GBBM6BKZPEHWYO3E3YKREDPQXMS4VK35YLNU7NFBRI26RAN7GI5POFBB",
        "funder": "GBIA4FH6TV64KSPDAJCNUQSM7PFL4ILGUVJDPCLUOPJ7ONMKBBVUQHRO",
        "id": 46316927324160,
        "paging_token": "46316927324160",
        "starting_balance": 1e+09,
        "transaction_successful": true,
        "type_i": 0,
        "type": "create_account"
      }
    ]
  },
  "_links": {
    "next": {
      "href": "/accounts/GBBM6BKZPEHWYO3E3YKREDPQXMS4VK35YLNU7NFBRI26RAN7GI5POFBB/operations?order=asc&limit=10&cursor=46316927324160"
    },
    "prev": {
      "href": "/accounts/GBBM6BKZPEHWYO3E3YKREDPQXMS4VK35YLNU7NFBRI26RAN7GI5POFBB/operations?order=desc&limit=10&cursor=46316927324160"
    },
    "self": {
      "href": "/accounts/GBBM6BKZPEHWYO3E3YKREDPQXMS4VK35YLNU7NFBRI26RAN7GI5POFBB/operations?order=asc&limit=10&cursor="
    }
  }
}
```

### Example Streaming Event

```json
{
  "_links": {
    "effects": {
      "href": "/operations/77309415424/effects/{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "/operations?cursor=77309415424&order=asc"
    },
    "self": {
      "href": "/operations/77309415424"
    },
    "succeeds": {
      "href": "/operations?cursor=77309415424&order=desc"
    },
    "transactions": {
      "href": "/transactions/77309415424"
    }
  },
  "account": "GBIA4FH6TV64KSPDAJCNUQSM7PFL4ILGUVJDPCLUOPJ7ONMKBBVUQHRO",
  "funder": "GCEZWKCA5VLDNRLN3RPRJMRZOX3Z6G5CHCGSNFHEYVXM3XOJMDS674JZ",
  "id": 77309415424,
  "paging_token": "77309415424",
  "starting_balance": "1000.0000000",
  "transaction_successful": true,
  "type_i": 0,
  "type": "create_account"
}
```


## Possible Errors

- The [standard errors](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors#standard-errors).
- [not_found](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors/not-found): A `not_found` error will be returned if there is no account whose ID matches the `account` argument.
