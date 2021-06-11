Each account in DigitalBits network can contain multiple key/value pairs associated with it. Frontier can be used to retrieve value of each data key.

When frontier returns information about a single account data key it uses the following format:

## Attributes

| Attribute | Type | | 
| --- | --- | --- |
| value | base64-encoded string | The base64-encoded value for the key |

## Example

```json
{
  "value": "MTAw"
}
```
