This endpoint represents successful [operations](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/operation) that were included in valid [transactions](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/transaction) that affected a particular [account](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/account).

This endpoint can also be used in [streaming](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/streaming) mode so it is possible to use it to listen for new operations that affect a given account as they happen.
If called in streaming mode Frontier will start at the earliest known operation unless a `cursor` is set. In that case it will start from the `cursor`. You can also set `cursor` value to `now` to only stream operations created since your request time.

## Request

```
GET /accounts/{account}/operations{?cursor,limit,order,include_failed}
```

### Arguments

| name     | notes                          | description                                                      | example                                                   |
| ------   | -------                        | -----------                                                      | -------                                                   |
| `account`| required, string               | Account ID                                                  | `GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY`|
| `?cursor`| optional, default _null_       | A paging token, specifying where to start returning records from.  When streaming this can be set to `now` to stream object created since your request time. | `1623820974`                                             |
| `?order` | optional, string, default `asc`| The order in which to return rows, "asc" or "desc".              | `asc`                                                     |
| `?limit` | optional, number, default `10` | Maximum number of records to return.                             | `200`
| `?include_failed` | optional, bool, default: `false` | Set to `true` to include operations of failed transactions in results. | `true` |                                                     |
| `?join` | optional, string, default: _null_ | Set to `transactions` to include the transactions which created each of the operations in the response. | `transactions` |

### curl Example Request

```sh
curl "https://frontier.testnet.digitalbits.io/accounts/GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY/operations"
```

### JavaScript Example Request

```js
var DigitalBitsSdk = require('digitalbits-sdk');
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

server.operations()
  .forAccount("GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY")
  .call()
  .then(function (operationsResult) {
    console.log(JSON.stringify(operationsResult.records))
  })
  .catch(function (err) {
    console.log(err)
  })
```

### JavaScript Streaming Example

```javascript
var DigitalBitsSdk = require('digitalbits-sdk')
var server = new DigitalBitsSdk.Server('https://frontier.testnet.digitalbits.io');

var operationHandler = function (operationResponse) {
    console.log(operationResponse);
};

var es = server.operations()
    .forAccount("GAKLBGHNHFQ3BMUYG5KU4BEWO6EYQHZHAXEWC33W34PH2RBHZDSQBD75")
    .cursor('now')
    .stream({
        onmessage: operationHandler
    })
```

## Response

This endpoint responds with a list of operations that affected the given account. See [operation resource](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/resources/operation) for reference.

### Example Response

