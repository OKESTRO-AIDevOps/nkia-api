package promquery

var PROM_COMM_URL = "http://localhost:9090/api/v1/query"

var PROM_CPU_USAGE = "rate(container_cpu_usage_seconds_total{namespace=***}[5m])[1h:5m]"

type PQOutputFormat struct {
	DataLabel string
	Timestamp []string
	Values    []float64
}
