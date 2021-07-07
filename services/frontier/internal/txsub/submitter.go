package txsub

import (
	"context"
	"net/http"
	"time"

	"github.com/xdbfoundation/go/clients/digitalbitscore"
	proto "github.com/xdbfoundation/go/protocols/digitalbitscore"
	"github.com/xdbfoundation/go/support/errors"
	"github.com/xdbfoundation/go/support/log"
)

// NewDefaultSubmitter returns a new, simple Submitter implementation
// that submits directly to the digitalbits-core at `url` using the http client
// `h`.
func NewDefaultSubmitter(h *http.Client, url string) Submitter {
	return &submitter{
		DigitalBitsCore: &digitalbitscore.Client{
			HTTP: h,
			URL:  url,
		},
		Log: log.DefaultLogger.WithField("service", "txsub.submitter"),
	}
}

// submitter is the default implementation for the Submitter interface.  It
// submits directly to the configured digitalbits-core instance using the
// configured http client.
type submitter struct {
	DigitalBitsCore *digitalbitscore.Client
	Log         *log.Entry
}

// Submit sends the provided envelope to digitalbits-core and parses the response into
// a SubmissionResult
func (sub *submitter) Submit(ctx context.Context, env string) (result SubmissionResult) {
	start := time.Now()
	defer func() {
		result.Duration = time.Since(start)
		sub.Log.Ctx(ctx).WithFields(log.F{
			"err":      result.Err,
			"duration": result.Duration.Seconds(),
		}).Info("Submitter result")
	}()

	cresp, err := sub.DigitalBitsCore.SubmitTransaction(ctx, env)
	if err != nil {
		result.Err = errors.Wrap(err, "failed to submit")
		return
	}

	// interpet response
	if cresp.IsException() {
		result.Err = errors.Errorf("digitalbits-core exception: %s", cresp.Exception)
		return
	}

	switch cresp.Status {
	case proto.TXStatusError:
		result.Err = &FailedTransactionError{cresp.Error}
	case proto.TXStatusPending, proto.TXStatusDuplicate, proto.TXStatusTryAgainLater:
		//noop.  A nil Err indicates success
	default:
		result.Err = errors.Errorf("Unrecognized digitalbits-core status response: %s", cresp.Status)
	}

	return
}
