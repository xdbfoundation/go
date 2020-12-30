---
title: Frontier
replacement: https://developers.digitalbits.org/docs/run-api-server/
---

Frontier is the server for the client facing API for the DigitalBits ecosystem.  It acts as the interface between [digitalbits-core](https://www.digitalbits.org/developers/software/#digitalbits-core) and applications that want to access the DigitalBits network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the DigitalBits ecosystem](https://www.digitalbits.org/developers/guides/) for more details.

You can interact directly with frontier via curl or a web browser but SDF provides a [JavaScript SDK](https://www.digitalbits.org/developers/js-digitalbits-sdk/reference/) for clients to use to interact with Frontier.

SDF runs a instance of Frontier that is connected to the test net [https://frontier-testnet.digitalbits.org/](https://frontier-testnet.digitalbits.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/digitalbits/js-digitalbits-sdk)
- [Go](https://github.com/digitalbits/go/tree/master/clients/frontierclient)
- [Java](https://github.com/digitalbits/java-digitalbits-sdk)

Community maintained libraries for interacting with Frontier in other languages:<br>
- [Python](https://github.com/DigitalBitsCN/py-digitalbits-base)
- [C# .NET Core 2.x](https://github.com/elucidsoft/dotnetcore-digitalbits-sdk)
- [Ruby](https://github.com/astroband/ruby-digitalbits-sdk)
- [iOS and macOS](https://github.com/Soneso/digitalbits-ios-mac-sdk)
- [Scala SDK](https://github.com/synesso/scala-digitalbits-sdk)
- [C++ SDK](https://github.com/bnogalm/DigitalBitsQtSDK)
