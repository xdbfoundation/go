This endpoint represents all [ledgers](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/ledger).
This endpoint can also be used in [streaming](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/streaming) mode so it is possible to use it to get notifications as ledgers are closed by the DigitalBits network.
If called in streaming mode Frontier will start at the earliest known ledger unless a `cursor` is set. In that case it will start from the `cursor`. You can also set `cursor` value to `now` to only stream ledgers created since your request time.

## Request

```
GET /ledgers{?cursor,limit,order}
```

### Arguments

| name | notes | description | example |
| ---- | ----- | ----------- | ------- |
| `?cursor` | optional, any, default _null_ | A paging token, specifying where to start returning records from. When streaming this can be set to `now` to stream object created since your request time. | `12884905984` |
| `?order`  | optional, string, default `asc` | The order in which to return rows, "asc" or "desc". | `asc` |
| `?limit`  | optional, number, default: `10` | Maximum number of records to return. | `200` |

### curl Example Request

```sh
# Retrieve the 200 latest ledgers, ordered chronologically
curl "https://frontier.testnet.digitalbits.io/ledgers?limit=200&order=desc"
```

### JavaScript Example Request

```javascript
var DigitalBitsSdk = require('digitalbits-sdk')
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

server.ledgers()
  .call()
  .then(function (ledgerResult) {
    // page 1
    console.log(JSON.stringify(ledgerResult.records))
  })
  .catch(function(err) {
    console.log(err)
  })

```


### JavaScript Streaming Example

```javascript
var DigitalBitsSdk = require('digitalbits-sdk')
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

var ledgerHandler = function (ledgerResponse) {
  console.log(ledgerResponse);
};

var es = server.ledgers()
  .cursor('now')
  .stream({
    onmessage: ledgerHandler
})
```

## Response

This endpoint responds with a list of ledgers.  See [ledger resource](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/ledger) for reference.

### Example Response

