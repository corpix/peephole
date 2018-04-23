package pool

import (
	"fmt"
)

type ErrNoSlotsAvailable struct {
	QueueSize int
}

func (e ErrNoSlotsAvailable) Error() string {
	return fmt.Sprintf(
		"No free slots available in a pool at the moment, total slots %d",
		e.QueueSize,
	)
}

func NewErrNoSlotsAvailable(queueSize int) ErrNoSlotsAvailable {
	return ErrNoSlotsAvailable{
		QueueSize: queueSize,
	}
}
