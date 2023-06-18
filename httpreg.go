package promcli

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type HTTPReg struct {
	reg *prometheus.Registry

	namespace string
	subsystem string

	requestResultCount     *prometheus.CounterVec
	requestDurationSeconds *prometheus.HistogramVec
}

func NewHTTP(namespace, subsystem string) *HTTPReg {
	httpReg := &HTTPReg{
		reg:       prometheus.NewRegistry(),
		namespace: namespace,
		subsystem: subsystem,
		requestResultCount: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "request_result_count",
				Help:      "A counter of requests results",
			},
			[]string{"path", "result"},
		),
		requestDurationSeconds: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "request_duration_seconds",
				Help:      "A histogram of latencies for requests",
				Buckets:   []float64{.0001, .00025, .0005, .001, .0025, .005, .01, .025, .05, .1, .25, .5, 1, 2, 3, 4, 5, 6, 7, 10},
			},
			[]string{"path", "code"},
		),
	}

	httpReg.reg.MustRegister(httpReg.requestResultCount)
	httpReg.reg.MustRegister(httpReg.requestDurationSeconds)

	return httpReg
}

func (h *HTTPReg) IncRequestResultCount(path string, result Label) {
	h.requestResultCount.WithLabelValues(path, result.String()).Inc()
}

func (h *HTTPReg) ObserveRequestDurationSeconds(path string, code int, duration time.Duration) {
	h.requestDurationSeconds.WithLabelValues(path, strconv.Itoa(code)).Observe(duration.Seconds())
}
