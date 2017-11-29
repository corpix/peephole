package prefixwrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/memory"
)

func TestPrefixWrapper(t *testing.T) {
	type sample struct {
		name   string
		prefix string
		msg    string
		xs     []interface{}
		output [][]byte
	}

	var (
		samples = []sample{
			{
				name:   "no format",
				prefix: "prefix: ",
				msg:    "hello",
				xs:     nil,
				output: [][]byte{
					[]byte("prefix: hello"),
				},
			},
			{
				name:   "with format",
				prefix: "prefix: ",
				msg:    "hello %s",
				xs:     []interface{}{"you"},
				output: [][]byte{
					[]byte("prefix: hello you"),
				},
			},
			{
				name:   "with format and multiple xs",
				prefix: "prefix: ",
				msg:    "hello %s %s!",
				xs:     []interface{}{"evil", "world"},
				output: [][]byte{
					[]byte("prefix: hello evil world!"),
				},
			},
			{
				name:   "with format and array xs",
				prefix: "prefix: ",
				msg:    "hello %#v!",
				xs:     []interface{}{[]string{"evil", "world"}},
				output: [][]byte{
					[]byte(`prefix: hello []string{"evil", "world"}!`),
				},
			},
		}

		newLogger = func(s sample) loggers.Logger {
			return New(
				s.prefix,
				memory.New(),
			)
		}

		check = func(t *testing.T, s sample, l loggers.Logger) {
			assert.Equal(
				t,
				s.output,
				l.(*PrefixWrapper).log.(*memory.Memory).GetBuffer(),
			)
		}

		test = func(t *testing.T, s sample) {
			var (
				log loggers.Logger
			)

			log = newLogger(s)
			log.Debug(s.msg)
			check(t, s, log)

			log = newLogger(s)
			log.Print(s.msg)
			check(t, s, log)

			log = newLogger(s)
			log.Error(s.msg)
			check(t, s, log)

			log = newLogger(s)
			log.Fatal(s.msg)
			check(t, s, log)
		}

		testf = func(t *testing.T, s sample) {
			var (
				log loggers.Logger
			)

			log = newLogger(s)
			log.Debugf(s.msg, s.xs...)
			check(t, s, log)

			log = newLogger(s)
			log.Printf(s.msg, s.xs...)
			check(t, s, log)

			log = newLogger(s)
			log.Errorf(s.msg, s.xs...)
			check(t, s, log)

			log = newLogger(s)
			log.Fatalf(s.msg, s.xs...)
			check(t, s, log)
		}

		testWriter = func(t *testing.T, s sample) {
			log := newLogger(s)
			n, err := log.Write([]byte(s.msg))

			assert.Equal(t, nil, err)
			assert.Equal(t, len(s.prefix+s.msg), n)
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
