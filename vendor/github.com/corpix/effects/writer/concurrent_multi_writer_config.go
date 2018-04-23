package writer

import (
	"github.com/corpix/pool"
	"github.com/corpix/time"
)

type BacklogConfig struct {
	// Size, determines how many writers could wait for free slots.
	Size int `default:"64" validate:"required"`

	// AddTimeout, determines how much time we could wait to write
	// into particular writer queue(when out of backlog capacity).
	// Error handler will be called if backlog queue is full and
	// this time is out.
	AddTimeout time.Duration `default:"50ms"`
}

// ConcurrentMultiWriterConfig is a configuration for ConcurrentMultiWriter.
// Backlog is a size of the writer personal buffer(if writer is slow or so).
// Pool is a configuration for github.com/corpix/pool
type ConcurrentMultiWriterConfig struct {
	// Backlog, writer backlog queue configuration.
	Backlog BacklogConfig

	// Pool consists of:
	// Workers, determines how many writers we could serve.
	// QueueSize, determines how many writes we could queue for each writer.
	Pool pool.Config `default:"{\"Workers\": 1024, \"QueueSize\": 16}" validate:"required"`
}

var (
	DefaultConcurrentMultiWriterConfig = ConcurrentMultiWriterConfig{
		Backlog: BacklogConfig{
			Size:       8,
			AddTimeout: time.Duration(10 * time.Millisecond),
		},
		Pool: pool.Config{
			Workers:   128,
			QueueSize: 8,
		},
	}
)
