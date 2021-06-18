**Transactions** are the basic unit of change in the DigitalBits Network.

A transaction is a grouping of [operations](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/operation).

To learn more about the concept of transactions in the DigitalBits network, take a look at the [DigitalBits transactions concept guide](https://developers.digitalbits.io/guides/docs/guides/concepts/transactions).

## Attributes

| Attribute               | Type                     |                                                                                                                                |
|-------------------------|--------------------------|--------------------------------------------------------------------------------------------------------------------------------|
| id                      | string                   | The canonical id of this transaction, suitable for use as the :id parameter for url templates that require a transaction's ID. |
| paging_token            | string                   | A [paging token](./page.md) suitable for use as the `cursor` parameter to transaction collection resources.                    |
| successful              | bool                     | Indicates if transaction was successful or not.                                                                                |
| hash                    | string                   | A hex-encoded, lowercase SHA-256 hash of the transaction's [XDR](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/xdr)-encoded form.                             |
| ledger                  | number                   | Sequence number of the ledger in which this transaction was applied.                                                           |
| created_at              | ISO8601 string           |                                                                                                                                |
| fee_account             | string                   | The account which paid for the transaction fees                                                                                |
| source_account          | string                   |                                                                                                                                |
| source_account_sequence | string                   |                                                                                                                                |
| max_fee                 | number                   | The the maximum fee the fee account was willing to pay.                                                                        |
| fee_charged             | number                   | The fee paid by the fee account of this transaction when the transaction was applied to the ledger.                            |
| operation_count         | number                   | The number of operations that are contained within this transaction.                                                           |
| envelope_xdr            | string                   | A base64 encoded string of the raw `TransactionEnvelope` xdr struct for this transaction                                       |
| result_xdr              | string                   | A base64 encoded string of the raw `TransactionResult` xdr struct for this transaction                                         |
| result_meta_xdr         | string                   | A base64 encoded string of the raw `TransactionMeta` xdr struct for this transaction                                           |
| fee_meta_xdr            | string                   | A base64 encoded string of the raw `LedgerEntryChanges` xdr struct produced by taking fees for this transaction.               |
| memo_type               | string                   | The type of memo set in the transaction. Possible values are `none`, `text`, `id`, `hash`, and `return`.                       |
| memo                    | string                   | The string representation of the memo set in the transaction. When `memo_type` is `id`, the `memo` is a decimal string representation of an unsigned 64 bit integer. When `memo_type` is `hash` or `return`, the `memo` is a base64 encoded string. When `memo_type` is `text`, the `memo` is a unicode string. However, if the original memo byte sequence in the transaction XDR is not valid unicode, Frontier will replace any invalid byte sequences with the utf-8 replacement character. Note this field is only present when `memo_type` is not `none`. |
| memo_bytes              | string                   | A base64 encoded string of the memo bytes set in the transaction's xdr envelope. Note this field is only present when `memo_type` is `text`. |
| signatures              | string[]                 | An array of signatures used to sign this transaction                                                                           |
| valid_after             | RFC3339 date-time string |                                                                                                                                |
| valid_before            | RFC3339 date-time string |                                                                                                                                |
| fee_bump_transaction    | object                   | This object is only present if the transaction is a fee bump transaction or is wrapped by a fee bump transaction. The object has two fields: `hash` (the hash of the fee bump transaction) and `signatures` (the signatures present in the fee bump transaction envelope)                                                                                                                               |
| inner_transaction       | object                   | This object is only present if the transaction is a fee bump transaction or is wrapped by a fee bump transaction. The object has three fields: `hash` (the hash of the inner transaction wrapped by the fee bump transaction), `max_fee` (the max fee set in the inner transaction), and `signatures` (the signatures present in the inner transaction envelope)                                                                                                                               |

## Links

| rel        | Example                                                                                                                                              | Description                                                                                |
|------------|------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------|
| self       | `https://frontier.testnet.digitalbits.io/transactions/a8b22539d1f62825c527dbdfee8ba8d4faf701126021ccfa33bbe8cb149de9fd`                                  |                                                                                            |
| account    | `https://frontier.testnet.digitalbits.io/accounts/GC3CLEUNQVWY36AHTGGX2NASAPHD6EBQXE63YH2B3PAASLCCIG4ELGTP`                                              | The source [account](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/accounts-single) for this transaction.                          |
| ledger     | `https://frontier.testnet.digitalbits.io/ledgers/257`                                                                                                | The [ledger](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/ledgers-single) in which this transaction was applied.                  |
| operations | `https://frontier.testnet.digitalbits.io/transactions/a8b22539d1f62825c527dbdfee8ba8d4faf701126021ccfa33bbe8cb149de9fd/operations{?cursor,limit,order}"` | [Operations](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/operations-for-transaction) included in this transaction.               |
| effects    | `https://frontier.testnet.digitalbits.io/transactions/a8b22539d1f62825c527dbdfee8ba8d4faf701126021ccfa33bbe8cb149de9fd/effects{?cursor,limit,order}"`    | [Effects](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/effects-for-transaction) which resulted by operations in this transaction. |
| precedes   | `https://frontier.testnet.digitalbits.io/transactions?order=asc&cursor=1103806599168`                                                                | A collection of transactions that occur after this transaction.                            |
| succeeds   | `https://frontier.testnet.digitalbits.io/transactions?order=desc&cursor=1103806599168`                                                               | A collection of transactions that occur before this transaction.                           |

## Example

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/847b33d9e54a8884a9b9c1fd68dc5560e1c61e181155aafc1145e934cc12535d"
    },
    "account": {
      "href": "https://frontier.testnet.digitalbits.io/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY"
    },
    "ledger": {
      "href": "https://frontier.testnet.digitalbits.io/ledgers/957773"
    },
    "operations": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/847b33d9e54a8884a9b9c1fd68dc5560e1c61e181155aafc1145e934cc12535d/operations{?cursor,limit,order}",
      "templated": true
    },
    "effects": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/847b33d9e54a8884a9b9c1fd68dc5560e1c61e181155aafc1145e934cc12535d/effects{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "https://frontier.testnet.digitalbits.io/transactions?order=asc&cursor=4113603711995904"
    },
    "succeeds": {
      "href": "https://frontier.testnet.digitalbits.io/transactions?order=desc&cursor=4113603711995904"
    },
    "transaction": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/847b33d9e54a8884a9b9c1fd68dc5560e1c61e181155aafc1145e934cc12535d"
    }
  },
  "id": "847b33d9e54a8884a9b9c1fd68dc5560e1c61e181155aafc1145e934cc12535d",
  "paging_token": "4113603711995904",
  "successful": true,
  "hash": "847b33d9e54a8884a9b9c1fd68dc5560e1c61e181155aafc1145e934cc12535d",
  "ledger": 957773,
  "created_at": "2021-06-16T09:08:23Z",
  "source_account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
  "source_account_sequence": "4113023891406850",
  "fee_account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
  "fee_charged": "100",
  "max_fee": "100",
  "operation_count": 1,
  "envelope_xdr": "AAAAAgAAAADK462YFeuR6EO+Vst63pspzDtfuUukBAUFfeZKbxjkrQAAAGQADpzGAAAAAgAAAAEAAAAAAAAAAAAAAABgyb/qAAAAAAAAAAEAAAAAAAAABgAAAAFFVVIAAAAAAMSIQxqiqtWuxkX/wbJkAUTIKArFLPC771zOg8VSLlv6AAAAAlQL5AAAAAAAAAAAAW8Y5K0AAABAhpPc6Er5j03kONRLQN+S0kb9RJuB71KaxWFEVq4MZtaA8WaRnOrNuKBo+wrKTzDsoyasz3GhylcaXMVPlhBbDQ==",
  "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAAGAAAAAAAAAAA=",
  "result_meta_xdr": "AAAAAgAAAAIAAAADAA6dTQAAAAAAAAAAyuOtmBXrkehDvlbLet6bKcw7X7lLpAQFBX3mSm8Y5K0AAAAXSHbnOAAOnMYAAAABAAAAAQAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAA6dTQAAAAAAAAAAyuOtmBXrkehDvlbLet6bKcw7X7lLpAQFBX3mSm8Y5K0AAAAXSHbnOAAOnMYAAAACAAAAAQAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAAAwAAAAMADp1NAAAAAAAAAADK462YFeuR6EO+Vst63pspzDtfuUukBAUFfeZKbxjkrQAAABdIduc4AA6cxgAAAAIAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEADp1NAAAAAAAAAADK462YFeuR6EO+Vst63pspzDtfuUukBAUFfeZKbxjkrQAAABdIduc4AA6cxgAAAAIAAAACAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAADp1NAAAAAQAAAADK462YFeuR6EO+Vst63pspzDtfuUukBAUFfeZKbxjkrQAAAAFFVVIAAAAAAMSIQxqiqtWuxkX/wbJkAUTIKArFLPC771zOg8VSLlv6AAAAAAAAAAAAAAACVAvkAAAAAAEAAAAAAAAAAAAAAAA=",
  "fee_meta_xdr": "AAAABAAAAAMADpzmAAAAAAAAAADK462YFeuR6EO+Vst63pspzDtfuUukBAUFfeZKbxjkrQAAABdIduecAA6cxgAAAAEAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEADp1NAAAAAAAAAADK462YFeuR6EO+Vst63pspzDtfuUukBAUFfeZKbxjkrQAAABdIduc4AA6cxgAAAAEAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAMADpzwAAAAAAAAAAC300+A8SGiACMZeKQTbc3s0U6aNTBLD14/5rrFIEl/hAAAAAAAAySwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEADp1NAAAAAAAAAAC300+A8SGiACMZeKQTbc3s0U6aNTBLD14/5rrFIEl/hAAAAAAAAyUUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
  "memo_type": "none",
  "signatures": [
    "hpPc6Er5j03kONRLQN+S0kb9RJuB71KaxWFEVq4MZtaA8WaRnOrNuKBo+wrKTzDsoyasz3GhylcaXMVPlhBbDQ=="
  ],
  "valid_after": "1970-01-01T00:00:00Z",
  "valid_before": "2021-06-16T09:10:02Z"
}

```

## Endpoints

| Resource                                               | Type       | Resource URI Template                |
|--------------------------------------------------------|------------|--------------------------------------|
| [All Transactions](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/transactions-all)             | Collection | `/transactions` (`GET`)              |
| [Post Transaction](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/transactions-create)          | Action     | `/transactions`  (`POST`)            |
| [Transaction Details](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/transactions-single)       | Single     | `/transactions/:id`                  |
| [Account Transactions](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/transactions-for-account) | Collection | `/accounts/:account_id/transactions` |
| [Ledger Transactions](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/transactions-for-ledger)   | Collection | `/ledgers/:ledger_id/transactions`   |


## Submitting transactions
To submit a new transaction to DigitalBits network, it must first be built and signed locally. Then you can submit a hex representation of your transaction’s [XDR](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/xdr) to the `/transactions` endpoint. Read more about submitting transactions in [Post Transaction](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/transactions-create) doc.
