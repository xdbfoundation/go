**Assets** are the units that are traded on the DigitalBits Network.

An asset consists of an type, code, and issuer.

To learn more about the concept of assets in the DigitalBits network, take a look at the [DigitalBits assets concept guide](https://developers.digitalbits.io/guides/docs/guides/concepts/assets).

## Attributes

|    Attribute     |  Type  |                                                                                                                                |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------ |
| asset_type               | string | The type of this asset: "credit_alphanum4", or "credit_alphanum12". |
| asset_code               | string | The code of this asset.   |
| asset_issuer             | string | The issuer of this asset. |
| amount                   | number | The number of units of credit issued. |
| num_accounts             | number | The number of accounts that: 1) trust this asset and 2) where if the asset has the auth_required flag then the account is authorized to hold the asset. |
| flags                    | object | The flags denote the enabling/disabling of certain asset issuer privileges. |
| paging_token             | string | A [paging token](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/page) suitable for use as the `cursor` parameter to transaction collection resources.                   |

#### Flag Object
|    Attribute     |  Type  |                                                                                                                                |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------ |
| auth_immutable             | bool | With this setting, none of the following authorization flags can be changed. |
| auth_required              | bool | With this setting, an anchor must approve anyone who wants to hold its asset.  |
| auth_revocable             | bool | With this setting, an anchor can set the authorize flag of an existing trustline to freeze the assets held by an asset holder.  |

## Links
| rel          | Example                                                                                           | Description                                                
|--------------|---------------------------------------------------------------------------------------------------|------------------------------------------------------------
| toml  | `https://livenet.digitalbits.io/.well-known/digitalbits.toml`| Link to the TOML file for this issuer |

## Example

```json
{
    "_links": {
      "toml": {
        "href": ""
      }
    },
    "asset_type": "credit_alphanum4",
    "asset_code": "EUR",
    "asset_issuer": "GDCIQQY2UKVNLLWGIX74DMTEAFCMQKAKYUWPBO7PLTHIHRKSFZN7V2FC",
    "paging_token": "EUR_GDCIQQY2UKVNLLWGIX74DMTEAFCMQKAKYUWPBO7PLTHIHRKSFZN7V2FC_credit_alphanum4",
    "amount": "15.0000000",
    "num_accounts": 2,
    "flags": {
      "auth_required": false,
      "auth_revocable": false,
      "auth_immutable": false
    }
}
```

## Endpoints

|  Resource                                |    Type    |    Resource URI Template     |
| ---------------------------------------- | ---------- | ---------------------------- |
| [All Assets](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/assets-all) | Collection | `/assets` (`GET`)            |