```json
[
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/192"
      },
      "transactions": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/192/transactions{?cursor,limit,order}",
        "templated": true
      },
      "operations": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/192/operations{?cursor,limit,order}",
        "templated": true
      },
      "payments": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/192/payments{?cursor,limit,order}",
        "templated": true
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/192/effects{?cursor,limit,order}",
        "templated": true
      }
    },
    "id": "97b39c3900e7f30cfca8e9de957969480bf2639d6bbae423433dde84ed6ee8bf",
    "paging_token": "824633720832",
    "hash": "97b39c3900e7f30cfca8e9de957969480bf2639d6bbae423433dde84ed6ee8bf",
    "prev_hash": "e575d362aaaca24584aa71a69b3dde6acfb925e80b7ceb1b8d815699a2c5aff9",
    "sequence": 192,
    "successful_transaction_count": 0,
    "failed_transaction_count": 0,
    "operation_count": 0,
    "tx_set_operation_count": 0,
    "closed_at": "2021-04-13T13:49:26Z",
    "total_coins": "20000000000.0000000",
    "fee_pool": "0.0000000",
    "base_fee_in_nibbs": 100,
    "base_reserve_in_nibbs": 100000000,
    "max_tx_set_size": 100,
    "protocol_version": 15,
    "header_xdr": "AAAAD+V102KqrKJFhKpxpps93mrPuSXoC3zrG42BVpmixa/5jzaKZSZCfJ/sNjMNgnDzXqPoYJkcZ65S0XcyMrno44YAAAAAYHWhZgAAAAAAAAABAAAAAKqV7E8NgJZicfvA5TYMz5Vh5pluOP8ASgCgSd3q3d1NAAAAQPDoYJsZ8CmnZRBiAs7v9C/jV9f+K3QeQnnhvUTfrXukLeviM1jDD9l2Aecp1hcYqd2U+h3SfvSV2Dp1/Gn8zATfP2GYBKkv20BXGS3EPddI6neK3FK8SYzoBSTAFLgRGdUqNRu1cyy6DH9XVFSYcjGSBOXl16wbBekMjr0XoBQuAAAAwALGivC7FAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGRcWfjB1FX6pk2n/C9G5i8Z102VUUTA+VRumlGGMdOMogAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/193"
      },
      "transactions": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/193/transactions{?cursor,limit,order}",
        "templated": true
      },
      "operations": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/193/operations{?cursor,limit,order}",
        "templated": true
      },
      "payments": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/193/payments{?cursor,limit,order}",
        "templated": true
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/193/effects{?cursor,limit,order}",
        "templated": true
      }
    },
    "id": "4b7a21fd6842e03cae999819cfcaecfd5ed37943f2a57bd2c94e932ee89a0e51",
    "paging_token": "828928688128",
    "hash": "4b7a21fd6842e03cae999819cfcaecfd5ed37943f2a57bd2c94e932ee89a0e51",
    "prev_hash": "97b39c3900e7f30cfca8e9de957969480bf2639d6bbae423433dde84ed6ee8bf",
    "sequence": 193,
    "successful_transaction_count": 0,
    "failed_transaction_count": 0,
    "operation_count": 0,
    "tx_set_operation_count": 0,
    "closed_at": "2021-04-13T13:49:31Z",
    "total_coins": "20000000000.0000000",
    "fee_pool": "0.0000000",
    "base_fee_in_nibbs": 100,
    "base_reserve_in_nibbs": 100000000,
    "max_tx_set_size": 100,
    "protocol_version": 15,
    "header_xdr": "AAAAD5eznDkA5/MM/Kjp3pV5aUgL8mOda7rkI0M93oTtbui/gS87y+GxiWBWnRoNrI2VG2ZVP8hBeU3OuBycJi1+YSYAAAAAYHWhawAAAAAAAAABAAAAAKvZnpGcHWYsIfQkvokpnA88t6aedQMkQ3LW/icyV30jAAAAQP0ccgtyAsmeI9IvnFd+e611NT5ymafDtoDPtL+nVNimWlsiBGE/33wZkUi5kISs+/z2K7PAN4J0EIeVBxrOMA3fP2GYBKkv20BXGS3EPddI6neK3FK8SYzoBSTAFLgRGdUqNRu1cyy6DH9XVFSYcjGSBOXl16wbBekMjr0XoBQuAAAAwQLGivC7FAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGRcWfjB1FX6pk2n/C9G5i8Z102VUUTA+VRumlGGMdOMogAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/194"
      },
      "transactions": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/194/transactions{?cursor,limit,order}",
        "templated": true
      },
      "operations": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/194/operations{?cursor,limit,order}",
        "templated": true
      },
      "payments": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/194/payments{?cursor,limit,order}",
        "templated": true
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/194/effects{?cursor,limit,order}",
        "templated": true
      }
    },
    "id": "821008648bc0a7723c66bec3a18048ed206274c6b53914e5fea487753e74c75c",
    "paging_token": "833223655424",
    "hash": "821008648bc0a7723c66bec3a18048ed206274c6b53914e5fea487753e74c75c",
    "prev_hash": "4b7a21fd6842e03cae999819cfcaecfd5ed37943f2a57bd2c94e932ee89a0e51",
    "sequence": 194,
    "successful_transaction_count": 0,
    "failed_transaction_count": 0,
    "operation_count": 0,
    "tx_set_operation_count": 0,
    "closed_at": "2021-04-13T13:49:37Z",
    "total_coins": "20000000000.0000000",
    "fee_pool": "0.0000000",
    "base_fee_in_nibbs": 100,
    "base_reserve_in_nibbs": 100000000,
    "max_tx_set_size": 100,
    "protocol_version": 15,
    "header_xdr": "AAAAD0t6If1oQuA8rpmYGc/K7P1e03lD8qV70slOky7omg5RHnZqtS4uhLSsaeZOMAOo8g+mI0PNbjCt8EiE5IVV/FUAAAAAYHWhcQAAAAAAAAABAAAAAKqV7E8NgJZicfvA5TYMz5Vh5pluOP8ASgCgSd3q3d1NAAAAQJMpEsQ0Qvj2cbxGnQU6WlPeGBCTzJwquIiY46820S/n6Ic93ov1eCuAEQCfigaMX0qMeTWMi8GMyUsEr5cWiQjfP2GYBKkv20BXGS3EPddI6neK3FK8SYzoBSTAFLgRGdUqNRu1cyy6DH9XVFSYcjGSBOXl16wbBekMjr0XoBQuAAAAwgLGivC7FAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGRcWfjB1FX6pk2n/C9G5i8Z102VUUTA+VRumlGGMdOMogAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/195"
      },
      "transactions": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/195/transactions{?cursor,limit,order}",
        "templated": true
      },
      "operations": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/195/operations{?cursor,limit,order}",
        "templated": true
      },
      "payments": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/195/payments{?cursor,limit,order}",
        "templated": true
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/195/effects{?cursor,limit,order}",
        "templated": true
      }
    },
    "id": "72295b3d2d665789792913498a9faf3d36db0e46bc30ffb805503484eccd875d",
    "paging_token": "837518622720",
    "hash": "72295b3d2d665789792913498a9faf3d36db0e46bc30ffb805503484eccd875d",
    "prev_hash": "821008648bc0a7723c66bec3a18048ed206274c6b53914e5fea487753e74c75c",
    "sequence": 195,
    "successful_transaction_count": 0,
    "failed_transaction_count": 0,
    "operation_count": 0,
    "tx_set_operation_count": 0,
    "closed_at": "2021-04-13T13:49:45Z",
    "total_coins": "20000000000.0000000",
    "fee_pool": "0.0000000",
    "base_fee_in_nibbs": 100,
    "base_reserve_in_nibbs": 100000000,
    "max_tx_set_size": 100,
    "protocol_version": 15,
    "header_xdr": "AAAAD4IQCGSLwKdyPGa+w6GASO0gYnTGtTkU5f6kh3U+dMdc9N497LH0Zi2MvkEWwZ0cKVsQC1A1P3+egyjq7INc524AAAAAYHWheQAAAAAAAAABAAAAAKvZnpGcHWYsIfQkvokpnA88t6aedQMkQ3LW/icyV30jAAAAQEzqAH9nxBTQiXapzL359D8TlBsE2kaj81BGWX5ujOo3pNkmt1+o74fH9KAhjnhHFIrU51jcwZCoLAo7rv9lOwnfP2GYBKkv20BXGS3EPddI6neK3FK8SYzoBSTAFLgRGdUqNRu1cyy6DH9XVFSYcjGSBOXl16wbBekMjr0XoBQuAAAAwwLGivC7FAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGRcWfjB1FX6pk2n/C9G5i8Z102VUUTA+VRumlGGMdOMogAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/196"
      },
      "transactions": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/196/transactions{?cursor,limit,order}",
        "templated": true
      },
      "operations": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/196/operations{?cursor,limit,order}",
        "templated": true
      },
      "payments": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/196/payments{?cursor,limit,order}",
        "templated": true
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/196/effects{?cursor,limit,order}",
        "templated": true
      }
    },
    "id": "e68c37978f51a04cce7cb936f6f68c084583b332a4d4d428b7717b1a6c9b2265",
    "paging_token": "841813590016",
    "hash": "e68c37978f51a04cce7cb936f6f68c084583b332a4d4d428b7717b1a6c9b2265",
    "prev_hash": "72295b3d2d665789792913498a9faf3d36db0e46bc30ffb805503484eccd875d",
    "sequence": 196,
    "successful_transaction_count": 0,
    "failed_transaction_count": 0,
    "operation_count": 0,
    "tx_set_operation_count": 0,
    "closed_at": "2021-04-13T13:49:50Z",
    "total_coins": "20000000000.0000000",
    "fee_pool": "0.0000000",
    "base_fee_in_nibbs": 100,
    "base_reserve_in_nibbs": 100000000,
    "max_tx_set_size": 100,
    "protocol_version": 15,
    "header_xdr": "AAAAD3IpWz0tZleJeSkTSYqfrz022w5GvDD/uAVQNITszYddeKzobnKSRAjo5yfp8G+KPKL8HJMmfCBJAyhc88zSwuwAAAAAYHWhfgAAAAAAAAABAAAAAKqV7E8NgJZicfvA5TYMz5Vh5pluOP8ASgCgSd3q3d1NAAAAQNUOlJ9UaI5/VA7XtZEOeXs/cjBw6tQrrQc+bCNZgUp8ehtptwx+oN0MdpkKtmmgAyFRag+y4wLFxOwjB/Zp6w3fP2GYBKkv20BXGS3EPddI6neK3FK8SYzoBSTAFLgRGdUqNRu1cyy6DH9XVFSYcjGSBOXl16wbBekMjr0XoBQuAAAAxALGivC7FAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGRcWfjB1FX6pk2n/C9G5i8Z102VUUTA+VRumlGGMdOMogAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/197"
      },
      "transactions": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/197/transactions{?cursor,limit,order}",
        "templated": true
      },
      "operations": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/197/operations{?cursor,limit,order}",
        "templated": true
      },
      "payments": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/197/payments{?cursor,limit,order}",
        "templated": true
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/197/effects{?cursor,limit,order}",
        "templated": true
      }
    },
    "id": "835ad8719d0eaee306b68d7757a4d302ad2529b8911a7996adbc86783470009d",
    "paging_token": "846108557312",
    "hash": "835ad8719d0eaee306b68d7757a4d302ad2529b8911a7996adbc86783470009d",
    "prev_hash": "e68c37978f51a04cce7cb936f6f68c084583b332a4d4d428b7717b1a6c9b2265",
    "sequence": 197,
    "successful_transaction_count": 0,
    "failed_transaction_count": 0,
    "operation_count": 0,
    "tx_set_operation_count": 0,
    "closed_at": "2021-04-13T13:49:55Z",
    "total_coins": "20000000000.0000000",
    "fee_pool": "0.0000000",
    "base_fee_in_nibbs": 100,
    "base_reserve_in_nibbs": 100000000,
    "max_tx_set_size": 100,
    "protocol_version": 15,
    "header_xdr": "AAAAD+aMN5ePUaBMzny5Nvb2jAhFg7MypNTUKLdxexpsmyJlgisjZdm6CHWyI6O8ZNX/bAu8lx9GV5du4v1k0dLeH8QAAAAAYHWhgwAAAAAAAAABAAAAAPljcPIHxOHkMhXc81cs1HPAkJg+/WTx8514Ln+Dz2C+AAAAQHrdl+gzfS4WXjBcafpHT5OkVkikxmCPdB0hgWkPlJ1r0wJLx9tJrw7VE9C8TKZdQ0t6N3ukzlExW/IurnIH+ArfP2GYBKkv20BXGS3EPddI6neK3FK8SYzoBSTAFLgRGdUqNRu1cyy6DH9XVFSYcjGSBOXl16wbBekMjr0XoBQuAAAAxQLGivC7FAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGRcWfjB1FX6pk2n/C9G5i8Z102VUUTA+VRumlGGMdOMogAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/198"
      },
      "transactions": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/198/transactions{?cursor,limit,order}",
        "templated": true
      },
      "operations": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/198/operations{?cursor,limit,order}",
        "templated": true
      },
      "payments": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/198/payments{?cursor,limit,order}",
        "templated": true
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/198/effects{?cursor,limit,order}",
        "templated": true
      }
    },
    "id": "d0be45a62bcdb2e626dc6cba3afe0aaae8cb58e20d953ac2d2fe55b4ce15bad3",
    "paging_token": "850403524608",
    "hash": "d0be45a62bcdb2e626dc6cba3afe0aaae8cb58e20d953ac2d2fe55b4ce15bad3",
    "prev_hash": "835ad8719d0eaee306b68d7757a4d302ad2529b8911a7996adbc86783470009d",
    "sequence": 198,
    "successful_transaction_count": 0,
    "failed_transaction_count": 0,
    "operation_count": 0,
    "tx_set_operation_count": 0,
    "closed_at": "2021-04-13T13:50:00Z",
    "total_coins": "20000000000.0000000",
    "fee_pool": "0.0000000",
    "base_fee_in_nibbs": 100,
    "base_reserve_in_nibbs": 100000000,
    "max_tx_set_size": 100,
    "protocol_version": 15,
    "header_xdr": "AAAAD4Na2HGdDq7jBraNd1ek0wKtJSm4kRp5lq28hng0cACdM2rhNslxLWKML2RPbSuUfYU7bDO5XzPPj5d04tOMEloAAAAAYHWhiAAAAAAAAAABAAAAAKqV7E8NgJZicfvA5TYMz5Vh5pluOP8ASgCgSd3q3d1NAAAAQLqNe0av1lxPGGC/5gD62f1V9wdMXZ1NGNUQM9gP4mksYKrlzYi1gfOx2TYbZm/KBJALVmfxRLkdYT1So4z5igvfP2GYBKkv20BXGS3EPddI6neK3FK8SYzoBSTAFLgRGdUqNRu1cyy6DH9XVFSYcjGSBOXl16wbBekMjr0XoBQuAAAAxgLGivC7FAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGRcWfjB1FX6pk2n/C9G5i8Z102VUUTA+VRumlGGMdOMogAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/199"
      },
      "transactions": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/199/transactions{?cursor,limit,order}",
        "templated": true
      },
      "operations": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/199/operations{?cursor,limit,order}",
        "templated": true
      },
      "payments": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/199/payments{?cursor,limit,order}",
        "templated": true
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/199/effects{?cursor,limit,order}",
        "templated": true
      }
    },
    "id": "959046e369f292470e233b5d6df67e5676edaa5581bca6f5aff04375e3b28cd7",
    "paging_token": "854698491904",
    "hash": "959046e369f292470e233b5d6df67e5676edaa5581bca6f5aff04375e3b28cd7",
    "prev_hash": "d0be45a62bcdb2e626dc6cba3afe0aaae8cb58e20d953ac2d2fe55b4ce15bad3",
    "sequence": 199,
    "successful_transaction_count": 0,
    "failed_transaction_count": 0,
    "operation_count": 0,
    "tx_set_operation_count": 0,
    "closed_at": "2021-04-13T13:50:09Z",
    "total_coins": "20000000000.0000000",
    "fee_pool": "0.0000000",
    "base_fee_in_nibbs": 100,
    "base_reserve_in_nibbs": 100000000,
    "max_tx_set_size": 100,
    "protocol_version": 15,
    "header_xdr": "AAAAD9C+RaYrzbLmJtxsujr+Cqroy1jiDZU6wtL+VbTOFbrTAQ6foG7hR5iXHKszJg+8cdl7NMZXq311lZxATKrwL/AAAAAAYHWhkQAAAAAAAAABAAAAAPljcPIHxOHkMhXc81cs1HPAkJg+/WTx8514Ln+Dz2C+AAAAQAc0WqFgDWbBeZ0x7BMUcHk5JaKPkH+8jAsXwtNDKnvQzZqmBxTIvpL6C6VDFys/5U7XujeDZ7bnJCHFWoYt/AbfP2GYBKkv20BXGS3EPddI6neK3FK8SYzoBSTAFLgRGdUqNRu1cyy6DH9XVFSYcjGSBOXl16wbBekMjr0XoBQuAAAAxwLGivC7FAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGRcWfjB1FX6pk2n/C9G5i8Z102VUUTA+VRumlGGMdOMogAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/200"
      },
      "transactions": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/200/transactions{?cursor,limit,order}",
        "templated": true
      },
      "operations": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/200/operations{?cursor,limit,order}",
        "templated": true
      },
      "payments": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/200/payments{?cursor,limit,order}",
        "templated": true
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/200/effects{?cursor,limit,order}",
        "templated": true
      }
    },
    "id": "47be64cea8c7f7d7c1eb47525d3d6bf514f62dbe3f4086d2f15d5d4a918c9752",
    "paging_token": "858993459200",
    "hash": "47be64cea8c7f7d7c1eb47525d3d6bf514f62dbe3f4086d2f15d5d4a918c9752",
    "prev_hash": "959046e369f292470e233b5d6df67e5676edaa5581bca6f5aff04375e3b28cd7",
    "sequence": 200,
    "successful_transaction_count": 0,
    "failed_transaction_count": 0,
    "operation_count": 0,
    "tx_set_operation_count": 0,
    "closed_at": "2021-04-13T13:50:14Z",
    "total_coins": "20000000000.0000000",
    "fee_pool": "0.0000000",
    "base_fee_in_nibbs": 100,
    "base_reserve_in_nibbs": 100000000,
    "max_tx_set_size": 100,
    "protocol_version": 15,
    "header_xdr": "AAAAD5WQRuNp8pJHDiM7XW32flZ27apVgbym9a/wQ3XjsozXWUqYq647Gy/SWC5clOa/YsqJmuDVkzoSqYKLPvobnysAAAAAYHWhlgAAAAAAAAABAAAAAPljcPIHxOHkMhXc81cs1HPAkJg+/WTx8514Ln+Dz2C+AAAAQG/K82a1NaNPPLuTRQNAJFOR4QKpeiCXUXNBQ37BzYKcuzI1ds4DewSWEVd5R/lSqruO3BJFkAhDH9Cq1vlqtAnfP2GYBKkv20BXGS3EPddI6neK3FK8SYzoBSTAFLgRGdUqNRu1cyy6DH9XVFSYcjGSBOXl16wbBekMjr0XoBQuAAAAyALGivC7FAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGTVKjUbtXMsugx/V1RUmHIxkgTl5desGwXpDI69F6AULgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/201"
      },
      "transactions": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/201/transactions{?cursor,limit,order}",
        "templated": true
      },
      "operations": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/201/operations{?cursor,limit,order}",
        "templated": true
      },
      "payments": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/201/payments{?cursor,limit,order}",
        "templated": true
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/ledgers/201/effects{?cursor,limit,order}",
        "templated": true
      }
    },
    "id": "3c1ea240afe5abba9a282da4d16c6ecd38616b27ddd0d020b414fbdfd542a8f6",
    "paging_token": "863288426496",
    "hash": "3c1ea240afe5abba9a282da4d16c6ecd38616b27ddd0d020b414fbdfd542a8f6",
    "prev_hash": "47be64cea8c7f7d7c1eb47525d3d6bf514f62dbe3f4086d2f15d5d4a918c9752",
    "sequence": 201,
    "successful_transaction_count": 0,
    "failed_transaction_count": 0,
    "operation_count": 0,
    "tx_set_operation_count": 0,
    "closed_at": "2021-04-13T13:50:20Z",
    "total_coins": "20000000000.0000000",
    "fee_pool": "0.0000000",
    "base_fee_in_nibbs": 100,
    "base_reserve_in_nibbs": 100000000,
    "max_tx_set_size": 100,
    "protocol_version": 15,
    "header_xdr": "AAAAD0e+ZM6ox/fXwetHUl09a/UU9i2+P0CG0vFdXUqRjJdSZN0faxhiYodqc9MDeSfurghqO9Y7bHAiAedtV/xzSMwAAAAAYHWhnAAAAAAAAAABAAAAAPljcPIHxOHkMhXc81cs1HPAkJg+/WTx8514Ln+Dz2C+AAAAQHh0anzYnaZ2hBeUbo49y6VdraQQTuO+8n3xuqWJBll/AJx9ERgRRZkjbXGzs5FCNmr378xjNk0CFZ5n8MdwewDfP2GYBKkv20BXGS3EPddI6neK3FK8SYzoBSTAFLgRGdUqNRu1cyy6DH9XVFSYcjGSBOXl16wbBekMjr0XoBQuAAAAyQLGivC7FAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGTVKjUbtXMsugx/V1RUmHIxkgTl5desGwXpDI69F6AULgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
  }
]

```