```json
[
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113023891410945"
      },
      "transaction": {
        "href": "https://frontier.testnet.digitalbits.io/transactions/82b00c2fa0840ea768b08b2dedc0613f5c28ad3fe03001df8f22694523a3a1bc"
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113023891410945/effects"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4113023891410945"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4113023891410945"
      }
    },
    "id": "4113023891410945",
    "paging_token": "4113023891410945",
    "transaction_successful": true,
    "source_account": "GC3CLEUNQVWY36AHTGGX2NASAPHD6EBQXE63YH2B3PAASLCCIG4ELGTP",
    "type": "create_account",
    "type_i": 0,
    "created_at": "2021-06-16T08:55:24Z",
    "transaction_hash": "82b00c2fa0840ea768b08b2dedc0613f5c28ad3fe03001df8f22694523a3a1bc",
    "starting_balance": "10000.0000000",
    "funder": "GC3CLEUNQVWY36AHTGGX2NASAPHD6EBQXE63YH2B3PAASLCCIG4ELGTP",
    "account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY"
  },
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
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113165625331713"
      },
      "transaction": {
        "href": "https://frontier.testnet.digitalbits.io/transactions/21187d611ad7a050bc2f700186b3b95fb1d2abdc52410d838c68aafba408be8c"
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113165625331713/effects"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4113165625331713"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4113165625331713"
      }
    },
    "id": "4113165625331713",
    "paging_token": "4113165625331713",
    "transaction_successful": true,
    "source_account": "GB4RZUSF3HZGCAKB3VBM2S7QOHHC5KTV3LLZXGBYR5ZO4B26CKHFZTSZ",
    "type": "payment",
    "type_i": 1,
    "created_at": "2021-06-16T08:58:35Z",
    "transaction_hash": "21187d611ad7a050bc2f700186b3b95fb1d2abdc52410d838c68aafba408be8c",
    "asset_type": "credit_alphanum4",
    "asset_code": "USD",
    "asset_issuer": "GB4RZUSF3HZGCAKB3VBM2S7QOHHC5KTV3LLZXGBYR5ZO4B26CKHFZTSZ",
    "from": "GB4RZUSF3HZGCAKB3VBM2S7QOHHC5KTV3LLZXGBYR5ZO4B26CKHFZTSZ",
    "to": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "amount": "10.0000000"
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113603711995905"
      },
      "transaction": {
        "href": "https://frontier.testnet.digitalbits.io/transactions/847b33d9e54a8884a9b9c1fd68dc5560e1c61e181155aafc1145e934cc12535d"
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113603711995905/effects"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4113603711995905"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4113603711995905"
      }
    },
    "id": "4113603711995905",
    "paging_token": "4113603711995905",
    "transaction_successful": true,
    "source_account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "change_trust",
    "type_i": 6,
    "created_at": "2021-06-16T09:08:23Z",
    "transaction_hash": "847b33d9e54a8884a9b9c1fd68dc5560e1c61e181155aafc1145e934cc12535d",
    "asset_type": "credit_alphanum4",
    "asset_code": "EUR",
    "asset_issuer": "GDCIQQY2UKVNLLWGIX74DMTEAFCMQKAKYUWPBO7PLTHIHRKSFZN7V2FC",
    "limit": "1000.0000000",
    "trustee": "GDCIQQY2UKVNLLWGIX74DMTEAFCMQKAKYUWPBO7PLTHIHRKSFZN7V2FC",
    "trustor": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY"
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113608006963201"
      },
      "transaction": {
        "href": "https://frontier.testnet.digitalbits.io/transactions/2edafcd4e44ce7fd166ee32d80f61c86593084fdcf4ad792841907e06a14230b"
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4113608006963201/effects"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4113608006963201"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4113608006963201"
      }
    },
    "id": "4113608006963201",
    "paging_token": "4113608006963201",
    "transaction_successful": true,
    "source_account": "GDCIQQY2UKVNLLWGIX74DMTEAFCMQKAKYUWPBO7PLTHIHRKSFZN7V2FC",
    "type": "payment",
    "type_i": 1,
    "created_at": "2021-06-16T09:08:28Z",
    "transaction_hash": "2edafcd4e44ce7fd166ee32d80f61c86593084fdcf4ad792841907e06a14230b",
    "asset_type": "credit_alphanum4",
    "asset_code": "EUR",
    "asset_issuer": "GDCIQQY2UKVNLLWGIX74DMTEAFCMQKAKYUWPBO7PLTHIHRKSFZN7V2FC",
    "from": "GDCIQQY2UKVNLLWGIX74DMTEAFCMQKAKYUWPBO7PLTHIHRKSFZN7V2FC",
    "to": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "amount": "5.0000000"
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4114797712904193"
      },
      "transaction": {
        "href": "https://frontier.testnet.digitalbits.io/transactions/642eb344f7e79b3340f607f09337e330bc8884a26702baf29891620d92fb70c3"
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4114797712904193/effects"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4114797712904193"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4114797712904193"
      }
    },
    "id": "4114797712904193",
    "paging_token": "4114797712904193",
    "transaction_successful": true,
    "source_account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "change_trust",
    "type_i": 6,
    "created_at": "2021-06-16T09:34:52Z",
    "transaction_hash": "642eb344f7e79b3340f607f09337e330bc8884a26702baf29891620d92fb70c3",
    "asset_type": "credit_alphanum4",
    "asset_code": "HUF",
    "asset_issuer": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "limit": "1000.0000000",
    "trustee": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "trustor": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY"
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4114849252511745"
      },
      "transaction": {
        "href": "https://frontier.testnet.digitalbits.io/transactions/300dce27226621ff08dd873fb9e167e0ed5356dc12c1e96c5bc670c59e2a0d74"
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4114849252511745/effects"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4114849252511745"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4114849252511745"
      }
    },
    "id": "4114849252511745",
    "paging_token": "4114849252511745",
    "transaction_successful": true,
    "source_account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "type": "change_trust",
    "type_i": 6,
    "created_at": "2021-06-16T09:35:59Z",
    "transaction_hash": "300dce27226621ff08dd873fb9e167e0ed5356dc12c1e96c5bc670c59e2a0d74",
    "asset_type": "credit_alphanum4",
    "asset_code": "HUF",
    "asset_issuer": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "limit": "922337203685.4775807",
    "trustee": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "trustor": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY"
  },
  {
    "_links": {
      "self": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4114853547479041"
      },
      "transaction": {
        "href": "https://frontier.testnet.digitalbits.io/transactions/c60f666cda020d13033bb44926adf7f6c2b659857f13959e3988351055c0b52f"
      },
      "effects": {
        "href": "https://frontier.testnet.digitalbits.io/operations/4114853547479041/effects"
      },
      "succeeds": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4114853547479041"
      },
      "precedes": {
        "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4114853547479041"
      }
    },
    "id": "4114853547479041",
    "paging_token": "4114853547479041",
    "transaction_successful": true,
    "source_account": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "type": "payment",
    "type_i": 1,
    "created_at": "2021-06-16T09:36:04Z",
    "transaction_hash": "c60f666cda020d13033bb44926adf7f6c2b659857f13959e3988351055c0b52f",
    "asset_type": "credit_alphanum4",
    "asset_code": "HUF",
    "asset_issuer": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "from": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
    "to": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
    "amount": "50000.0000000"
  },
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
]
```

