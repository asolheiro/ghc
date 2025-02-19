package metrics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	httputils "github.com/asolheiro/gita-healthcheck/internal/http_utils"
)

type Node struct {
	Name           string `json:"name"`
	UID            string `json:"uid"`
	CPUCapacity    string `json:"cpu_capacity"`
	MemoryCapacity string `json:"memory_capacity"`
	PodCapacity    string `json:"pod_capacity"`
	InUseCPU       string `json:"in_use_cpu"`
	InUseMemory    string `json:"in_use_memory"`
	PodsWithoutIP  int    `json:"podsWithoutIP"`
	PodsWithIP     int    `json:"podsWithIP"`
}

type Response struct {
	Status string `json:"status"`
	Nodes  []Node `json:"nodes"`
}

type TotalMetrics struct {
	TotalCPU            int
	TotalMemory         int
	TotalPods           int
	TotalPodCapacity    int
	TotalInUseCPU       int
	TotalInUseMemory    int
	CPUUsePercentage    float64
	MemoryUsePercentage float64
}

func GetMetrics(token, clusterId string) (TotalMetrics, error) {
	url := fmt.Sprintf("https://api-principal-geral.api.gita.cloud/dashboard/all/node/%s", clusterId)
	body, err := httputils.HttpRequest(
		nil,
		http.MethodGet,
		url,
		token,
	)

	if err != nil {
		return TotalMetrics{}, nil
	}

	var response Response
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&response); err != nil {
		return TotalMetrics{}, fmt.Errorf("error deconding JSON response, err: %s", err)
	}

	totalMetrics := calculateTotalMetrics(response)
	return totalMetrics, nil

}

func calculateTotalMetrics(response Response) TotalMetrics {
	var total TotalMetrics

	for _, node := range response.Nodes {
		cpuCapacity, _ := strconv.Atoi(node.CPUCapacity)
		memoryCapacity, _ := strconv.Atoi(strings.TrimSuffix(node.MemoryCapacity, "Ki"))
		podCapacity, _ := strconv.Atoi(node.PodCapacity)
		inUseCPU, _ := strconv.Atoi(strings.TrimSuffix(node.InUseCPU, "n"))
		inUseMemory, _ := strconv.Atoi(strings.TrimSuffix(node.InUseMemory, "Ki"))
		
		total.TotalCPU += cpuCapacity
		total.TotalMemory += memoryCapacity
		total.TotalPodCapacity += podCapacity
		total.TotalInUseCPU += inUseCPU
		total.TotalInUseMemory += inUseMemory
		total.TotalPods += node.PodsWithIP + node.PodsWithoutIP
	}

	if total.TotalCPU > 0 {
		total.CPUUsePercentage = math.Round((float64(total.TotalInUseCPU) / float64(total.TotalCPU*1e9)) * 100 * 100) / 100
	}
	if total.TotalMemory > 0 {
		total.MemoryUsePercentage = math.Round((float64(total.TotalInUseMemory) / float64(total.TotalMemory)) * 100 * 100) / 100
	}

	return total
}
