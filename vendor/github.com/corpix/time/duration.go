package time

import (
	"encoding/json"
	"time"
)

// XXX: All this mess here just to make parsing of
// time.Duration with JSON parser less painfull.

const (
	Nanosecond  = Duration(time.Nanosecond)
	Microsecond = Duration(time.Microsecond)
	Millisecond = Duration(time.Millisecond)
	Second      = Duration(time.Second)
	Minute      = Duration(time.Minute)
	Hour        = Duration(time.Hour)
)

// Duration is a wrapper around time.Duration which
// implements json.Unmarshaler and json.Marshaler.
// It marshals and unmarshals the duration as a string
// in the format accepted by time.ParseDuration and
// returned by time.Duration.String.
type Duration time.Duration

func (d Duration) String() string {
	return time.Duration(d).String()
}

func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

// MarshalJSON implements the json.Marshaler interface.
// The duration is a quoted-string in the format accepted
// by time.ParseDuration and returned by time.Duration.String.
func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The duration is expected to be a quoted-string of a duration
// in the format accepted by time.ParseDuration.
func (d *Duration) UnmarshalJSON(buf []byte) error {
	var (
		s   string
		t   time.Duration
		err error
	)

	err = json.Unmarshal(buf, &s)
	if err != nil {
		return err
	}

	t, err = time.ParseDuration(s)
	if err != nil {
		return err
	}

	*d = Duration(t)

	return nil
}
