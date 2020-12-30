---
title: Overview
---

The Go SDK is a set of packages for interacting with most aspects of the DigitalBits ecosystem. The primary component is the Frontier SDK, which provides convenient access to Frontier services. There are also packages for other DigitalBits services such as [TOML support](https://github.com/digitalbits/digitalbits-protocol/blob/master/ecosystem/sep-0001.md) and [federation](https://github.com/digitalbits/digitalbits-protocol/blob/master/ecosystem/sep-0002.md).

## Frontier SDK

The Frontier SDK is composed of two complementary libraries: `txnbuild` + `frontierclient`.
The `txnbuild` ([source](https://github.com/digitalbits/go/tree/master/txnbuild), [docs](https://godoc.org/github.com/digitalbits/go/txnbuild)) package enables the construction, signing and encoding of DigitalBits [transactions](https://www.digitalbits.org/developers/guides/concepts/transactions.html) and [operations](https://www.digitalbits.org/developers/guides/concepts/list-of-operations.html) in Go. The `frontierclient` ([source](https://github.com/digitalbits/go/tree/master/clients/frontierclient), [docs](https://godoc.org/github.com/digitalbits/go/clients/frontierclient)) package provides a web client for interfacing with [Frontier](https://www.digitalbits.org/developers/guides/get-started/) server REST endpoints to retrieve ledger information, and to submit transactions built with `txnbuild`.

## List of major SDK packages

- `frontierclient` ([source](https://github.com/digitalbits/go/tree/master/clients/frontierclient), [docs](https://godoc.org/github.com/digitalbits/go/clients/frontierclient)) - programmatic client access to Frontier
- `txnbuild` ([source](https://github.com/digitalbits/go/tree/master/txnbuild), [docs](https://godoc.org/github.com/digitalbits/go/txnbuild)) - construction, signing and encoding of DigitalBits transactions and operations
- `digitalbitstoml` ([source](https://github.com/digitalbits/go/tree/master/clients/digitalbitstoml), [docs](https://godoc.org/github.com/digitalbits/go/clients/digitalbitstoml)) - parse [DigitalBits.toml](../../guides/concepts/digitalbits-toml.md) files from the internet
- `federation` ([source](https://godoc.org/github.com/digitalbits/go/clients/federation)) - resolve federation addresses  into digitalbits account IDs, suitable for use within a transaction

