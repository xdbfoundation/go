package history

import (
	"testing"
	"time"

	"github.com/xdbfoundation/go/services/frontier/internal/test"
)

func TestLatestLedger(t *testing.T) {
	tt := test.Start(t)
	tt.Scenario("base")
	defer tt.Finish()
	q := &Q{tt.FrontierSession()}

	var seq int
	err := q.LatestLedger(&seq)

	if tt.Assert.NoError(err) {
		tt.Assert.Equal(3, seq)
	}
}

func TestLatestLedgerSequenceClosedAt(t *testing.T) {
	tt := test.Start(t)
	tt.Scenario("base")
	defer tt.Finish()
	q := &Q{tt.FrontierSession()}

	sequence, closedAt, err := q.LatestLedgerSequenceClosedAt()
	if tt.Assert.NoError(err) {
		tt.Assert.Equal(int32(3), sequence)
		tt.Assert.Equal("2019-10-31T13:19:46Z", closedAt.Format(time.RFC3339))
	}

	test.ResetFrontierDB(t, tt.FrontierDB)

	sequence, closedAt, err = q.LatestLedgerSequenceClosedAt()
	if tt.Assert.NoError(err) {
		tt.Assert.Equal(int32(0), sequence)
		tt.Assert.Equal("0001-01-01T00:00:00Z", closedAt.Format(time.RFC3339))
	}
}

func TestGetLatestLedgerEmptyDB(t *testing.T) {
	tt := test.Start(t)
	defer tt.Finish()
	test.ResetFrontierDB(t, tt.FrontierDB)
	q := &Q{tt.FrontierSession()}

	value, err := q.GetLatestLedger()
	tt.Assert.NoError(err)
	tt.Assert.Equal(uint32(0), value)
}

func TestElderLedger(t *testing.T) {
	tt := test.Start(t)
	tt.Scenario("base")
	defer tt.Finish()
	q := &Q{tt.FrontierSession()}

	var seq int
	err := q.ElderLedger(&seq)

	if tt.Assert.NoError(err) {
		tt.Assert.Equal(1, seq)
	}
}
