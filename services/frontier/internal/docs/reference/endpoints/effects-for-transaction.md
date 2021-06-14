This endpoint represents all [effects](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/effect) that occurred as a result of a given [transaction](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/transaction).

## Request

```
GET /transactions/{hash}/effects{?cursor,limit,order}
```

## Arguments

| name | notes | description | example |
| ---- | ----- | ----------- | ------- |
| `hash` | required, string | A transaction hash, hex-encoded, lowercase. | `7e2050abc676003efc3eaadd623c927f753b7a6c37f50864bf284f4e1510d088` |
| `?cursor` | optional, default _null_ | A paging token, specifying where to start returning records from. | `12884905984` |
| `?order` | optional, string, default `asc` | The order in which to return rows, "asc" or "desc". | `asc` |
| `?limit` | optional, number, default `10` | Maximum number of records to return. | `200` |

### curl Example Request

```sh
curl "https://frontier.testnet.digitalbits.io/transactions/7e2050abc676003efc3eaadd623c927f753b7a6c37f50864bf284f4e1510d088/effects?limit=1"
```

### JavaScript Example Request

```javascript
var DigitalBitsSdk = require('digitalbits-sdk');
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

server.effects()
  .forTransaction("7e2050abc676003efc3eaadd623c927f753b7a6c37f50864bf284f4e1510d088")
  .call()
  .then(function (effectResults) {
    //page 1
    console.log(effectResults.records)
  })
  .catch(function (err) {
    console.log(err)
  })

```

## Response

This endpoint responds with a list of effects on the ledger as a result of a given transaction. See [effect resource](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/effect) for reference.

### Example Response

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/7e2050abc676003efc3eaadd623c927f753b7a6c37f50864bf284f4e1510d088/effects?cursor=&limit=10&order=asc"
    },
    "next": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/7e2050abc676003efc3eaadd623c927f753b7a6c37f50864bf284f4e1510d088/effects?cursor=1919197546291201-3&limit=10&order=asc"
    },
    "prev": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/7e2050abc676003efc3eaadd623c927f753b7a6c37f50864bf284f4e1510d088/effects?cursor=1919197546291201-1&limit=10&order=desc"
    }
  },
  "_embedded": {
    "records": [
      {
        "_links": {
          "operation": {
            "href": "https://frontier.testnet.digitalbits.io/operations/1919197546291201"
          },
          "succeeds": {
            "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=1919197546291201-1"
          },
          "precedes": {
            "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=1919197546291201-1"
          }
        },
        "id": "0001919197546291201-0000000001",
        "paging_token": "1919197546291201-1",
        "account": "GBYUUJHG6F4EPJGNLERINATVQLNDOFRUD7SGJZ26YZLG5PAYLG7XUSGF",
        "type": "account_created",
        "type_i": 0,
        "created_at": "2019-03-25T22:43:38Z",
        "starting_balance": "10000.0000000"
      },
      {
        "_links": {
          "operation": {
            "href": "https://frontier.testnet.digitalbits.io/operations/1919197546291201"
          },
          "succeeds": {
            "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=1919197546291201-2"
          },
          "precedes": {
            "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=1919197546291201-2"
          }
        },
        "id": "0001919197546291201-0000000002",
        "paging_token": "1919197546291201-2",
        "account": "GAIH3ULLFQ4DGSECF2AR555KZ4KNDGEKN4AFI4SU2M7B43MGK3QJZNSR",
        "type": "account_debited",
        "type_i": 3,
        "created_at": "2019-03-25T22:43:38Z",
        "asset_type": "native",
        "amount": "10000.0000000"
      },
      {
        "_links": {
          "operation": {
            "href": "https://frontier.testnet.digitalbits.io/operations/1919197546291201"
          },
          "succeeds": {
            "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=1919197546291201-3"
          },
          "precedes": {
            "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=1919197546291201-3"
          }
        },
        "id": "0001919197546291201-0000000003",
        "paging_token": "1919197546291201-3",
        "account": "GBYUUJHG6F4EPJGNLERINATVQLNDOFRUD7SGJZ26YZLG5PAYLG7XUSGF",
        "type": "signer_created",
        "type_i": 10,
        "created_at": "2019-03-25T22:43:38Z",
        "weight": 1,
        "public_key": "GBYUUJHG6F4EPJGNLERINATVQLNDOFRUD7SGJZ26YZLG5PAYLG7XUSGF",
        "key": ""
      }
    ]
  }
}
```

## Errors

- The [standard errors](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors#standard-errors).
- [not_found](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors/not-found): A `not_found` error will be returned if there are no effects for transaction whose hash matches the `hash` argument.
