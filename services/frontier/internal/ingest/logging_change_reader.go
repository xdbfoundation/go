package ingest

import (
	"runtime"

	"github.com/xdbfoundation/go/ingest/io"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func getMemStats() (uint64, uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return bToMb(m.HeapAlloc), bToMb(m.HeapSys)
}

type loggingChangeReader struct {
	io.ChangeReader
	profile    bool
	entryCount int
	// how often should the logger report
	frequency int
	source    string
	sequence  uint32
}

func newloggingChangeReader(
	reader io.ChangeReader,
	source string,
	sequence uint32,
	every int,
	profile bool,
) *loggingChangeReader {
	return &loggingChangeReader{
		ChangeReader: reader,
		frequency:    every,
		profile:      profile,
		source:       source,
		sequence:     sequence,
	}
}

// Read returns a new ledger entry change on each call, returning io.EOF when the stream ends.
func (lcr *loggingChangeReader) Read() (io.Change, error) {
	change, err := lcr.ChangeReader.Read()

	if err == nil {
		lcr.entryCount++

		if lcr.entryCount%lcr.frequency == 0 {
			logger := log.WithField("numEntries", lcr.entryCount).
				WithField("source", lcr.source).
				WithField("sequence", lcr.sequence)

			if lcr.profile {
				curHeap, sysHeap := getMemStats()
				logger = logger.
					WithField("currentHeapSizeMB", curHeap).
					WithField("systemHeapSizeMB", sysHeap)
			}
			logger.Info("Processing ledger entry changes")
		}
	}

	return change, err
}
