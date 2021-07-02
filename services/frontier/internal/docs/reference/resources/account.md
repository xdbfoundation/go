---
title: Account
---

In the DigitalBits network, users interact using **accounts** which can be controlled by a corresponding keypair that can authorize transactions. One can create a new account with the [Create Account](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/resources/operation.md#create-account) operation.

To learn more about the concept of accounts in the DigitalBits network, take a look at the [DigitalBits account concept guide](https://github.com/xdbfoundation/docs/blob/master/guides/concepts/accounts.md).

When frontier returns information about an account it uses the following format:

## Attributes
| Attribute      | Type             | Description                                                                                                                                  |
|----------------|------------------|------------------------------------------------------------------------------------------------------------------------                      |
| id             | string           | The canonical id of this account, suitable for use as the :id parameter for url templates that require an account's ID.                      |
| account_id     | string           | The account's public key encoded into a base32 string representation.                                                                        |
| sequence       | number           | The current sequence number that can be used when submitting a transaction from this account.                                                |
| subentry_count | number           | The number of [account subentries](https://github.com/xdbfoundation/docs/blob/master/guides/concepts/ledger.md#ledger-entries).                           |
| balances       | array of objects | An array of the native asset or credits this account holds.                                                                                  |
| thresholds     | object           | An object of account thresholds.                                                                                                             |
| flags          | object           | The flags denote the enabling/disabling of certain asset issuer privileges.                                                                  |
| signers        | array of objects | An array of [account signers](https://github.com/xdbfoundation/docs/blob/master/guides/concepts/multi-sig.md#additional-signing-keys) with their weights. |
| data           | object           | An array of account [data](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/resources/data.md) fields.                                                                                                |

### Signer Object
| Attribute  | Type   | Description                                                                                                      |
|------------|--------|------------------------------------------------------------------------------------------------------------------|
| weight     | number | The numerical weight of a signer, necessary to determine whether a transaction meets the threshold requirements. |
| key        | string | Different depending on the type of the signer.                                                                   |
| type       | string | See below.                                                                                                       |

### Possible Signer Types
| Type               | Description                                                                                                                                                                                                         |
|--------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ed25519_public_key | A normal DigitalBits public key.                                                                                                                                                                                        |
| sha256_hash        | The SHA256 hash of some arbitrary `x`. Adding a signature of this type allows anyone who knows `x` to sign a transaction from this account. *Note: Once this transaction is broadcast, `x` will be known publicly.* |
| preauth_tx         | The hash of a pre-authorized transaction. This signer is automatically removed from the account when a matching transaction is properly applied.                                                                    |
### Balances Object
| Attribute           | Type             |                                                                                                                                                                                                                                                               |
|---------------------|------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| balance             | string           | How much of an asset is owned.                                                                                                                                                                                                                                |
| buying_liabilities  | string           | The total amount of an asset offered to buy aggregated over all offers owned by this account.                                                                                                                                                                 |
| selling_liabilities | string           | The total amount of an asset offered to sell aggregated over all offers owned by this account.                                                                                                                                                                |
| limit               | optional, number | The maximum amount of an asset that this account is willing to accept (this is specified when an account opens a trustline).                                                                                                                                  |
| asset_type          | string           | Either native, credit_alphanum4, or credit_alphanum12.                                                                                                                                                                                                        |
| asset_code          | optional, string | The code for the asset.                                                                                                                                                                                                                                       |
| asset_issuer        | optional, string | The digitalbits address of the given asset's issuer.                                                                                                                                                                                                              |
| is_authorized       | optional, bool   | The trustline status for an `auth_required` asset.  If true, the issuer of the asset has granted the account permission to send, receive, buy, or sell the asset.  If false, the issuer has not, so the account cannot send, receive, buy, or sell the asset. |

### Flag Object
| Attribute      | Type |                                                                                                                                |
|----------------|------|--------------------------------------------------------------------------------------------------------------------------------|
| auth_immutable | bool | With this setting, none of the following authorization flags can be changed.                                                   |
| auth_required  | bool | With this setting, an anchor must approve anyone who wants to hold its asset.                                                  |
| auth_revocable | bool | With this setting, an anchor can set the authorize flag of an existing trustline to freeze the assets held by an asset holder. |

### Threshold Object
| Attribute      | Type   |                                                                                                                                                                                                                                                                                                                                                                           |
|----------------|--------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| low_threshold  | number | The weight required for a valid transaction including the [Allow Trust][allow_trust] and [Bump Sequence][bump_seq] operations.                                                                                                                                                                                                                                            |
| med_threshold  | number | The weight required for a valid transaction including the [Create Account][create_acc], [Payment][payment], [Path Payment Strict Send][path_payment_send], [Path Payment Strict Receive][path_payment_receive], [Manage Buy Offer][manage_buy_offer], [Manage Sell Offer][manage_sell_offer], [Create Passive Sell Offer][passive_sell_offer], [Change Trust][change_trust], [Inflation][inflation], and [Manage Data][manage_data] operations. |
| high_threshold | number | The weight required for a valid transaction including the [Account Merge][account_merge] and [Set Options]() operations.                                                                                                                                                                                                                                                  |

[account_merge]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#account-merge
[allow_trust]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#allow-trust
[bump_seq]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#bump-sequence
[change_trust]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#change-trust
[create_acc]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#create-account
[inflation]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#inflation
[manage_data]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#manage-data
[manage_buy_offer]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#manage-buy-offer
[manage_sell_offer]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#manage-sell-offer
[passive_sell_offer]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#create-passive-sell-offer
[path_payment_receive]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#path-payment-strict-receive
[path_payment_send]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#path-payment-strict-send
[payment]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#payment
[set_options]: https://github.com/xdbfoundation/docs/blob/master/guides/concepts/list-of-operations.md#set-options

## Links
| rel          | Example                                                                                                 | Description                                                  | `templated` |
|--------------|---------------------------------------------------------------------------------------------------------|--------------------------------------------------------------|-------------|
| data         | `/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/data/{key}`                         | [Data fields](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/resources/data.md) related to this account             | true        |
| effects      | `/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/effects/{?cursor,limit,order}`      | The [effects](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/resources/effect.md) related to this account           | true        |
| offers       | `/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/offers/{?cursor,limit,order}`       | The [offers](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/resources/offer.md) related to this account             | true        |
| operations   | `/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/operations/{?cursor,limit,order}`   | The [operations](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/resources/operation.md) related to this account     | true        |
| payments     | `/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/payments/{?cursor,limit,order}`     | The [payments](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/resources/payment.md) related to this account         | true        |
| trades       | `/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/trades/{?cursor,limit,order}`       | The [trades](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/resources/trade.md) related to this account             | true        |
| transactions | `/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/transactions/{?cursor,limit,order}` | The [transactions](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/resources/transaction.md) related to this account | true        |

## Example

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY"
    },
    "transactions": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/transactions{?cursor,limit,order}",
      "templated": true
    },
    "operations": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/operations{?cursor,limit,order}",
      "templated": true
    },
    "payments": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/payments{?cursor,limit,order}",
      "templated": true
    },
    "effects": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/effects{?cursor,limit,order}",
      "templated": true
    },
    "offers": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/offers{?cursor,limit,order}",
      "templated": true
    },
    "trades": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/trades{?cursor,limit,order}",
      "templated": true
    },
    "data": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/data/{key}",
      "templated": true
    }
  },
  "id": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
  "account_id": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
  "sequence": "4113023891406853",
  "subentry_count": 4,
  "last_modified_ledger": 958207,
  "last_modified_time": "2021-06-16T09:49:53Z",
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
      "balance": "5.0000000",
      "limit": "1000.0000000",
      "buying_liabilities": "0.0000000",
      "selling_liabilities": "0.0000000",
      "last_modified_ledger": 957774,
      "is_authorized": true,
      "is_authorized_to_maintain_liabilities": true,
      "asset_type": "credit_alphanum4",
      "asset_code": "EUR",
      "asset_issuer": "GDCIQQY2UKVNLLWGIX74DMTEAFCMQKAKYUWPBO7PLTHIHRKSFZN7V2FC"
    },
    {
      "balance": "50000.0000000",
      "limit": "922337203685.4775807",
      "buying_liabilities": "0.0000000",
      "selling_liabilities": "0.0000000",
      "last_modified_ledger": 958064,
      "is_authorized": true,
      "is_authorized_to_maintain_liabilities": true,
      "asset_type": "credit_alphanum4",
      "asset_code": "HUF",
      "asset_issuer": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK"
    },
    {
      "balance": "10.0000000",
      "limit": "1000.0000000",
      "buying_liabilities": "0.0000000",
      "selling_liabilities": "0.0000000",
      "last_modified_ledger": 957671,
      "is_authorized": true,
      "is_authorized_to_maintain_liabilities": true,
      "asset_type": "credit_alphanum4",
      "asset_code": "USD",
      "asset_issuer": "GB4RZUSF3HZGCAKB3VBM2S7QOHHC5KTV3LLZXGBYR5ZO4B26CKHFZTSZ"
    },
    {
      "balance": "9999.9999500",
      "buying_liabilities": "0.0000000",
      "selling_liabilities": "0.0000000",
      "asset_type": "native"
    }
  ],
  "signers": [
    {
      "weight": 1,
      "key": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
      "type": "ed25519_public_key"
    }
  ],
  "data": {
    "user-id": "WERCRm91bmRhdGlvbg=="
  },
  "num_sponsoring": 0,
  "num_sponsored": 0,
  "paging_token": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY"
}
```

## Endpoints
| Resource                                                         | Type       | Resource URI Template                |
|------------------------------------------------------------------|------------|--------------------------------------|
| [Account Details](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/accounts-single.md)               | Single     | `/accounts/:id`                      |
| [Account Data](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/data-for-account.md)                 | Single     | `/accounts/:id/data/:key`            |
| [Account Transactions](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/transactions-for-account.md) | Collection | `/accounts/:account_id/transactions` |
| [Account Operations](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/operations-for-account.md)     | Collection | `/accounts/:account_id/operations`   |
| [Account Payments](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/payments-for-account.md)         | Collection | `/accounts/:account_id/payments`     |
| [Account Effects](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/effects-for-account.md)           | Collection | `/accounts/:account_id/effects`      |
| [Account Offers](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/offers-for-account.md)             | Collection | `/accounts/:account_id/offers`       |
