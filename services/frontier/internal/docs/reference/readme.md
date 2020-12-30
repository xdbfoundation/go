---
title: Overview
---

Frontier is an API server for the DigitalBits ecosystem.  It acts as the interface between [digitalbits-core](https://github.com/digitalbits/digitalbits-core) and applications that want to access the DigitalBits network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the DigitalBits ecosystem](https://www.digitalbits.org/developers/guides/) for details of where Frontier fits in.

Frontier provides a RESTful API to allow client applications to interact with the DigitalBits network. You can communicate with Frontier using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a DigitalBits SDK in the language of your client.
SDF provides a [JavaScript SDK](https://www.digitalbits.org/developers/js-digitalbits-sdk/reference/index.html) for clients to use to interact with Frontier.

SDF runs a instance of Frontier that is connected to the test net: [https://frontier-testnet.digitalbits.org/](https://frontier-testnet.digitalbits.org/) and one that is connected to the public DigitalBits network:
[https://frontier.digitalbits.org/](https://frontier.digitalbits.org/).

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