### Example Streaming Event

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/ledgers/958657"
    },
    "transactions": {
      "href": "https://frontier.testnet.digitalbits.io/ledgers/958657/transactions{?cursor,limit,order}",
      "templated": true
    },
    "operations": {
      "href": "https://frontier.testnet.digitalbits.io/ledgers/958657/operations{?cursor,limit,order}",
      "templated": true
    },
    "payments": {
      "href": "https://frontier.testnet.digitalbits.io/ledgers/958657/payments{?cursor,limit,order}",
      "templated": true
    },
    "effects": {
      "href": "https://frontier.testnet.digitalbits.io/ledgers/958657/effects{?cursor,limit,order}",
      "templated": true
    }
  },
  "id": "d2b742d31acf6af9e41da960e74199da8f128e88ca55aac6cded14cf5fcec566",
  "paging_token": "4117400463081472",
  "hash": "d2b742d31acf6af9e41da960e74199da8f128e88ca55aac6cded14cf5fcec566",
  "prev_hash": "0430688a5394cbd3bdb20ace3fd46371df9821c4ebfc96c333783c26abfb2476",
  "sequence": 958657,
  "successful_transaction_count": 0,
  "failed_transaction_count": 0,
  "operation_count": 0,
  "tx_set_operation_count": 0,
  "closed_at": "2021-06-16T10:33:00Z",
  "total_coins": "20000000000.0000000",
  "fee_pool": "0.0207200",
  "base_fee_in_nibbs": 100,
  "base_reserve_in_nibbs": 100000000,
  "max_tx_set_size": 100,
  "protocol_version": 15,
  "header_xdr": "AAAADwQwaIpTlMvTvbIKzj/UY3HfmCHE6/yWwzN4PCar+yR28aQM7p8RkbDHmZP9TX4tU51fI3fmhaGp6f+ttaQC8jMAAAAAYMnTXAAAAAAAAAABAAAAAKvZnpGcHWYsIfQkvokpnA88t6aedQMkQ3LW/icyV30jAAAAQFYZQ3XhVMuHn6Myb3k6+8tklb2h3K1GbJF0KAlQTJJUG1isziJPXNnuy/dlujVhSTSU7vu/j1d374WfhzGnkgzfP2GYBKkv20BXGS3EPddI6neK3FK8SYzoBSTAFLgRGdFxiUWEuBYBeP+5fiimEr9ZxcFFuRh8iTRJ4wYjCJkTAA6gwQLGivC7FAAAAAAAAAADKWAAAAAAAAAAAAAAAAAAAABkBfXhAAAAAGTRcYlFhLgWAXj/uX4ophK/WcXBRbkYfIk0SeMGIwiZE7GnoJtIrqPU36tDFU4XORBgsCvIi04GG/A0tVIWclCY2pL7Nkua71s2zrhLvP2xk17wI1QdTs2NbP8p4hUvqO96TzEzCTu1IfbrP9QD0x0cN77mrkt2Hhi4BP6sYQcbDgAAAAA="
}
```

## Errors

- The [standard errors](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors#standard-errors).
