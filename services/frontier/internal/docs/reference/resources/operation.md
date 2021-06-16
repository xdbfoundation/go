[Operations](https://developers.digitalbits.io/guides/docs/guides/concepts/operations) are objects that represent a desired change to the ledger: payments,
offers to exchange currency, changes made to account options, etc.  Operations
are submitted to the DigitalBits network grouped in a [Transaction](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/transaction).

To learn more about the concept of operations in the DigitalBits network, take a look at the [DigitalBits operations concept guide](https://developers.digitalbits.io/guides/docs/guides/concepts/operations).

## Operation Types

| type                                                    | type_i | description                                                                                                |
|---------------------------------------------------------|--------|------------------------------------------------------------------------------------------------------------|
| [CREATE_ACCOUNT](#create-account)                       | 0      | Creates a new account in DigitalBits network.                                                                  |
| [PAYMENT](#payment)                                     | 1      | Sends a simple payment between two accounts in DigitalBits network.                                            |
| [PATH_PAYMENT_STRICT_RECEIVE](#path-payment)            | 2      | Sends a path payment strict receive between two accounts in the DigitalBits network.                           |
| [PATH_PAYMENT_STRICT_SEND](#path-payment-strict-send)   | 13     | Sends a path payment strict send between two accounts in the DigitalBits network.                              |
| [MANAGE_SELL_OFFER](#manage-sell-offer)                 | 3      | Creates, updates or deletes a sell offer in the DigitalBits network.                                           |
| [MANAGE_BUY_OFFER](#manage-buy-offer)                   | 12     | Creates, updates or deletes a buy offer in the DigitalBits network.                                            |
| [CREATE_PASSIVE_SELL_OFFER](#create-passive-sell-offer) | 4      | Creates an offer that won't consume a counter offer that exactly matches this offer.                       |
| [SET_OPTIONS](#set-options)                             | 5      | Sets account options (inflation destination, adding signers, etc.)                                         |
| [CHANGE_TRUST](#change-trust)                           | 6      | Creates, updates or deletes a trust line.                                                                  |
| [ALLOW_TRUST](#allow-trust)                             | 7      | Updates the "authorized" flag of an existing trust line this is called by the issuer of the related asset. |
| [ACCOUNT_MERGE](#account-merge)                         | 8      | Deletes account and transfers remaining balance to destination account.                                    |
| [INFLATION](#inflation)                                 | 9      | Runs inflation.                                                                                            |
| [MANAGE_DATA](#manage-data)                             | 10     | Set, modify or delete a Data Entry (name/value pair) for an account.                                       |
| [BUMP_SEQUENCE](#bump-sequence)                         | 11     | Bumps forward the sequence number of an account.                                                           |


Every operation type shares a set of common attributes and links, some operations also contain
additional attributes and links specific to that operation type.



## Common Attributes

|                        | Type   |                                                                                                                             |
|------------------------|--------|-----------------------------------------------------------------------------------------------------------------------------|
| id                     | number | The canonical id of this operation, suitable for use as the :id parameter for url templates that require an operation's ID. |
| paging_token           | any    | A [paging token](./page.md) suitable for use as a `cursor` parameter.                                                       |
| transaction_successful | bool   | Indicates if this operation is part of successful transaction.                                                              |
| type                   | string | A string representation of the type of operation.                                                                           |
| type_i                 | number | Specifies the type of operation, See "Types" section below for reference.                                                   |

## Common Links

|             | Relation                                                                  |
|-------------|---------------------------------------------------------------------------|
| self        | Relative link to the current operation                                    |
| succeeds    | Relative link to the list of operations succeeding the current operation. |
| precedes    | Relative link to the list of operations preceding the current operation.  |
| effects     | The effects this operation triggered                                      |
| transaction | The transaction this operation is part of                                 |


Each operation type will have a different set of attributes, in addition to the
common attributes listed above.

<a id="create-account"></a>
### Create Account

Create Account operation represents a new account creation.

#### Attributes

| Field            | Type   | Description                        |
|------------------|--------|------------------------------------|
| account          | string | A new account that was funded.     |
| funder           | string | Account that funded a new account. |
| starting_balance | string | Amount the account was funded.     |


#### Example
```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/a94ea88f6de5d4a9018c536096921879ff04053d6a865b3810b548ef7854bfae/operations?cursor=&limit=10&order=asc"
    },
    "next": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/a94ea88f6de5d4a9018c536096921879ff04053d6a865b3810b548ef7854bfae/operations?cursor=4126119246696449&limit=10&order=asc"
    },
    "prev": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/a94ea88f6de5d4a9018c536096921879ff04053d6a865b3810b548ef7854bfae/operations?cursor=4126119246696449&limit=10&order=desc"
    }
  },
  "_embedded": {
    "records": [
      {
        "_links": {
          "self": {
            "href": "https://frontier.testnet.digitalbits.io/operations/4126119246696449"
          },
          "transaction": {
            "href": "https://frontier.testnet.digitalbits.io/transactions/a94ea88f6de5d4a9018c536096921879ff04053d6a865b3810b548ef7854bfae"
          },
          "effects": {
            "href": "https://frontier.testnet.digitalbits.io/operations/4126119246696449/effects"
          },
          "succeeds": {
            "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4126119246696449"
          },
          "precedes": {
            "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4126119246696449"
          }
        },
        "id": "4126119246696449",
        "paging_token": "4126119246696449",
        "transaction_successful": true,
        "source_account": "GC3CLEUNQVWY36AHTGGX2NASAPHD6EBQXE63YH2B3PAASLCCIG4ELGTP",
        "type": "create_account",
        "type_i": 0,
        "created_at": "2021-06-16T13:46:30Z",
        "transaction_hash": "a94ea88f6de5d4a9018c536096921879ff04053d6a865b3810b548ef7854bfae",
        "starting_balance": "10000.0000000",
        "funder": "GC3CLEUNQVWY36AHTGGX2NASAPHD6EBQXE63YH2B3PAASLCCIG4ELGTP",
        "account": "GDZTA2VF4L3PKQYGKZO3OM4NNHD2WIPWFGRXHJTVOPQWITQ7PJEKBDD5"
      }
    ]
  }
}
```

<a id="payment"></a>
### Payment

A payment operation represents a payment from one account to another.  This payment
can be either a simple native asset payment or a fiat asset payment.

#### Attributes

| Field        | Type   | Description                                  |
|--------------|--------|----------------------------------------------|
| from         | string | Sender of a payment.                         |
| to           | string | Destination of a payment.                    |
| asset_type   | string | Asset type (native / alphanum4 / alphanum12) |
| asset_code   | string | Code of the destination asset.               |
| asset_issuer | string | Asset issuer.                                |
| amount       | string | Amount sent.                                 |

#### Links

|          | Example                                                            | Relation          |
|----------|--------------------------------------------------------------------|-------------------|
| sender   | `/accounts/GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK` | Sending account   |
| receiver | `/accounts/GCKY3VKRJDSRORRMHRDHA6IKRXMGSBRZE42P64AHX4NHVGB3Y224WM3M` | Receiving account |

#### Example

```json
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4109381759143937"
      },
      "transaction": {
        "href": "https://frontier.testnet.digitalbits.io/transactions/f50342ab33a932dceb23bb44cbc03a13515563a07c51d3e76ecb430163345bec"
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4109381759143937/effects"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4109381759143937"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4109381759143937"
      }
    },
    "id": "4109381759143937",
    "paging_token": "4109381759143937",
    "transaction_successful": true,
    "source_account": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "type": "payment",
    "type_i": 1,
    "created_at": "2021-06-16T07:35:18Z",
    "transaction_hash": "f50342ab33a932dceb23bb44cbc03a13515563a07c51d3e76ecb430163345bec",
    "asset_type": "credit_alphanum4",
    "asset_code": "UAH",
    "asset_issuer": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "from": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "to": "GCKY3VKRJDSRORRMHRDHA6IKRXMGSBRZE42P64AHX4NHVGB3Y224WM3M",
    "amount": "10.0000000"
  }
```

<a id="path-payment"></a>
### Path Payment Strict Receive

A path payment strict receive operation represents a payment from one account to another through a path.  This type of payment starts as one type of asset and ends as another type of asset. There can be other assets that are traded into and out of along the path.


#### Attributes

| Field               | Type                          | Description                                                                 |
|---------------------|-------------------------------|-----------------------------------------------------------------------------|
| from                | string                        | Sender of a payment.                                                        |
| to                  | string                        | Destination of a payment.                                                   |
| asset_code          | string                        | Code of the destination asset.                                              |
| asset_issuer        | string                        | Destination asset issuer.                                                   |
| asset_type          | string                        | Destination asset type (native / alphanum4 / alphanum12)                    |
| amount              | string                        | Amount received.                                                            |
| source_asset_code   | string                        | Code of the source asset.                                                   |
| source_asset_issuer | string                        | Source asset issuer.                                                        |
| source_asset_type   | string                        | Source asset type (native / alphanum4 / alphanum12)                         |
| source_max          | string                        | Max send amount.                                                            |
| source_amount       | string                        | Amount sent.                                                                |
| path                | array of [Assets](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/asset) | Additional hops the operation went through to get to the destination asset. |

#### Example

```json
{
  "_links": {
    "effects": {
      "href": "/operations/25769807873/effects/{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "/operations?cursor=25769807873\u0026order=asc"
    },
    "self": {
      "href": "/operations/25769807873"
    },
    "succeeds": {
      "href": "/operations?cursor=25769807873\u0026order=desc"
    },
    "transaction": {
      "href": "/transactions/25769807872"
    }
  },
  "amount": "10.0",
  "asset_code": "EUR",
  "asset_issuer": "GCQPYGH4K57XBDENKKX55KDTWOTK5WDWRQOH2LHEDX3EKVIQRLMESGBG",
  "asset_type": "credit_alphanum4",
  "from": "GCXKG6RN4ONIEPCMNFB732A436Z5PNDSRLGWK7GBLCMQLIFO4S7EYWVU",
  "id": "25769807873",
  "paging_token": "25769807873",
  "source_asset_code": "USD",
  "source_asset_issuer": "GC23QF2HUE52AMXUFUH3AYJAXXGXXV2VHXYYR6EYXETPKDXZSAW67XO4",
  "source_asset_type": "credit_alphanum4",
  "source_amount": "10.0",
  "source_max": "10.0",
  "to": "GA5WBPYA5Y4WAEHXWR2UKO2UO4BUGHUQ74EUPKON2QHV4WRHOIRNKKH2",
  "transaction_successful": true,
  "type_i": 2,
  "type": "path_payment_strict_receive"
}
```

<a id="path-payment-strict-send"></a>
### Path Payment Strict Send

A path payment strict send operation represents a payment from one account to another through a path.  This type of payment starts as one type of asset and ends as another type of asset. There can be other assets that are traded into and out of along the path.

Unlike [path payment strict receive](#path-payment), this operation sends precisely the source amount, ensuring that the destination account receives at least the minimum specified amount (the amount received will vary based on offers in the order books).


#### Attributes

| Field               | Type                          | Description                                                                 |
|---------------------|-------------------------------|-----------------------------------------------------------------------------|
| from                | string                        | Sender of a payment.                                                        |
| to                  | string                        | Destination of a payment.                                                   |
| asset_type          | string                        | Destination asset type (native / alphanum4 / alphanum12)                    |
| asset_code          | string                        | Code of the destination asset.                                              |
| asset_issuer        | string                        | Destination asset issuer.                                                   |
| amount              | string                        | Amount received.                                                            |
| source_asset_type   | string                        | Source asset type (native / alphanum4 / alphanum12)                         |
| source_asset_code   | string                        | Source asset code.                                                          |
| source_asset_issuer | string                        | Source asset issuer.                                                        |
| source_amount       | string                        | Amount sent.                                                                |
| destination_min     | string                        | The minimum amount of destination asset expected to be received.            |
| path                | array of [Assets](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/asset) | Additional hops the operation went through to get to the destination asset. |


#### Example

```json
{
  "_links": {
    "self": {
      "href": "/operations/120903307907612673"
    },
    "transaction": {
      "href": "/transactions/f60f32eff7f1dd0649cfe2986955d12f6ff45288357fe1526600642ea1b418aa"
    },
    "effects": {
      "href": "/operations/120903307907612673/effects"
    },
    "succeeds": {
      "href": "/effects?order=desc&cursor=120903307907612673"
    },
    "precedes": {
      "href": "/effects?order=asc&cursor=120903307907612673"
    }
  },
  "id": "120903307907612673",
  "paging_token": "120903307907612673",
  "transaction_successful": true,
  "source_account": "GCXVEEBWI4YMRK6AFJQSEUBYDQL4PZ24ECAPJE2ZIAPIQZLZIBAX3LIF",
  "type": "path_payment_strict_send",
  "type_i": 13,
  "created_at": "2020-02-09T20:32:53Z",
  "transaction_hash": "f60f32eff7f1dd0649cfe2986955d12f6ff45288357fe1526600642ea1b418aa",
  "asset_type": "native",
  "from": "GCXVEEBWI4YMRK6AFJQSEUBYDQL4PZ24ECAPJE2ZIAPIQZLZIBAX3LIF",
  "to": "GCXVEEBWI4YMRK6AFJQSEUBYDQL4PZ24ECAPJE2ZIAPIQZLZIBAX3LIF",
  "amount": "0.0859000",
  "path": [

  ],
  "source_amount": "1000.0000000",
  "destination_min": "0.0859000",
  "source_asset_type": "credit_alphanum4",
  "source_asset_code": "KIN",
  "source_asset_issuer": "GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR"
}
```

<a id="manage-sell-offer"></a>
### Manage Sell Offer

A "Manage Sell Offer" operation can create, update or delete a sell
offer to trade assets in the DigitalBits network.
It specifies an issuer, a price and amount of a given asset to
buy or sell.

When this operation is applied to the ledger, trades can potentially be executed if
this offer crosses others that already exist in the ledger.

In the event that there are not enough crossing orders to fill the order completely
a new "Offer" object will be created in the ledger.  As other accounts make
offers or payments, this offer can potentially be filled.

#### Sell Offer vs. Buy Offer

A [sell offer](#manage-sell-offer) specifies a certain amount of the `selling` asset that should be sold in exchange for the maximum quantity of the `buying` asset. It additionally only crosses offers where the price is higher than `price`.

A [buy offer](#manage-buy-offer) specifies a certain amount of the `buying` asset that should be bought in exchange for the minimum quantity of the `selling` asset. It additionally only crosses offers where the price is lower than `price`.

Both will fill only partially (or not at all) if there are few (or no) offers that cross them.

#### Attributes

| Field                | Type   | Description                                             |
|----------------------|--------|---------------------------------------------------------|
| offer_id             | string | Offer ID.                                               |
| amount               | string | Amount of asset to be sold.                             |
| buying_asset_code    | string | The code of asset to buy.                               |
| buying_asset_issuer  | string | The issuer of asset to buy.                             |
| buying_asset_type    | string | Type of asset to buy (native / alphanum4 / alphanum12)  |
| price                | string | Price of selling_asset in units of buying_asset         |
| price_r              | Object | n: price numerator, d: price denominator                |
| selling_asset_code   | string | The code of asset to sell.                              |
| selling_asset_issuer | string | The issuer of asset to sell.                            |
| selling_asset_type   | string | Type of asset to sell (native / alphanum4 / alphanum12) |

#### Example

```json
{
  "_links": {
    "effects": {
      "href": "/operations/592323234762753/effects{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "/operations?cursor=592323234762753\u0026order=asc"
    },
    "self": {
      "href": "/operations/592323234762753"
    },
    "succeeds": {
      "href": "/operations?cursor=592323234762753\u0026order=desc"
    },
    "transaction": {
      "href": "/transactions/592323234762752"
    }
  },
  "amount": "100.0",
  "buying_asset_code": "CHP",
  "buying_asset_issuer": "GAC2ZUXVI5266NMMGDPBMXHH4BTZKJ7MMTGXRZGX2R5YLMFRYLJ7U5EA",
  "buying_asset_type": "credit_alphanum4",
  "id": "592323234762753",
  "offer_id": "8",
  "paging_token": "592323234762753",
  "price": "2.0",
  "price_r": {
    "d": 1,
    "n": 2
  },
  "selling_asset_code": "YEN",
  "selling_asset_issuer": "GDVXG2FMFFSUMMMBIUEMWPZAIU2FNCH7QNGJMWRXRD6K5FZK5KJS4DDR",
  "selling_asset_type": "credit_alphanum4",
  "transaction_successful": true,
  "type_i": 3,
  "type": "manage_sell_offer"
}
```

<a id="manage-buy-offer"></a>
### Manage Buy Offer

A "Manage Buy Offer" operation can create, update or delete a buy
offer to trade assets in the DigitalBits network.
It specifies an issuer, a price and amount of a given asset to
buy or sell.

When this operation is applied to the ledger, trades can potentially be executed if
this offer crosses others that already exist in the ledger.

In the event that there are not enough crossing orders to fill the order completely
a new "Offer" object will be created in the ledger.  As other accounts make
offers or payments, this offer can potentially be filled.

#### Attributes

| Field                | Type   | Description                                                   |
|----------------------|--------|---------------------------------------------------------------|
| offer_id             | string | Offer ID.                                                     |
| buy_amount           | string | Amount of asset to be bought.                                 |
| buying_asset_code    | string | The code of asset to buy.                                     |
| buying_asset_issuer  | string | The issuer of asset to buy.                                   |
| buying_asset_type    | string | Type of asset to buy (native / alphanum4 / alphanum12)        |
| price                | string | Price of thing being bought in terms of what you are selling. |
| price_r              | Object | n: price numerator, d: price denominator                      |
| selling_asset_code   | string | The code of asset to sell.                                    |
| selling_asset_issuer | string | The issuer of asset to sell.                                  |
| selling_asset_type   | string | Type of asset to sell (native / alphanum4 / alphanum12)       |
#### Example

```json
{
  "_links": {
    "effects": {
      "href": "/operations/592323234762753/effects{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "/operations?cursor=592323234762753\u0026order=asc"
    },
    "self": {
      "href": "/operations/592323234762753"
    },
    "succeeds": {
      "href": "/operations?cursor=592323234762753\u0026order=desc"
    },
    "transaction": {
      "href": "/transactions/592323234762752"
    }
  },
  "amount": "100.0",
  "buying_asset_code": "CHP",
  "buying_asset_issuer": "GAC2ZUXVI5266NMMGDPBMXHH4BTZKJ7MMTGXRZGX2R5YLMFRYLJ7U5EA",
  "buying_asset_type": "credit_alphanum4",
  "id": "592323234762753",
  "offer_id": "8",
  "paging_token": "592323234762753",
  "price": "2.0",
  "price_r": {
    "d": 1,
    "n": 2
  },
  "selling_asset_code": "YEN",
  "selling_asset_issuer": "GDVXG2FMFFSUMMMBIUEMWPZAIU2FNCH7QNGJMWRXRD6K5FZK5KJS4DDR",
  "selling_asset_type": "credit_alphanum4",
  "transaction_successful": true,
  "type_i": 12,
  "type": "manage_buy_offer"
}
```

<a id="create-passive-sell-offer"></a>
### Create Passive Sell Offer

“Create Passive Sell Offer” operation creates an offer that won't consume a counter offer that exactly matches this offer. This is useful for offers just used as 1:1 exchanges for path payments. Use Manage Sell Offer to manage this offer after using this operation to create it.

#### Attributes

As in [Manage Sell Offer](#manage-sell-offer) operation.

#### Example

```json
{
  "_links": {
    "effects": {
      "href": "/operations/1127729562914817/effects{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "/operations?cursor=1127729562914817\u0026order=asc"
    },
    "self": {
      "href": "/operations/1127729562914817"
    },
    "succeeds": {
      "href": "/operations?cursor=1127729562914817\u0026order=desc"
    },
    "transaction": {
      "href": "/transactions/1127729562914816"
    }
  },
  "amount": "11.27827",
  "buying_asset_code": "USD",
  "buying_asset_issuer": "GDS5JW5E6DRSSN5XK4LW7E6VUMFKKE2HU5WCOVFTO7P2RP7OXVCBLJ3Y",
  "buying_asset_type": "credit_alphanum4",
  "id": "1127729562914817",
  "offer_id": "9",
  "paging_token": "1127729562914817",
  "price": "1.0",
  "price_r": {
    "d": 1,
    "n": 1
  },
  "selling_asset_type": "native",
  "transaction_successful": true,
  "type_i": 4,
  "type": "create_passive_sell_offer"
}
```


<a id="set-options"></a>
### Set Options

Use “Set Options” operation to set following options to your account:
* Set/clear account flags:
  * AUTH_REQUIRED_FLAG (0x1) - if set, TrustLines are created with authorized set to `false` requiring the issuer to set it for each TrustLine.
  * AUTH_REVOCABLE_FLAG (0x2) - if set, the authorized flag in TrustLines can be cleared. Otherwise, authorization cannot be revoked.
* Set the account’s inflation destination.
* Add new signers to the account.
* Set home domain.


#### Attributes

| Field             | Type   | Description                                                                  |
|-------------------|--------|------------------------------------------------------------------------------|
| signer_key        | string | The public key of the new signer.                                            |
| signer_weight     | int    | The weight of the new signer (1-255).                                        |
| master_key_weight | int    | The weight of the master key (1-255).                                        |
| low_threshold     | int    | The sum weight for the low threshold.                                        |
| med_threshold     | int    | The sum weight for the medium threshold.                                     |
| high_threshold    | int    | The sum weight for the high threshold.                                       |
| home_domain       | string | The home domain used for reverse federation lookup                           |
| set_flags         | array  | The array of numeric values of flags that has been set in this operation     |
| set_flags_s       | array  | The array of string values of flags that has been set in this operation      |
| clear_flags       | array  | The array of numeric values of flags that has been cleared in this operation |
| clear_flags_s     | array  | The array of string values of flags that has been cleared in this operation  |


#### Example

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/operations/4126510088720385/effects?cursor=&limit=10&order=asc"
    },
    "next": {
      "href": "https://frontier.testnet.digitalbits.io/operations/4126510088720385/effects?cursor=4126510088720385-1&limit=10&order=asc"
    },
    "prev": {
      "href": "https://frontier.testnet.digitalbits.io/operations/4126510088720385/effects?cursor=4126510088720385-1&limit=10&order=desc"
    }
  },
  "_embedded": {
    "records": [
      {
        "_links": {
          "operation": {
            "href": "https://frontier.testnet.digitalbits.io/operations/4126510088720385"
          },
          "succeeds": {
            "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4126510088720385-1"
          },
          "precedes": {
            "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4126510088720385-1"
          }
        },
        "id": "0004126510088720385-0000000001",
        "paging_token": "4126510088720385-1",
        "account": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
        "type": "account_home_domain_updated",
        "type_i": 5,
        "created_at": "2021-06-16T13:55:12Z",
        "home_domain": "developers.digitalbits.io"
      }
    ]
  }
}

```

<a id="change-trust"></a>
### Change Trust

Use “Change Trust” operation to create/update/delete a trust line from the source account to another. The issuer being trusted and the asset code are in the given Asset object.

#### Attributes

| Field        | Type   | Description                                  |
|--------------|--------|----------------------------------------------|
| asset_code   | string | Asset code.                                  |
| asset_issuer | string | Asset issuer.                                |
| asset_type   | string | Asset type (native / alphanum4 / alphanum12) |
| trustee      | string | Trustee account.                             |
| trustor      | string | Trustor account.                             |
| limit        | string | The limit for the asset.                     |

#### Example

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/operations/4113161330364417"
    },
    "transaction": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/49d8adc71d64c6fba21883a59ffc3239194b6ace7b889f2835eda92d8bb41f0b"
    },
    "effects": {
      "href": "https://frontier.testnet.digitalbits.io/operations/4113161330364417/effects"
    },
    "succeeds": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4113161330364417"
    },
    "precedes": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4113161330364417"
    }
  },
  "id": "4113161330364417",
  "paging_token": "4113161330364417",
  "transaction_successful": true,
  "source_account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
  "type": "change_trust",
  "type_i": 6,
  "created_at": "2021-06-16T08:58:29Z",
  "transaction_hash": "49d8adc71d64c6fba21883a59ffc3239194b6ace7b889f2835eda92d8bb41f0b",
  "asset_type": "credit_alphanum4",
  "asset_code": "USD",
  "asset_issuer": "GB4RZUSF3HZGCAKB3VBM2S7QOHHC5KTV3LLZXGBYR5ZO4B26CKHFZTSZ",
  "limit": "1000.0000000",
  "trustee": "GB4RZUSF3HZGCAKB3VBM2S7QOHHC5KTV3LLZXGBYR5ZO4B26CKHFZTSZ",
  "trustor": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY"
}

```

<a id="allow-trust"></a>
### Allow Trust

Updates the "authorized" flag of an existing trust line this is called by the issuer of the asset.

Heads up! Unless the issuing account has `AUTH_REVOCABLE_FLAG` set than the "authorized" flag can only be set and never cleared.

#### Attributes

| Field        | Type   | Description                                             |
|--------------|--------|---------------------------------------------------------|
| asset_code   | string | Asset code.                                             |
| asset_issuer | string | Asset issuer.                                           |
| asset_type   | string | Asset type (native / alphanum4 / alphanum12)            |
| authorize    | bool   | `true` when allowing trust, `false` when revoking trust |
| trustee      | string | Trustee account.                                        |
| trustor      | string | Trustor account.                                        |

#### Example

```json
{
  "_links": {
    "effects": {
      "href": "/operations/34359742465/effects/{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "/operations?cursor=34359742465\u0026order=asc"
    },
    "self": {
      "href": "/operations/34359742465"
    },
    "succeeds": {
      "href": "/operations?cursor=34359742465\u0026order=desc"
    },
    "transaction": {
      "href": "/transactions/34359742464"
    }
  },
  "asset_code": "USD",
  "asset_issuer": "GC23QF2HUE52AMXUFUH3AYJAXXGXXV2VHXYYR6EYXETPKDXZSAW67XO4",
  "asset_type": "credit_alphanum4",
  "authorize": true,
  "id": "34359742465",
  "paging_token": "34359742465",
  "trustee": "GC23QF2HUE52AMXUFUH3AYJAXXGXXV2VHXYYR6EYXETPKDXZSAW67XO4",
  "trustor": "GBXGQJWVLWOYHFLVTKWV5FGHA3LNYY2JQKM7OAJAUEQFU6LPCSEFVXON",
  "transaction_successful": true,
  "type_i": 7,
  "type": "allow_trust"
}
```

<a id="account-merge"></a>
### Account Merge

Removes the account and transfers all remaining XDB to the destination account.

#### Attributes

| Field | Type   | Description                                                 |
|-------|--------|-------------------------------------------------------------|
| into  | string | Account ID where funds of deleted account were transferred. |

#### Example
```json
{
  "_links": {
    "effects": {
      "href": "/operations/799357838299137/effects{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "/operations?cursor=799357838299137\u0026order=asc"
    },
    "self": {
      "href": "/operations/799357838299137"
    },
    "succeeds": {
      "href": "/operations?cursor=799357838299137\u0026order=desc"
    },
    "transaction": {
      "href": "/transactions/799357838299136"
    }
  },
  "account": "GBCR5OVQ54S2EKHLBZMK6VYMTXZHXN3T45Y6PRX4PX4FXDMJJGY4FD42",
  "id": "799357838299137",
  "into": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K",
  "paging_token": "799357838299137",
  "transaction_successful": true,
  "type_i": 8,
  "type": "account_merge"
}
```

<a id="inflation"></a>
### Inflation

Runs inflation.

#### Example

```json
{
  "_links": {
    "effects": {
      "href": "/operations/12884914177/effects/{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "/operations?cursor=12884914177\u0026order=asc"
    },
    "self": {
      "href": "/operations/12884914177"
    },
    "succeeds": {
      "href": "/operations?cursor=12884914177\u0026order=desc"
    },
    "transaction": {
      "href": "/transactions/12884914176"
    }
  },
  "id": "12884914177",
  "paging_token": "12884914177",
  "transaction_successful": true,
  "type_i": 9,
  "type": "inflation"
}
```

<a id="manage-data"></a>
### Manage Data

Set, modify or delete a Data Entry (name/value pair) for an account.

#### Example

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/operations/4115467727802369"
    },
    "transaction": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/896ebf830df1ad883603a9b08a3486110e8a81e877fb4078f8edb341d38266c9"
    },
    "effects": {
      "href": "https://frontier.testnet.digitalbits.io/operations/4115467727802369/effects"
    },
    "succeeds": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4115467727802369"
    },
    "precedes": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4115467727802369"
    }
  },
  "id": "4115467727802369",
  "paging_token": "4115467727802369",
  "transaction_successful": true,
  "source_account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
  "type": "manage_data",
  "type_i": 10,
  "created_at": "2021-06-16T09:49:53Z",
  "transaction_hash": "896ebf830df1ad883603a9b08a3486110e8a81e877fb4078f8edb341d38266c9",
  "name": "user-id",
  "value": "WERCRm91bmRhdGlvbg=="
}

```

<a id="bump-sequence"></a>
### Bump Sequence

Bumps forward the sequence number of the source account of the operation, allowing it to invalidate any transactions with a smaller sequence number.

#### Attributes

| Field  | Type   | Description                                                       |
|--------|--------|-------------------------------------------------------------------|
| bumpTo | number | Desired value for the operation’s source account sequence number. |

#### Example
```json
{
  "_links": {
    "self": {
      "href": "/operations/1743756726273"
    },
    "transaction": {
      "href": "/transactions/328436a8dffaf6ca33c08a93279234c7d3eaf1c028804152614187dc76b7168d"
    },
    "effects": {
      "href": "/operations/1743756726273/effects"
    },
    "succeeds": {
      "href": "/effects?order=desc&cursor=1743756726273"
    },
    "precedes": {
      "href": "/effects?order=asc&cursor=1743756726273"
    }
  },
  "id": "1743756726273",
  "paging_token": "1743756726273",
  "source_account": "GBHPJ3VMVT3X7Y6HIIAPK7YPTZCF3CWO4557BKGX2GVO4O7EZHIBELLH",
  "type": "bump_sequence",
  "type_i": 11,
  "transaction_hash": "328436a8dffaf6ca33c08a93279234c7d3eaf1c028804152614187dc76b7168d",
  "bump_to": "1273737228"
}
```

## Endpoints

| Resource                                           | Type       | Resource URI Template                           |
|----------------------------------------------------|------------|-------------------------------------------------|
| [All Operations](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/operations-all)             | Collection | `/operations`                                   |
| [Operations Details](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/operations-single)      | Single     | `/operations/:id`                               |
| [Ledger Operations](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/operations-for-ledger)   | Collection | `/ledgers/{id}/operations{?cursor,limit,order}` |
| [Account Operations](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/operations-for-account) | Collection | `/accounts/:account_id/operations`              |
| [Account Payments](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/payments-for-account)     | Collection | `/accounts/:account_id/payments`                |
