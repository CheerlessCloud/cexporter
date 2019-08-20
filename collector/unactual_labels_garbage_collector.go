package collector

import (
    "github.com/prometheus/client_golang/prometheus"
    "sync"
    )

var (
	rememberedMetricLabels      = make(map[string]prometheus.Labels)
	rememberedMetricLabelsMutex = sync.RWMutex{}
    labelsToDelete              = make(map[string]prometheus.Labels)
    labelsToDeleteMutex         = sync.RWMutex{}
)

func refreshUnactualMetricsList() {
	labelsToDelete = make(map[string]prometheus.Labels)
	for key, value := range rememberedMetricLabels {
		labelsToDelete[key] = value
	}
}

func flushUnactualMetrics() {
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
    rememberedMetricLabelsMutex.Lock()
	rememberedMetricLabels[containerID] = labels
    rememberedMetricLabelsMutex.Unlock()
}

func markMetricAsActualInCollectItration(containerID string) {
    labelsToDeleteMutex.Lock()
	delete(labelsToDelete, containerID)
    labelsToDeleteMutex.Unlock()
}
