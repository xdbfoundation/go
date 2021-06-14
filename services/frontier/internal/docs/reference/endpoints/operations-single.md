The operation details endpoint provides information on a single
[operation](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/operation). The operation ID provided in the `id` argument specifies
which operation to load.

### Warning - failed transactions

Operations can be part of successful or failed transactions (failed transactions are also included
in DigitalBits ledger). Always check operation status using `transaction_successful` field!

## Request

```
GET /operations/{id}
```

### Arguments

| name | notes | description | example |
| ---- | ----- | ----------- | ------- |
| `id` | required, number | An operation ID. | 2927608622747649 |
| `?join` | optional, string, default: _null_ | Set to `transactions` to include the transactions which created each of the operations in the response. | `transactions` |

### curl Example Request

```sh
curl https://frontier.testnet.digitalbits.io/operations/2927608622747649
```

### JavaScript Example Request

```javascript
var DigitalBitsSdk = require('digitalbits-sdk');
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

server.operations()
  .operation('2927608622747649')
  .call()
  .then(function (operationsResult) {
    console.log(operationsResult)
  })
  .catch(function (err) {
    console.log(err)
  })
```

## Response

This endpoint responds with a single Operation.  See [operation resource](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/operation) for reference.

### Example Response

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/operations/2927608622747649"
    },
    "transaction": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/4a3365180521e16b478d9f0c9198b97a9434fc9cb07b34f83ecc32fc54d0ca8a"
    },
    "effects": {
      "href": "https://frontier.testnet.digitalbits.io/operations/2927608622747649/effects"
    },
    "succeeds": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=2927608622747649"
    },
    "precedes": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=2927608622747649"
    }
  },
  "id": "2927608622747649",
  "paging_token": "2927608622747649",
  "transaction_successful": true,
  "source_account": "GCGXZPH2QNKJP4GI2J77EFQQUMP3NYY4PCUZ4UPKHR2XYBKRUYKQ2DS6",
  "type": "payment",
  "type_i": 1,
  "created_at": "2019-04-08T21:59:27Z",
  "transaction_hash": "4a3365180521e16b478d9f0c9198b97a9434fc9cb07b34f83ecc32fc54d0ca8a",
  "asset_type": "native",
  "from": "GCGXZPH2QNKJP4GI2J77EFQQUMP3NYY4PCUZ4UPKHR2XYBKRUYKQ2DS6",
  "to": "GDGEQS64ISS6Y2KDM5V67B6LXALJX4E7VE4MIA54NANSUX5MKGKBZM5G",
  "amount": "404.0000000"
}
```

## Possible Errors

- The [standard errors](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors#standard-errors).
- [not_found](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors/not-found): A `not_found` error will be returned if the
  there is no operation that matches the ID argument, i.e. the operation does not exist.
