package metrics

import (
	metrics "github.com/armon/go-metrics"

	"github.com/corpix/peephole/pkg/log"
)

type Metrics = metrics.Metrics

func Create(c Config, l log.Logger) (*Metrics, error) {
	var (
		statsdAddressesCount = len(c.StatsdAddresses)

		s   metrics.MetricSink
		err error
	)

	switch {
	case statsdAddressesCount == 0:
		s = &metrics.BlackholeSink{}
		l.Warn().Msg("will NOT report any metrics, no metrics endpoint configured")
	case statsdAddressesCount == 1:
		s, err = metrics.NewStatsdSink(c.StatsdAddresses[0])
		if err != nil {
			return nil, err
		}
		l.Info().Msgf("will report metrics to %v", c.StatsdAddresses)
	default:
		fanoutSink := make(metrics.FanoutSink, len(c.StatsdAddresses))
		for k, v := range c.StatsdAddresses {
			s, err = metrics.NewStatsdSink(v)
			if err != nil {
				return nil, err
			}
			fanoutSink[k] = s
		}
		s = fanoutSink
		l.Info().Msgf("will report metrics to %v", c.StatsdAddresses)
	}

	return metrics.New(
		metrics.DefaultConfig(c.ServiceName),
		s,
	)
}
