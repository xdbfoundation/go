# frontierclient


`frontierclient` is a [DigitalBits Go SDK](https://www.digitalbits.org/developers/reference/) package that provides client access to a frontier server. It supports all endpoints exposed by the [frontier API](https://www.digitalbits.org/developers/frontier/reference/index.html).

This project is maintained by the DigitalBits Development Foundation.

## Getting Started
This library is aimed at developers building Go applications that interact with the [DigitalBits network](https://www.digitalbits.org/). It allows users to query the network and submit transactions to the network. The recommended transaction builder for Go programmers is [txnbuild](https://github.com/digitalbits/go/tree/master/txnbuild). Together, these two libraries provide a complete DigitalBits SDK.

* The [frontierclient API reference](https://godoc.org/github.com/digitalbits/go/clients/frontierclient).
* The [txnbuild API reference](https://godoc.org/github.com/digitalbits/go/txnbuild).

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
For more examples, refer to the [documentation](https://godoc.org/github.com/digitalbits/go/clients/frontierclient).

## Running the tests
Run the unit tests from the package directory: `go test`

## Contributing
Please read [Code of Conduct](https://github.com/digitalbits/.github/blob/master/CODE_OF_CONDUCT.md) to understand this project's communication rules.

To submit improvements and fixes to this library, please see [CONTRIBUTING](../CONTRIBUTING.md).

## License
This project is licensed under the Apache License - see the [LICENSE](../../LICENSE-APACHE.txt) file for details.
