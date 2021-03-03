# frontierclient


`frontierclient` is a [DigitalBits Go SDK](https://github.com/xdbfoundation/docs/tree/master/reference) package that provides client access to a frontier server. It supports all endpoints exposed by the [frontier API](https://github.com/xdbfoundation/go/tree/master/services/frontier).

This project is maintained by the XDB Foundation.

## Getting Started
This library is aimed at developers building Go applications that interact with the [DigitalBits network](https://www.digitalbits.io/). It allows users to query the network and submit transactions to the network. The recommended transaction builder for Go programmers is [txnbuild](https://github.com/xdbfoundation/go/tree/master/txnbuild). Together, these two libraries provide a complete DigitalBits SDK.

* The [frontierclient API reference](https://github.com/xdbfoundation/go/tree/master/clients/frontierclient).
* The [txnbuild API reference](https://github.com/xdbfoundation/go/tree/master/txnbuild).

### Prerequisites
* Go 1.14 or greater
* [Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies

### Installing
* `go get github.com/digitalbits/go/clients/frontierclient`

### Usage

``` golang
    ...
    import hClient "github.com/digitalbits/go/clients/frontierclient"
    ...

    // Use the default pubnet client
    client := hClient.DefaultPublicNetClient

    // Create an account request
    accountRequest := hClient.AccountRequest{AccountID: "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU"}

    // Load the account detail from the network
    account, err := client.AccountDetail(accountRequest)
    if err != nil {
        fmt.Println(err)
        return
    }
    // Account contains information about the digitalbits account
    fmt.Print(account)
```
For more examples, refer to the [documentation](https://github.com/xdbfoundation/go/tree/master/clients/frontierclient).

## Running the tests
Run the unit tests from the package directory: `go test`

## Contributing
Please read [Code of Conduct](https://digitalbits.io/community-guidelines/) to understand this project's communication rules.

To submit improvements and fixes to this library, please see [CONTRIBUTING](https://github.com/xdbfoundation/docs/blob/master/CONTRIBUTING.md).

## License
This project is licensed under the Apache License - see the [LICENSE](https://www.apache.org/licenses/LICENSE-2.0) file for details.
