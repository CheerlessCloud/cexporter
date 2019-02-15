package collector

import "github.com/prometheus/client_golang/prometheus"

var (
	rememberedMetricLabels = make(map[string]prometheus.Labels)
	labelsToDelete         = make(map[string]prometheus.Labels)
)

func initLabelsToDelete() {
	labelsToDelete = make(map[string]prometheus.Labels)
	for key, value := range rememberedMetricLabels {
		labelsToDelete[key] = value
	}
}

func flushAllMetrics() {
	for key, labels := range labelsToDelete {
		cpuUsageRatio.Delete(labels)
		memoryUsageRatio.Delete(labels)
		memoryUsageBytes.Delete(labels)
		memoryLimitBytes.Delete(labels)
		cpuThrottledTime.Delete(labels)
		restartsCount.Delete(labels)
		containerState.Delete(labels)
		delete(labelsToDelete, key)
		delete(rememberedMetricLabels, key)
	}
}

func rememberMetricLabels(containerID string, labels prometheus.Labels) {
	rememberedMetricLabels[containerID] = labels
}

func markMetricAsActualInCollectItration(containerID string) {
	delete(labelsToDelete, containerID)
}