### Example Streaming Event

```json
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/operations/4114849252511745"
    },
    "transaction": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/300dce27226621ff08dd873fb9e167e0ed5356dc12c1e96c5bc670c59e2a0d74"
    },
    "effects": {
      "href": "https://frontier.testnet.digitalbits.io/operations/4114849252511745/effects"
    },
    "succeeds": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4114849252511745"
    },
    "precedes": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4114849252511745"
    }
  },
  "id": "4114849252511745",
  "paging_token": "4114849252511745",
  "transaction_successful": true,
  "source_account": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
  "type": "change_trust",
  "type_i": 6,
  "created_at": "2021-06-16T09:35:59Z",
  "transaction_hash": "300dce27226621ff08dd873fb9e167e0ed5356dc12c1e96c5bc670c59e2a0d74",
  "asset_type": "credit_alphanum4",
  "asset_code": "HUF",
  "asset_issuer": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
  "limit": "922337203685.4775807",
  "trustee": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
  "trustor": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY"
}
{
  "_links": {
    "self": {
      "href": "https://frontier.testnet.digitalbits.io/operations/4114853547479041"
    },
    "transaction": {
      "href": "https://frontier.testnet.digitalbits.io/transactions/c60f666cda020d13033bb44926adf7f6c2b659857f13959e3988351055c0b52f"
    },
    "effects": {
      "href": "https://frontier.testnet.digitalbits.io/operations/4114853547479041/effects"
    },
    "succeeds": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=desc&cursor=4114853547479041"
    },
    "precedes": {
      "href": "https://frontier.testnet.digitalbits.io/effects?order=asc&cursor=4114853547479041"
    }
  },
  "id": "4114853547479041",
  "paging_token": "4114853547479041",
  "transaction_successful": true,
  "source_account": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
  "type": "payment",
  "type_i": 1,
  "created_at": "2021-06-16T09:36:04Z",
  "transaction_hash": "c60f666cda020d13033bb44926adf7f6c2b659857f13959e3988351055c0b52f",
  "asset_type": "credit_alphanum4",
  "asset_code": "HUF",
  "asset_issuer": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
  "from": "GCHQ6AOZST6YPMROCQWPE3SVFY57FHPYC3WJGGSFCHOQ5HFZC5HSHQYK",
  "to": "GDFOHLMYCXVZD2CDXZLMW6W6TMU4YO27XFF2IBAFAV66MSTPDDSK2LAY",
  "amount": "50000.0000000"
}
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


## Possible Errors

- The [standard errors](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors#standard-errors).
- [not_found](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/errors/not-found): A `not_found` error will be returned if there is no account whose ID matches the `account` argument.
