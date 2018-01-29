package collector

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"sync"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	log "github.com/CheerlessCloud/logrus"
)

func fetchContainerStats(ctx context.Context, cli *docker.Client, containerID string) (*types.StatsJSON, error) {
	stats := new(types.StatsJSON)
	statBuf, err := cli.ContainerStats(ctx, containerID, false)

	if err != nil {
		return stats, err
	}

	buf, err := ioutil.ReadAll(statBuf.Body)
	if err != nil {
		return stats, err
	}

	if err := json.Unmarshal(buf, &stats); err != nil {
		return stats, err
	}

	return stats, nil
}

// status - "created", "running", "paused", "restarting", "removing", "exited", or "dead"
func containerStatusTextToInt(status string) int {
	switch status {
	case "created":
		return 1
	case "running":
		return 2
	case "paused":
		return 3
	case "restarting":
		return 4
	case "removing":
		return 5
	case "exited":
		return 6
	case "dead":
		return 7
	default:
		return 0
	}
}

func fetchContainerMetrics(ctx context.Context, cli *docker.Client, container types.Container) (*ContainerMetrics, error) {
	metrics := ContainerMetrics{
		ID:        container.ID,
		StateName: container.State,
		State:     int32(containerStatusTextToInt(container.State)),
		Image:     container.Image,
	}

	if len(container.Names) < 1 {
		metrics.Name = container.ID[12:]
	} else {
		metrics.Name = container.Names[0][1:]
	}

	if metrics.Image != "" {
		if len(metrics.Image) > 7 && metrics.Image[:7] == "sha256:" {
			metrics.Image = metrics.Image[7:19]
		}
	} else if len(container.ImageID) >= 19 {
		metrics.Image = container.ImageID[7:19]
	}

	containerInfo, err := cli.ContainerInspect(ctx, container.ID)
	if err != nil {
		return nil, err
	}

	metrics.RestartCount = containerInfo.ContainerJSONBase.RestartCount
	metrics.Labels = containerInfo.Config.Labels

	if metrics.State >= 2 && metrics.State < 6 {
		stats, err := fetchContainerStats(ctx, cli, metrics.ID)
		if err != nil {
			return &metrics, err
		}

		metrics.MemoryUsage = stats.MemoryStats.Usage
		metrics.MemoryLimit = stats.MemoryStats.Limit
		if stats.MemoryStats.Usage != 0 && stats.MemoryStats.Limit != 0 {
			metrics.MemoryUsagePercent = (float32(metrics.MemoryUsage) / float32(metrics.MemoryLimit)) * 100 // TODO: round
		} else {
			metrics.MemoryUsagePercent = 0
		}

		metrics.CPUUsage = calculateCPUPercentUnix(stats)
		// TODO: it's really work?
		metrics.CPUThrottledTime = stats.CPUStats.ThrottlingData.ThrottledTime - stats.PreCPUStats.ThrottlingData.ThrottledTime
	}

	return &metrics, nil
}

// =========== start code from docker/cli =================

func calculateCPUPercentUnix(stats *types.StatsJSON) float64 {
	cpuPercent := 0.0
	// calculate the change for the cpu usage of the container in between readings
	cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage) - float64(stats.PreCPUStats.CPUUsage.TotalUsage)
	// calculate the change for the entire system between readings
	systemDelta := float64(stats.CPUStats.SystemUsage) - float64(stats.PreCPUStats.SystemUsage)
	onlineCPUs := float64(stats.CPUStats.OnlineCPUs)

	if onlineCPUs == 0.0 {
		onlineCPUs = float64(len(stats.CPUStats.CPUUsage.PercpuUsage))
	}
	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * onlineCPUs * 100.0
	}
	return cpuPercent
}

// ============= end code from docker-cli ================

// FetchMetrics - Fetch metrics for all containers
func FetchMetrics(ctx context.Context) {
	client, err := docker.NewEnvClient()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	containers, err := client.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error on fetching containers list")
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(containers))

	for _, container := range containers {
		go func(container types.Container) {
			defer wg.Done()

			if containerMetrics, err := fetchContainerMetrics(ctx, client, container); err != nil {
				log.WithFields(log.Fields{"err": err, "containerId": container.ID, "containerName": container.Names[0][1:]}).Error("Error on fetching metrics")
			} else if containerMetrics != nil {
				containerMetrics.ExportToRegistry()
			}
		}(container)
	}

	wg.Wait()

	return
}
