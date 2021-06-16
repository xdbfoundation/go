Returns information and links relating to a single [account](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/account).

The balances section in the returned JSON will also list all the
[trustlines](https://developers.digitalbits.io/guides/docs/guides/concepts/assets#trustlines) this account
established. Note this will only return trustlines that have the necessary authorization to work.
Meaning if an account `A` trusts another account `B` that has the
[authorization required](https://developers.digitalbits.io/guides/docs/guides/concepts/accounts#flags)
flag set, the trustline won't show up until account `B`
[allows](https://developers.digitalbits.io/guides/docs/guides/concepts/list-of-operations#allow-trust)
account `A` to hold its assets.

## Request

```
GET /accounts/{account}
```

### Arguments

| name | notes | description | example |
| ---- | ----- | ----------- | ------- |
| `account` | required, string | Account ID | `GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT` |

### curl Example Request

```sh
curl "https://frontier.testnet.digitalbits.io/accounts/GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT"
```

### JavaScript Example Request

```javascript
var DigitalBitsSdk = require('digitalbits-sdk');
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

server.accounts()
  .accountId("GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT")
  .call()
  .then(function (accountResult) {
    console.log(JSON.stringify(accountResult));
  })
  .catch(function (err) {
    console.error(err);
  })
```

## Response

This endpoint responds with the details of a single account for a given ID. See [account resource](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/account) for reference.

### Example Response
```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT"
    },
    "transactions": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT/transactions{?cursor,limit,order}",
      "templated": true
    },
    "operations": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT/operations{?cursor,limit,order}",
      "templated": true
    },
    "payments": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT/payments{?cursor,limit,order}",
      "templated": true
    },
    "effects": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT/effects{?cursor,limit,order}",
      "templated": true
    },
    "offers": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT/offers{?cursor,limit,order}",
      "templated": true
    },
    "trades": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT/trades{?cursor,limit,order}",
      "templated": true
    },
    "data": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT/data/{key}",
      "templated": true
    }
  },
  "id": "GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT",
  "account_id": "GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT",
  "sequence": "4056944503422978",
  "subentry_count": 0,
  "last_modified_ledger": 944602,
  "last_modified_time": "2021-06-15T12:08:01Z",
  "thresholds": {
    "low_threshold": 0,
    "med_threshold": 0,
    "high_threshold": 0
  },
  "flags": {
    "auth_required": false,
    "auth_revocable": false,
    "auth_immutable": false
  },
  "balances": [
    {
      "balance": "9599.9999800",
      "buying_liabilities": "0.0000000",
      "selling_liabilities": "0.0000000",
      "asset_type": "native"
    }
  ],
  "signers": [
    {
      "weight": 1,
      "key": "GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT",
      "type": "ed25519_public_key"
    }
  ],
  "num_sponsoring": 0,
  "num_sponsored": 0,
  "paging_token": "GCLHMNEO2XW24POXS7HCWKHK5O5ZTM7R3DNRIKJS3DTHBWVLSVACYAOT",
  "data_attr": {}
}
```

## Possible Errors

- The [standard errors](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors#standard-errors).
- [not_found](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors/not-found): A `not_found` error will be returned if there is no account whose ID matches the `account` argument.
