Pages represent a subset of a larger collection of objects.
As an example, it would be unfeasible to provide the
[All Transactions](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/transactions-all) endpoint without paging.  Over time there
will be millions of transactions in the DigitalBits network's ledger and returning
them all over a single request would be unfeasible.

## Attributes

A page itself exposes no attributes.  It is merely a container for embedded
records and some links to aid in iterating the entire collection the page is
part of.

## Cursor
A `cursor` is a number that points to a specific location in a collection of resources.

The `cursor` attribute itself is an opaque value meaning that users should not try to parse it.

## Embedded Resources

A page contains an embedded set of `records`, regardless of the contained resource.

## Links

A page provides a couple of links to ease in iteration.

|      |                        Example                         |           Relation           |
| ---- | ------------------------------------------------------ | ---------------------------- |
| self | `/transactions`                                        |                              |
| prev | `/transactions?cursor=12884905984&order=desc&limit=10` | The previous page of results |
| next | `/transactions?cursor=12884905984&order=asc&limit=10`  | The next page of results     |

## Example

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/operations/1099511631874"
    },
    "transaction": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/081c8114fe004413a325294413c9372ce47ac4fc6925b5b994d80f854e0bddf9"
    },
    "effects": {
      "href": "https://frontier.testnet.digitalbits.io/operations/1099511631874/effects"
    },
    "succeeds": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=1099511631874"
    },
    "precedes": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=1099511631874"
    }
  },
  "id": "1099511631874",
  "paging_token": "1099511631874",
  "transaction_successful": true,
  "source_account": "GC3CLEUNQVWY36AHTGGX2NASAPHD6EBQXE63YH2B3PAASLCCIG4ELGTP",
  "type": "create_account",
  "type_i": 0,
  "created_at": "2021-04-13T13:55:32Z",
  "transaction_hash": "081c8114fe004413a325294413c9372ce47ac4fc6925b5b994d80f854e0bddf9",
  "starting_balance": "101.0000000",
  "funder": "GC3CLEUNQVWY36AHTGGX2NASAPHD6EBQXE63YH2B3PAASLCCIG4ELGTP",
  "account": "GBPULLXKNDHPAP25N66JA4SH5SOQSNAIWKPVFTATMY6DDV43GBH2TUGV"
}


```

## Endpoints

Any endpoint that provides a collection of resources will represent them as pages.

