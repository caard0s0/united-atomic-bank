package api

import "github.com/prometheus/client_golang/prometheus"

var (
	successfulCreatedAccounts = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "uab",
			Name:      "successful_created_accounts_total",
			Help:      "Total number of accounts successfully created",
		},
		[]string{"path", "method", "status"},
	)

	successfulLogins = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "uab",
			Name:      "successful_logins_total",
			Help:      "Total number of successful logins",
		},
		[]string{"path", "method", "status"},
	)

	successfulTransfers = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "uab",
			Name:      "successful_transfers_total",
			Help:      "Total number of successful transfers",
		},
		[]string{"path", "method", "status"},
	)
)

func init() {
	prometheus.MustRegister(successfulCreatedAccounts)
	prometheus.MustRegister(successfulLogins)
	prometheus.MustRegister(successfulTransfers)
}
