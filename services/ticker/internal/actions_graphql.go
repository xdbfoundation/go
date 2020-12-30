package ticker

import (
	"github.com/digitalbits/go/services/ticker/internal/gql"
	"github.com/digitalbits/go/services/ticker/internal/tickerdb"
	hlog "github.com/digitalbits/go/support/log"
)

func StartGraphQLServer(s *tickerdb.TickerSession, l *hlog.Entry, port string) {
	graphql := gql.New(s, l)

	graphql.Serve(port)
}
