Returns information and links relating to a single [offer](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/offer).

## Request

```
GET /offers/{offer}
```

### Arguments

| name | notes | description | example |
| ---- | ----- | ----------- | ------- |
| `offer` | required, string | Offer ID | `126628073` |

### curl Example Request

```sh
curl "https://frontier.testnet.digitalbits.io/offers/1347876"
```

<!-- ### JavaScript Example Request -->

## Response

This endpoint responds with the details of a single offer for a given ID. See [offer resource](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/offer) for reference.

### Example Response

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/offers/1347876"
    },
    "offer_maker": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GAQHWQYBBW272OOXNQMMLCA5WY2XAZPODGB7Q3S5OKKIXVESKO55ZQ7C"
    }
  },
  "id": "1347876",
  "paging_token": "1347876",
  "seller": "GAQHWQYBBW272OOXNQMMLCA5WY2XAZPODGB7Q3S5OKKIXVESKO55ZQ7C",
  "selling": {
    "asset_type": "credit_alphanum4",
    "asset_code": "DSQ",
    "asset_issuer": "GBDQPTQJDATT7Z7EO4COS4IMYXH44RDLLI6N6WIL5BZABGMUOVMLWMQF"
  },
  "buying": {
    "asset_type": "credit_alphanum4",
    "asset_code": "USD",
    "asset_issuer": "GAA4MFNZGUPJAVLWWG6G5XZJFZDHLKQNG3Q6KB24BAD6JHNNVXDCF4XG"
  },
  "amount": "60.4544008",
  "price_r": {
    "n": 84293,
    "d": 2000000
  },
  "price": "0.0421465",
  "last_modified_ledger": 1429506,
  "last_modified_time": "2019-10-29T22:08:23Z"
}
```

## Possible Errors

- The [standard errors](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors#standard-errors).
- [not_found](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors/not-found): A `not_found` error will be returned if there is no offer whose ID matches the `offer` argument.
