package ingest

import (
	"fmt"
	"time"

	"github.com/xdbfoundation/go/ingest/io"
	"github.com/xdbfoundation/go/support/errors"
)

type processorsRunDurations map[string]time.Duration

func (d processorsRunDurations) AddRunDuration(name string, startTime time.Time) {
	d[name] += time.Since(startTime)
}

type groupChangeProcessors struct {
	processors []frontierChangeProcessor
	processorsRunDurations
}

func newGroupChangeProcessors(processors []frontierChangeProcessor) *groupChangeProcessors {
	return &groupChangeProcessors{
		processors:             processors,
		processorsRunDurations: make(map[string]time.Duration),
	}
}

func (g groupChangeProcessors) ProcessChange(change io.Change) error {
	for _, p := range g.processors {
		startTime := time.Now()
		if err := p.ProcessChange(change); err != nil {
			return errors.Wrapf(err, "error in %T.ProcessChange", p)
		}
		g.AddRunDuration(fmt.Sprintf("%T", p), startTime)
	}
	return nil
}

func (g groupChangeProcessors) Commit() error {
	for _, p := range g.processors {
		startTime := time.Now()
		if err := p.Commit(); err != nil {
			return errors.Wrapf(err, "error in %T.Commit", p)
		}
		g.AddRunDuration(fmt.Sprintf("%T", p), startTime)
	}
	return nil
}

type groupTransactionProcessors struct {
	processors []frontierTransactionProcessor
	processorsRunDurations
}

func newGroupTransactionProcessors(processors []frontierTransactionProcessor) *groupTransactionProcessors {
	return &groupTransactionProcessors{
		processors:             processors,
		processorsRunDurations: make(map[string]time.Duration),
	}
}

func (g groupTransactionProcessors) ProcessTransaction(tx io.LedgerTransaction) error {
	for _, p := range g.processors {
		startTime := time.Now()
		if err := p.ProcessTransaction(tx); err != nil {
			return errors.Wrapf(err, "error in %T.ProcessTransaction", p)
		}
		g.AddRunDuration(fmt.Sprintf("%T", p), startTime)
	}
	return nil
}

func (g groupTransactionProcessors) Commit() error {
	for _, p := range g.processors {
		startTime := time.Now()
		if err := p.Commit(); err != nil {
			return errors.Wrapf(err, "error in %T.Commit", p)
		}
		g.AddRunDuration(fmt.Sprintf("%T", p), startTime)
	}
	return nil
}
