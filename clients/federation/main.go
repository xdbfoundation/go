package federation

import (
	"net/http"
	"net/url"

	hc "github.com/digitalbits/go/clients/frontierclient"
	"github.com/digitalbits/go/clients/digitalbitstoml"
	proto "github.com/digitalbits/go/protocols/federation"
)

// FederationResponseMaxSize is the maximum size of response from a federation server
const FederationResponseMaxSize = 100 * 1024

// DefaultTestNetClient is a default federation client for testnet
var DefaultTestNetClient = &Client{
	HTTP:        http.DefaultClient,
	Frontier:     hc.DefaultTestNetClient,
	DigitalBitsTOML: digitalbitstoml.DefaultClient,
}

// DefaultPublicNetClient is a default federation client for pubnet
var DefaultPublicNetClient = &Client{
	HTTP:        http.DefaultClient,
	Frontier:     hc.DefaultPublicNetClient,
	DigitalBitsTOML: digitalbitstoml.DefaultClient,
}

// Client represents a client that is capable of resolving a federation request
// using the internet.
type Client struct {
	DigitalBitsTOML DigitalBitsTOML
	HTTP        HTTP
	Frontier     Frontier
	AllowHTTP   bool
}

type ClientInterface interface {
	LookupByAddress(addy string) (*proto.NameResponse, error)
	LookupByAccountID(aid string) (*proto.IDResponse, error)
	ForwardRequest(domain string, fields url.Values) (*proto.NameResponse, error)
}

// Frontier represents a frontier client that can be consulted for data when
// needed as part of the federation protocol
type Frontier interface {
	HomeDomainForAccount(aid string) (string, error)
}

// HTTP represents the http client that a federation client uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// DigitalBitsTOML represents a client that can resolve a given domain name to
// digitalbits.toml file.  The response is used to find the federation server that a
// query should be made against.
type DigitalBitsTOML interface {
	GetDigitalBitsToml(domain string) (*digitalbitstoml.Response, error)
}

// confirm interface conformity
var _ DigitalBitsTOML = digitalbitstoml.DefaultClient
var _ HTTP = http.DefaultClient
var _ ClientInterface = &Client{}
