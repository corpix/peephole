package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/corpix/loggers"
)

func TestMemory(t *testing.T) {
	type sample struct {
		name   string
		msg    string
		xs     []interface{}
		output [][]byte
	}

	var (
		samples = []sample{
			{
				name: "no format",
				msg:  "hello",
				xs:   nil,
				output: [][]byte{
					[]byte("hello"),
				},
			},
			{
				name: "with format",
				msg:  "hello %s",
				xs:   []interface{}{"you"},
				output: [][]byte{
					[]byte("hello you"),
				},
			},
			{
				name: "with format and multiple xs",
				msg:  "hello %s %s!",
				xs:   []interface{}{"evil", "world"},
				output: [][]byte{
					[]byte("hello evil world!"),
				},
			},
		}

		check = func(t *testing.T, s sample, l loggers.Logger) {
			assert.Equal(
				t,
				s.output,
				l.(*Memory).GetBuffer(),
			)
		}

		test = func(t *testing.T, s sample) {
			var (
				log loggers.Logger
			)

			log = New()
			log.Debug(s.msg)
			check(t, s, log)

			log = New()
			log.Print(s.msg)
			check(t, s, log)

			log = New()
			log.Error(s.msg)
			check(t, s, log)

			log = New()
			log.Fatal(s.msg)
			check(t, s, log)
		}

		testf = func(t *testing.T, s sample) {
			var (
				log loggers.Logger
			)

			log = New()
			log.Debugf(s.msg, s.xs...)
			check(t, s, log)

			log = New()
			log.Printf(s.msg, s.xs...)
			check(t, s, log)

			log = New()
			log.Errorf(s.msg, s.xs...)
			check(t, s, log)

			log = New()
			log.Fatalf(s.msg, s.xs...)
			check(t, s, log)
		}

		testWriter = func(t *testing.T, s sample) {
			log := New()
			n, err := log.Write([]byte(s.msg))

			assert.Equal(t, nil, err)
			assert.Equal(t, len(s.msg), n)
		}
	)

	for _, sample := range samples {
		t.Run(
			sample.name,
			func(t *testing.T) {
				if sample.xs == nil {
					test(t, sample)
					testWriter(t, sample)
					return
				}

				testf(t, sample)
			},
		)
	}
}
