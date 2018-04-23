package writer

import (
	"fmt"
	"io"
)

type ErrBacklogOverflow struct {
	Slots  int
	Writer io.Writer
}

func (e ErrBacklogOverflow) Error() string {
	return fmt.Sprintf(
		"No free slots available in writer %p backlog(slow writer?), slots '%d'",
		e.Writer,
		e.Slots,
	)
}

func NewErrBacklogOverflow(slots int, w io.Writer) ErrBacklogOverflow {
	return ErrBacklogOverflow{
		Slots:  slots,
		Writer: w,
	}
}
