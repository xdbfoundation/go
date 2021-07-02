---
title: Streaming
---

## Streaming

Certain endpoints in Frontier can be called in streaming mode using Server-Sent Events. This mode will keep the connection to Frontier open and Frontier will continue to return responses as ledgers close. All parameters for the endpoints that allow this mode are the same. The way a caller initiates this mode is by setting `Accept: text/event-stream` in the HTTP header when you make the request.
You can read an example of using the streaming mode in the [Follow Received Payments](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/tutorials/follow-received-payments.md) tutorial.

Endpoints that currently support streaming:

- [Account](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/accounts-single.md)
- [Effects](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/effects-all.md)
- [Ledgers](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/ledgers-all.md)
- [Offers](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/offers-for-account.md)
- [Operations](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/operations-all.md)
- [Orderbook](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/orderbook-details.md)
- [Payments](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/payments-all.md)
- [Transactions](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/transactions-all.md)
- [Trades](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/endpoints/trades.md)
