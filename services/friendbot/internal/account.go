package internal

import (
	"strconv"

	"github.com/xdbfoundation/go/clients/frontierclient"
	"github.com/xdbfoundation/go/support/errors"
)

// Account implements the `txnbuild.Account` interface.
type Account struct {
	AccountID string
	Sequence  int64
}

// GetAccountID returns the Account ID.
func (a Account) GetAccountID() string {
	return a.AccountID
}

// IncrementSequenceNumber increments the internal record of the
// account's sequence number by 1.
func (a Account) IncrementSequenceNumber() (int64, error) {
	a.Sequence++
	return a.Sequence, nil
}

func (a Account) GetSequenceNumber() (int64, error) {
	return a.Sequence, nil
}

// RefreshSequenceNumber gets an Account's correct in-memory sequence number from Frontier.
func (a *Account) RefreshSequenceNumber(hclient *frontierclient.Client) error {
	accountRequest := frontierclient.AccountRequest{AccountID: a.GetAccountID()}
	accountDetail, err := hclient.AccountDetail(accountRequest)
	if err != nil {
		return errors.Wrap(err, "getting account detail")
	}
	seq, err := strconv.ParseInt(accountDetail.Sequence, 10, 64)
	if err != nil {
		return errors.Wrap(err, "parsing account seqnum")
	}
	a.Sequence = seq
	return nil
}
