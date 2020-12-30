package resourceadapter

import (
	"context"
	protocol "github.com/digitalbits/go/protocols/frontier"
	"github.com/digitalbits/go/services/frontier/internal/txsub"
)

// Populate fills out the details
func PopulateTransactionResultCodes(ctx context.Context,
	transactionHash string,
	dest *protocol.TransactionResultCodes,
	fail *txsub.FailedTransactionError,
) (err error) {

	dest.TransactionCode, err = fail.TransactionResultCode(transactionHash)
	if err != nil {
		return
	}

	dest.OperationCodes, err = fail.OperationResultCodes()
	if err != nil {
		return
	}

	return
}
