package promquery

var PROM_COMM_URL = "http://localhost:9090/api/v1/query"

var Q_POD_SCHEDULED = "count_over_time(kube_pod_status_scheduled{namespace=***}[5m])[24h:1h]"

var Q_POD_UNSCHEDULED = "count_over_time(kube_pod_status_unschedulable{namespace=***}[5m])[24h:1h]"

var Q_CONTAINER_CPU_USAGE = "rate(container_cpu_usage_seconds_total{namespace=***}[5m])[1h:5m]"

var Q_CONTAINER_MEM_USAGE = "rate(container_memory_usage_bytes{namespace=***}[5m])[1h:30s]"

var Q_CONTAINER_FS_READ = "rate(container_fs_reads_bytes_total{namespace=***}[5m])[1h:30s]"

var Q_CONTAINER_FS_WRITE = "rate(container_fs_writes_bytes_total{namespace=***}[5m])[1h:30s]"

var Q_CONTAINER_NETWORK_RECEIVE = "rate(container_network_receive_bytes_total{namespace=***}[5m])[1h:30s]"

var Q_CONTAINER_NETWORK_TRANSMIT = "rate(container_network_transmit_bytes_total{namespace=***}[5m])[1h:30s]"

var Q_KUBELET_VOLUME_AVAILABLE = "kubelet_volume_stats_available_bytes[1h:30s]"

var Q_KUBELET_VOLUME_CAPACITY = "kubelet_volume_stats_capacity_bytes[1h:30s]"

var Q_KUBELET_VOLUME_USED = "kubelet_volume_stats_used_bytes[1h:30s]"

var Q_NODE_TEMPERATURE_CELSIUS = "node_hwmon_temp_celsius{namespace=***}"

type PQOutputFormat struct {
	DataLabel string
	Timestamp []string
	Values    []float64
}
