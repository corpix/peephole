package time

import (
	"time"
)

func Sleep(d Duration) {
	time.Sleep(d.Duration())
}
