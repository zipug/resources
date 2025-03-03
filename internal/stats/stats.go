package statistics

import (
	"encoding/json"
	"fmt"
	"resources/internal/models"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func GetCPUStats() (*models.CPU, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("Error getting CPU usage: %v\n", err)
		return nil, err
	}
	cores, err := cpu.Counts(false)
	if err != nil {
		fmt.Printf("Error getting CPU cores: %v\n", err)
		return nil, err
	}
	return &models.CPU{
		Usage: cpuPercent[0],
		Cores: cores,
	}, nil
}

func GetRAMStats() (*models.RAM, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error getting RAM usage: %v\n", err)
		return nil, err
	}
	return &models.RAM{
		Usage: memInfo.UsedPercent,
		Total: memInfo.Total,
		Used:  memInfo.Used,
		Free:  memInfo.Free,
	}, nil
}

func GetHDDStats() (*models.HDD, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Printf("Error getting partitions: %v\n", err)
		return nil, err
	}
	partition := partitions[0]
	usage, err := disk.Usage(partition.Mountpoint)
	if err != nil {
		fmt.Printf("Error getting usage for %v: %v\n", partition.Mountpoint, err)
		return nil, err
	}
	return &models.HDD{
		Partition: partition.Mountpoint,
		Usage:     usage.UsedPercent,
		Total:     usage.Total,
		Used:      usage.Used,
		Free:      usage.Free,
	}, nil
}

func GetAllStats() ([]byte, error) {
	cpuStats, err := GetCPUStats()
	if err != nil {
		return nil, err
	}
	ramStats, err := GetRAMStats()
	if err != nil {
		return nil, err
	}
	hddStats, err := GetHDDStats()
	if err != nil {
		return nil, err
	}
	stats := models.Stats{
		CPU: *cpuStats,
		RAM: *ramStats,
		HDD: *hddStats,
	}
	res, err := json.Marshal(stats)
	if err != nil {
		return nil, err
	}
	return res, nil
}
