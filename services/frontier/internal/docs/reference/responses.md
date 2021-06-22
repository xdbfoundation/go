Rather than using a fully custom way of representing the resources we expose in
Frontier, we use [HAL](http://stateless.co/hal_specification.html). HAL is a
hypermedia format in JSON that remains simple while giving us a couple of
benefits such as simpler client integration for several languages. See [this
wiki page](https://github.com/mikekelly/hal_specification/wiki/Libraries) for a
list of libraries.

## Attributes, Links, Embedded Resources

At its simplest, a HAL response is just a JSON object with a couple of reserved
property names:  `_links` is used for expressing links and `_embedded` is used
for bundling other HAL objects with the response.  Other than links and embedded
objects, **HAL is just JSON**.

### Links

HAL is a hypermedia format, like HTML, in that it has a mechanism to express
links between documents.  Let's look at a simple example:

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

The above response is for the genesis ledger of the DigitalBits test network, and
the links in the `_links` attribute provide links to other relavant resources in
Frontier.  Notice the object beneath the `transactions` key.  The key of each
link specifies that links relation to the current resource, and in this case
`transactions` means "Transactions that occurred in this ledger".  Logically,
you should expect that resource to respond with a collection of transactions
with all of the results having a `ledger_sequence` attribute equal to 1.

The `transactions` link is also _templated_, which means that the `href`
attribute of the link is actually a URI template, as specified by  [RFC
6570](https://tools.ietf.org/html/rfc6570).  We use URI templates to show you
what parameters a given resource can take. You must evaluate the template to a
valid URI before navigating to it.

## Pages

Pages represent a subset of a larger collection of objects.  
As an example, it would be unfeasible to provide the
[All Transactions](./endpoints/transactions-all.md) endpoint without paging.  
Over time there will be millions of transactions in the DigitalBits network's ledger
and returning them all over a single request would be unfeasible.

Read more about paging in following docs:

- [Page](./resources/page.md)
- [Paging](./paging.md)
