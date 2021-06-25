A **path** resource contains information about a payment path.  A path can be used by code to populate necessary fields on path payment operation, such as `path` and `sendMax`.


## Attributes
| Attribute                | Type             |                                                                                                                                |
|--------------------------|------------------|--------------------------------------------------------------------------------------------------------------------------------|
| path                     | array of objects            | An array of assets that represents the intermediary assets this path hops through                                               |
| source_amount            | string           | An estimated cost for making a payment of destination_amount on this path. Suitable for use in a path payments `sendMax` field |
| destination_amount       | string           | The destination amount specified in the search that found this path                                                            |
| destination_asset_type   | string           | The type for the destination asset specified in the search that found this path                                                |
| destination_asset_code   | optional, string | The code for the destination asset specified in the search that found this path                                                |
| destination_asset_issuer | optional, string | The issuer for the destination asset specified in the search that found this path                                              |
| source_asset_type        | string           | The type for the source asset specified in the search that found this path                                                     |
| source_asset_code        | optional, string | The code for the source asset specified in the search that found this path                                                     |
| source_asset_issuer      | optional, string | The issuer for the source asset specified in the search that found this path                                                   |

#### Asset Object
| Attribute    | Type             |                                                                                                                        |
|--------------|------------------|------------------------------------------------------------------------------------------------------------------------
| asset_code     | optional, string           | The code for the asset.                       |
| asset_type     | string           | Either native, credit_alphanum4, or credit_alphanum12.                        |
| asset_issuer     | optional, string           | The digitalbits address of the given asset's issuer.  |

## Example

```json
[
  {
    "source_asset_type": "native",
    "source_amount": "0.2500000",
    "destination_asset_type": "credit_alphanum4",
    "destination_asset_code": "USD",
    "destination_asset_issuer": "GB4RZUSF3HZGCAKB3VBM2S7QOHHC5KTV3LLZXGBYR5ZO4B26CKHFZTSZ",
    "destination_amount": "1.0000000",
    "path": [
      {
        "asset_type": "credit_alphanum4",
        "asset_code": "UAH",
        "asset_issuer": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK"
      },
      {
        "asset_type": "credit_alphanum4",
        "asset_code": "EUR",
        "asset_issuer": "GDCIQQY2UKVNLLWGIX74DMTEAFCMQKAKYUWPBO7PLTHIHRKSFZN7V2FC"
      }
    ]
  },
  {
    "source_asset_type": "native",
    "source_amount": "1.0000000",
    "destination_asset_type": "credit_alphanum4",
    "destination_asset_code": "USD",
    "destination_asset_issuer": "GB4RZUSF3HZGCAKB3VBM2S7QOHHC5KTV3LLZXGBYR5ZO4B26CKHFZTSZ",
    "destination_amount": "1.0000000",
    "path": []
  }
]
```

## Endpoints
| Resource                                 | Type       | Resource URI Template |
|------------------------------------------|------------|-----------------------|
| [Find Payment Paths](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/path-finding.md) | Collection | `/paths`              |
