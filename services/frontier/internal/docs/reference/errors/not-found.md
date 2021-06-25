When Frontier can't find whatever data you are requesting, it will return a `not_found` error. This
is similar to a
[HTTP 404 Error](https://developer.mozilla.org/en-US/docs/Web/HTTP/Response_codes) error
response.

Incorrect URL path parameters or missing data are the common reasons for this error. If you
navigate using a link from a valid response, you should never receive this error message.

## Attributes

As with all errors Frontier returns, `not_found` follows the
[Problem Details for HTTP APIs](https://tools.ietf.org/html/draft-ietf-appsawg-http-problem-00)
draft specification guide and thus has the following attributes:

| Attribute   | Type   | Description                                                                     |
| ----------- | ------ | ------------------------------------------------------------------------------- |
| `type`      | URL    | The identifier for the error.  This is a URL that can be visited in the browser.|
| `title`     | String | A short title describing the error.                                             |
| `status`    | Number | An HTTP status code that maps to the error.                                     |
| `detail`    | String | A more detailed description of the error.                                       |

## Example

```shell
$ curl -X GET "https://frontier.testnet.digitalbits.io/accounts/accountthatdoesntexist"
{
  "type": "https://digitalbits.org/frontier-errors/bad_request",
  "title": "Bad Request",
  "status": 400,
  "detail": "The request you sent was invalid in some way",
  "extras": {
    "invalid_field": "account_id",
    "reason": "invalid address"
  }
}
```

## Related

- [Not Acceptable](https://github.com/xdbfoundation/go/tree/master/services/frontier/internal/docs/reference/errors/not-acceptable.md)
