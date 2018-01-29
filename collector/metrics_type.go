package collector

// ContainerMetrics - Struct with metrics of docker container
type ContainerMetrics struct {
	StateName          string            `json:"stateName"`          // current container state, represent as string
	State              int32             `json:"state"`              // current container state, represent as int
	Name               string            `json:"name"`               // container name
	ID                 string            `json:"id"`                 // container id
	Image              string            `json:"image"`              // tag of image
	Labels             map[string]string `json:"labels"`             // container labels map
	RestartCount       int               `json:"restartCount"`       // count of container restarts
	MemoryUsage        uint64            `json:"memoryUsage"`        // memory usage in bytes
	MemoryLimit        uint64            `json:"memoryLimit"`        // memory limit for container in bytes
	MemoryUsagePercent float32           `json:"memoryUsagePercent"` // percent of usage memory in container
	CPUUsage           float64           `json:"cpuUsage"`           // percent of usage CPU
	CPUThrottledTime   uint64            `json:"cpuThrottledTime"`   // aggregate time the container was throttled for in nanoseconds.
	// TODO: implement block/io and network/io
}

// ExportToRegistry - Export this metrics to prometheus registry
func (m *ContainerMetrics) ExportToRegistry() {
	exportToRegistry(m)
}
