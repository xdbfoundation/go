## Streaming

Certain endpoints in Frontier can be called in streaming mode using Server-Sent Events. This mode will keep the connection to Frontier open and Frontier will continue to return responses as ledgers close. All parameters for the endpoints that allow this mode are the same. The way a caller initiates this mode is by setting `Accept: text/event-stream` in the HTTP header when you make the request.
You can read an example of using the streaming mode in the [Follow Received Payments](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/tutorials/follow-received-payments) tutorial.

Endpoints that currently support streaming:

- [Account](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/accounts-single)
- [Effects](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/effects-all)
- [Ledgers](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/ledgers-all)
- [Offers](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/offers-for-account)
- [Operations](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/operations-all)
- [Orderbook](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/orderbook-details)
- [Payments](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/payments-all)
- [Transactions](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/transactions-all)
- [Trades](https://developers.digitalbits.io/reference/go/services/frontier/internal/docs/reference/endpoints/trades)
