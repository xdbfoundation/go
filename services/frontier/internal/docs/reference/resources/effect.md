---
title: Effect
---

A successful operation will yield zero or more **effects**.  These effects
represent specific changes that occur in the ledger, but are not necessarily
directly reflected in the [ledger](https://github.com/xdbfoundation/docs/blob/master/guides/concepts/ledger.md) or [history](https://github.com/xdbfoundation/DigitalBits/blob/master/docs/history.md), as [transactions](https://github.com/xdbfoundation/docs/blob/master/guides/concepts/transactions.md) and [operations](https://github.com/xdbfoundation/docs/blob/master/guides/concepts/operations.md) are.

## Effect types

We can distinguish 6 effect groups:

- Account effects
- Signer effects
- Trustline effects
- Trading effects
- Data effects
- Misc effects

### Account effects

| Type                                  | Operation                                            |
|---------------------------------------|------------------------------------------------------|
| Account Created                       | create_account                                       |
| Account Removed                       | merge_account                                        |
| Account Credited                      | create_account, payment, path_payment, merge_account |
| Account Debited                       | create_account, payment, path_payment, merge_account |
| Account Thresholds Updated            | set_options                                          |
| Account Home Domain Updated           | set_options                                          |
| Account Flags Updated                 | set_options                                          |
| Account Inflation Destination Updated | set_options                                          |

### Signer effects

| Type           | Operation   |
|----------------|-------------|
| Signer Created | set_options |
| Signer Removed | set_options |
| Signer Updated | set_options |

### Trustline effects

| Type                   | Operation                 |
|------------------------|---------------------------|
| Trustline Created      | change_trust              |
| Trustline Removed      | change_trust              |
| Trustline Updated      | change_trust, allow_trust |
| Trustline Authorized   | allow_trust               |
| Trustline Deauthorized | allow_trust               |

### Trading effects

| Type          | Operation                                                                    |
|---------------|------------------------------------------------------------------------------|
| Offer Created | manage_buy_offer, manage_sell_offer, create_passive_sell_offer               |
| Offer Removed | manage_buy_offer, manage_sell_offer, create_passive_sell_offer, path_payment |
| Offer Updated | manage_buy_offer, manage_sell_offer, create_passive_sell_offer, path_payment |
| Trade         | manage_buy_offer, manage_sell_offer, create_passive_sell_offer, path_payment |
### Data effects

| Type         | Operation   |
|--------------|-------------|
| Data Created | manage_data |
| Data Removed | manage_data |
| Data Updated | manage_data |
### Misc effects

| Type            | Operation     |
|-----------------|---------------|
| Sequence Bumped | bump_sequence |

## Attributes

Attributes depend on effect type.

## Links

| rel       | Example                                                       | Relation                          |
|-----------|---------------------------------------------------------------|-----------------------------------|
| self      | `/effects?order=asc\u0026limit=1`                             |                                   |
| prev      | `/effects?order=desc\u0026limit=1\u0026cursor=141733924865-1` |                                   |
| next      | `/effects?order=asc\u0026limit=1\u0026cursor=141733924865-1`  |                                   |
| operation | `/operations/141733924865`                                    | Operation that created the effect |

## Example

```json
  {
    "_links": {
      "operation": {
        "href": "https://frontier.testnet.digitalbits.io/operations/1099511631876"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=1099511631876-1"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=1099511631876-1"
      }
    },
    "id": "0000001099511631876-0000000001",
    "paging_token": "1099511631876-1",
    "account": "GA4E6NUVQE7UO6723X6BYVPZMV3CLYXV64KJOOJ7K37F7KYHUD2C6FV4",
    "type": "account_created",
    "type_i": 0,
    "created_at": "2021-04-13T13:55:32Z",
    "starting_balance": "101.0000000"
  }

```

## Endpoints

| Resource                                                                                                                                   | Type       | Resource URI Template           |
|--------------------------------------------------------------------------------------------------------------------------------------------|------------|---------------------------------|
| [All Effects](../endpoints/effects-all.md)                 | Collection | `/effects`                      |
| [Operation Effects](../endpoints/effects-for-operation.md) | Collection | `/operations/:id/effects`       |
| [Account Effects](../endpoints/effects-for-account.md)     | Collection | `/accounts/:account_id/effects` |
| [Ledger Effects](../endpoints/effects-for-ledger.md)       | Collection | `/ledgers/:ledger_id/effects`   |
