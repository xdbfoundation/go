package frontierclient

import "github.com/xdbfoundation/go/support/errors"

// BuildURL returns the url for getting fee stats about a running frontier instance
func (fr feeStatsRequest) BuildURL() (endpoint string, err error) {
	endpoint = fr.endpoint
	if endpoint == "" {
		err = errors.New("invalid request: too few parameters")
	}

	return
}
