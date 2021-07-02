---
title: Ledger
---

A **ledger** resource contains information about a given ledger.

To learn more about the concept of ledgers in the DigitalBits network, take a look at the [DigitalBits ledger concept guide](https://developers.digitalbits.io/guides/concepts/ledger.html).

## Attributes

| Attribute                    | Type   |                                                                                                                              |
|------------------------------|--------|------------------------------------------------------------------------------------------------------------------------------|
| id                           | string | The id is a unique identifier for this ledger.                                                                               |
| paging_token                 | number | A [paging token](./page.md) suitable for use as a `cursor` parameter.                                                        |
| hash                         | string | A hex-encoded, lowercase SHA-256 hash of the ledger's [XDR](../xdr.md)-encoded form.                                |
| prev_hash                    | string | The hash of the ledger that chronologically came before this one.                                                            |
| sequence                     | number | Sequence number of this ledger, suitable for use as the as the :id parameter for url templates that require a ledger number. |
| successful_transaction_count | number | The number of successful transactions in this ledger.                                                                        |
| failed_transaction_count     | number | The number of failed transactions in this ledger.                                                                            |
| operation_count              | number | The number of operations applied in this ledger.                                                                             |
| tx_set_operation_count       | number | The number of operations in this ledger. This number includes operations from failed and successful transactions.            |
| closed_at                    | string | An [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) formatted string of when this ledger was closed.                       |
| total_coins                  | string | The total number of digitalbits in circulation.                                                                                   |
| fee_pool                     | string | The sum of all transaction fees *(in digitalbits)* since the last inflation operation. They are redistributed during [inflation]. |
| base_fee                     | number | The [fee] the network charges per operation in a transaction.                                                                |
| base_reserve                 | string | The [reserve][fee] the network uses when calculating an account's minimum balance.                                           |
| max_tx_set_size              | number | The maximum number of transactions validators have agreed to process in a given ledger.                                      |
| protocol_version             | number | The protocol version that the digitalbits network was running when this ledger was committed.                                    |
| header_xdr                   | string | A base64 encoded string of the raw `LedgerHeader` xdr struct for this ledger.                                                |
| base_fee_in_nibbs          | number | The [fee] the network charges per operation in a transaction.  Expressed in nibbs.                                         |
| base_reserve_in_nibbs      | number | The [reserve][fee] the network uses when calculating an account's minimum balance. Expressed in nibbs.                     |

## Links
|              | Example                                           | Relation                        | templated |
|--------------|---------------------------------------------------|---------------------------------|-----------|
| self         | `/ledgers/957773`                                    |                                 |           |
| effects      | `/ledgers/957773/effects/{?cursor,limit,order}`      | The effects in this transaction | true      |
| operations   | `/ledgers/957773/operations/{?cursor,limit,order}`   | The operations in this ledger   | true      |
| transactions | `/ledgers/957773/transactions/{?cursor,limit,order}` | The transactions in this ledger | true      |


## Example

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/ledgers/957773"
    },
    "transactions": {
      "href": "https://frontier.testnet.digitalbits.io/ledgers/957773/transactions{?cursor,limit,order}",
      "templated": true
    },
    "operations": {
      "href": "https://frontier.testnet.digitalbits.io/ledgers/957773/operations{?cursor,limit,order}",
      "templated": true
    },
    "payments": {
      "href": "https://frontier.testnet.digitalbits.io/ledgers/957773/payments{?cursor,limit,order}",
      "templated": true
    },
    "effects": {
      "href": "https://frontier.testnet.digitalbits.io/ledgers/957773/effects{?cursor,limit,order}",
      "templated": true
    }
  },
  "id": "65e18185725ec118092b4915e341c7c75917738b238bee8998deead91946be75",
  "paging_token": "4113603711991808",
  "hash": "65e18185725ec118092b4915e341c7c75917738b238bee8998deead91946be75",
  "prev_hash": "d5c1154590395e887757be1809a48135d2ca8822e505ef50f52065f0cb21c1eb",
  "sequence": 957773,
  "successful_transaction_count": 1,
  "failed_transaction_count": 0,
  "operation_count": 1,
  "tx_set_operation_count": 1,
  "closed_at": "2021-06-16T09:08:23Z",
  "total_coins": "20000000000.0000000",
  "fee_pool": "0.0206000",
  "base_fee_in_nibbs": 100,
  "base_reserve_in_nibbs": 100000000,
  "max_tx_set_size": 100,
  "protocol_version": 15,
  "header_xdr": "AAAAD9XBFUWQOV6Id1e+GAmkgTXSyogi5QXvUPUgZfDLIcHrux5WF1JwRkTjTUEZ93mXRq4N8U6Xp53revUuRqqGiksAAAAAYMm/hwAAAAAAAAABAAAAAKvZnpGcHWYsIfQkvokpnA88t6aedQMkQ3LW/icyV30jAAAAQNqKT73RmwY7exn3h85m8RAlZ57SXrMH/TfYk6Gxvy1owUiHbyL1m1LcWDfjGLY429i3Ppwqb+XW35132vSOFgjTwPMb1bVYrBmvDFs/huBlb3dyFHWxx0guIrpv+rJCRlm1t9zYqfOiTdbG7MebuJxNp/r7H+dD0foh9v12OJ3FAA6dTQLGivC7FAAAAAAAAAADJLAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGSaOd07uLsDMIPf6nvuLuj6ev7+suXe3mfFES+inzTMUbGnoJtIrqPU36tDFU4XORBgsCvIi04GG/A0tVIWclCY2pL7Nkua71s2zrhLvP2xk17wI1QdTs2NbP8p4hUvqO96TzEzCTu1IfbrP9QD0x0cN77mrkt2Hhi4BP6sYQcbDgAAAAA="
}

```

## Endpoints
| Resource                | Type       | Resource URI Template              |
|-------------------------|------------|------------------------------------|
| [All ledgers](../endpoints/ledgers-all.md)         | Collection | `/ledgers`                         |
| [Single Ledger](../endpoints/ledgers-single.md)       | Single     | `/ledgers/:id`                     |
| [Ledger Transactions](../endpoints/transactions-for-ledger.md) | Collection | `/ledgers/:ledger_id/transactions` |
| [Ledger Operations](../endpoints/operations-for-ledger.md)   | Collection | `/ledgers/:ledger_id/operations`   |
| [Ledger Payments](../endpoints/payments-for-ledger.md)     | Collection | `/ledgers/:ledger_id/payments`     |
| [Ledger Effects](../endpoints/effects-for-ledger.md)      | Collection | `/ledgers/:ledger_id/effects`      |



[inflation]: https://developers.digitalbits.io/guides/concepts/inflation.html
[fee]: https://developers.digitalbits.io/guides/concepts/fees.html
