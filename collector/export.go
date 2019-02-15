package collector

import (
	"context"
	"math"
	"time"

	configPackage "github.com/CheerlessCloud/cexporter/config"

	_ "github.com/CheerlessCloud/cexporter/logger" // init logger

	log "github.com/CheerlessCloud/logrus"
	"github.com/prometheus/client_golang/prometheus"
)

var config = configPackage.ConfigSingleton

var labelsList = func() []string {
	labels := []string{"container_id", "container_name", "image"}

	for _, label := range config.ContainerLabelsList {
		labels = append(labels, "label_"+label)
	}

	return labels
}()

// Registry - root prometheus registry for exposing to http server
var Registry = prometheus.NewRegistry()

func defineContainerMetric(name string, help string) *prometheus.GaugeVec {
	metric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "docker",
			Subsystem: "container",
			Name:      name,
			Help:      help,
		},
		labelsList,
	)
	Registry.Register(metric)
	return metric
}

var (
	cpuUsageRatio = defineContainerMetric(
		"cpu_usage_ratio",
		"Current percent of CPU usage per container.",
	)

	memoryUsageRatio = defineContainerMetric(
		"memory_usage_ratio",
		"Current percent of memory usage per container.",
	)

	memoryUsageBytes = defineContainerMetric(
		"memory_usage_bytes",
		"Current bytes of memory usage per container.",
	)

	memoryLimitBytes = defineContainerMetric(
		"memory_limit_bytes",
		"Current limit bytes of memory usage per container.",
	)

	restartsCount = defineContainerMetric(
		"restarts_count",
		"Count of container restarts.",
	)

	cpuThrottledTime = defineContainerMetric(
		"cpu_throttled_time",
		"Count of milliseconds when trottling cpu for conrainer is enabled.",
	)

	containerState = defineContainerMetric(
		"state",
		"ID of current container state",
	)

	fetchMetricsTime = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "docker",
		Subsystem: "exporter",
		Name:      "fetch_metrics_time_ms",
		Help:      "Time of fetching all metrics from docker daemon(s)",
	}, []string{})
)

func init() {
	Registry.Register(fetchMetricsTime)
	if config.EnableSelfMetrics {
		Registry.MustRegister(prometheus.NewGoCollector())
	}
}

func round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func exportToRegistry(metrics *ContainerMetrics) {
	log.WithFields(log.Fields{"metric": metrics}).Debug("metric for container " + metrics.Name)

	labels := prometheus.Labels{
		"container_id":   metrics.ID[:12],
		"container_name": metrics.Name,
		"image":          metrics.Image,
	}

	for _, label := range config.ContainerLabelsList { // copy by list to awoid "inconsistent label cardinality"
		labels["label_"+label] = metrics.Labels[label] // labels of container will place as {...label_myLabel-foo="bar"}
	}

	markMetricAsActualInCollectItration(metrics.ID)
	rememberMetricLabels(metrics.ID, labels)

	cpuUsageRatio.With(labels).Set(round(metrics.CPUUsage, 3))
	memoryUsageRatio.With(labels).Set(round(float64(metrics.MemoryUsagePercent), 3))

	memoryUsageBytes.With(labels).Set(float64(metrics.MemoryUsage))
	memoryLimitBytes.With(labels).Set(float64(metrics.MemoryLimit))
	cpuThrottledTime.With(labels).Set(float64(metrics.CPUThrottledTime))
	restartsCount.With(labels).Set(float64(metrics.RestartCount))
	containerState.With(labels).Set(float64(metrics.State))
}

// StartCollectingMetrics - start exporting metrics to prometheus registry.
func StartCollectingMetrics(fetchInterval int64, fetchTimeout int64) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(fetchInterval))

	for _ = range ticker.C {
		ctx, ctxCancel := context.WithTimeout(context.Background(), time.Duration(fetchTimeout)*time.Millisecond)
		defer ctxCancel()

		startTime := time.Now()

		refreshUnactualMetricsList()

		FetchMetrics(ctx)

		flushUnactualMetrics()

		timeout := math.Floor(float64(time.Now().Sub(startTime).Nanoseconds() / 1000 / 1000))

		log.WithField("time", timeout).Debug("Time to fetch metrics")

		fetchMetricsTime.With(prometheus.Labels{}).Set(float64(timeout))
	}
}
