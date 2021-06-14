This endpoint represents a single [data](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/data) associated with a given [account](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/account).

## Request

```
GET /accounts/{account}/data/{key}
```

### Arguments

| name     | notes                          | description                                                      | example                                                   |
| ------   | -------                        | -----------                                                      | -------                                                   |
| `key`| required, string               | Key name | `user-id`|

### curl Example Request

```sh
curl "https://frontier.testnet.digitalbits.io/accounts/GA2HGBJIJKI6O4XEM7CZWY5PS6GKSXL6D34ERAJYQSPYA6X6AI7HYW36/data/user-id"
```

### JavaScript Example Request

```javascript
var DigitalBitsSdk = require('digitalbits-sdk');
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

server.accounts()
  .accountId("GAKLBGHNHFQ3BMUYG5KU4BEWO6EYQHZHAXEWC33W34PH2RBHZDSQBD75")
  .call()
  .then(function (account) {
    return account.data({key: 'user-id'})
  })
  .then(function(dataValue) {
    console.log(dataValue)
  })
  .catch(function (err) {
    console.log(err)
  })
```

## Response

This endpoint responds with a value of the data field for the given account. See [data resource](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/data) for reference.

### Example Response

```json
{
  "value": "MTAw"
}
```

## Possible Errors

- The [standard errors](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors#standard-errors).
- [not_found](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors/not-found): A `not_found` error will be returned if there is no account whose ID matches the `account` argument or there is no data field with a given key.
