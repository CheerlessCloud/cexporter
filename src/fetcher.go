package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/buger/jsonparser"
)

// ContainerMetrics - Struct with metrics of docker container
type ContainerMetrics struct {
	name               string // container name
	id                 string // container id
	image              string // container image
	restartCount       int64  // count of container restarts
	memoryUsage        int64  // memory usage in bytes
	momoryLimit        int64  // memory limit for container in bytes
	memoryUsagePercent int64  // percent of usage memory in container
	cpuUsage           int64  // percent of usage CPU
	cpuThrottledTime   int64  // aggregate time the container was throttled for in nanoseconds.
	// TODO: implement block/io and network/io
}

func fetchRestartCount(cli docker.Client, containerID string) (int, error) {
	inspect, err := cli.ContainerInspect(context.Background(), containerID)

	if err != nil {
		return 0, err
	}

	return inspect.ContainerJSONBase.RestartCount, nil
}

// type containerStats struct {
// 	Name            string    `json:"name"`
// 	Id              string    `json:"id"`
// 	Name            string    `json:"name"`
// 	ReadDateTime    time.Time `json:"read"`
// 	PreReadDateTime time.Time `json:"preread"`
// }

// type containerCPUStats struct {
// 	Usage struct {
// 		Total        uint64   `json:"total_usage"`
// 		PerCPU       []uint64 `json:"percpu_usage"`
// 		InKernelMode uint64   `json:"usage_in_kernelmode"`
// 		InUserSpace  uint64   `json:"usage_in_usermode"`
// 	} `json:"cpu_usage"`
// }

func fetchStat(cli docker.Client, containerID string) (*ContainerMetrics, error) {
	met1 := new(types.StatsJSON)
	metrics := new(ContainerMetrics)
	statBuf, err := cli.ContainerStats(context.Background(), containerID, false)

	if err != nil {
		return metrics, err
	}

	buf, err := ioutil.ReadAll(statBuf.Body)
	if err == nil {
		return metrics, err
	}

	if metrics.memoryUsage, err = jsonparser.GetInt(buf, "memory_stats", "usage"); err != nil {
		return metrics, err
	}

	if metrics.momoryLimit, err = jsonparser.GetInt(buf, "memory_stats", "limit"); err != nil {
		return metrics, err
	}

	metrics.memoryUsagePercent = metrics.memoryUsage / metrics.momoryLimit * 100

	return metrics, nil
}

// FetchMetrics - Fetch metrics per container from docker daemon
func FetchMetrics(dockerURL string) {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		stat, err := cli.ContainerStats(context.Background(), container.ID, false)

		if err != nil {
			panic(err)
		}

		if b, err := ioutil.ReadAll(stat.Body); err == nil {
			fmt.Println(string(b))
		}

		inspect, err := cli.ContainerInspect(context.Background(), container.ID)

		if err != nil {
			panic(err)
		}

		fmt.Println(inspect.ContainerJSONBase.RestartCount)

		// var buf []byte
		// bytesCount, err := stat.Body.Read(buf)
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("%s %s\n", string(bytesCount), buf)
		// fmt.Printf("%s %s %s %s\n", container.ID[:10], container.Image, string(bytesCount), buf)
	}
}
