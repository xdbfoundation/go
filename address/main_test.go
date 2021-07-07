package address

import (
	"testing"

	"github.com/xdbfoundation/go/support/errors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cases := []struct {
		Name            string
		Domain          string
		ExpectedAddress string
	}{
		{"scott", "digitalbits.org", "scott*digitalbits.org"},
		{"", "digitalbits.org", "*digitalbits.org"},
		{"scott", "", "scott*"},
	}

	for _, c := range cases {
		actual := New(c.Name, c.Domain)
		assert.Equal(t, c.ExpectedAddress, actual)
	}
}

func TestSplit(t *testing.T) {
	cases := []struct {
		CaseName       string
		Address        string
		ExpectedName   string
		ExpectedDomain string
		ExpectedError  error
	}{
		{"happy path", "scott*digitalbits.org", "scott", "digitalbits.org", nil},
		{"blank", "", "", "", ErrInvalidAddress},
		{"blank name", "*digitalbits.org", "", "", ErrInvalidName},
		{"blank domain", "scott*", "", "", ErrInvalidDomain},
		{"invalid domain", "scott*--3.com", "", "", ErrInvalidDomain},
	}

	for _, c := range cases {
		name, domain, err := Split(c.Address)

		if c.ExpectedError == nil {
			assert.Equal(t, c.ExpectedName, name)
			assert.Equal(t, c.ExpectedDomain, domain)
		} else {
			assert.Equal(t, c.ExpectedError, errors.Cause(err))
		}
	}
}
