package usage

import (
	"github.com/gookit/goutil/jsonutil"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"nodepanels-probe/config"
	"nodepanels-probe/util"
)

var initNetRx uint64 = 0
var initNetTx uint64 = 0

// InitUsageNet 初始化网络使用量
func InitUsageNet() {
	initNetTx = 0
	initNetRx = 0
	ioCounters, _ := net.IOCounters(false)
	for _, val := range ioCounters {
		initNetTx += val.BytesSent
		initNetRx += val.BytesRecv
	}
}

// Usage 获取系统使用率
func Usage() string {

	loadStat, _ := load.Avg()
	cpuStat, _ := cpu.Percent(0, false)
	swapStat, _ := mem.SwapMemory()
	memStat, _ := mem.VirtualMemory()

	//磁盘使用量
	var partitionTotal uint64 = 0
	var partitionUsed uint64 = 0

	partitionInfo, _ := disk.Partitions(false)
	for _, val := range partitionInfo {
		usage, _ := disk.Usage(val.Mountpoint)
		partitionTotal += usage.Total
		partitionUsed += usage.Used
	}

	//网络使用量
	var netTx uint64 = 0
	var netRx uint64 = 0

	ioCounters, _ := net.IOCounters(false)
	for _, val := range ioCounters {
		netTx += val.BytesSent
		netRx += val.BytesRecv
	}

	sysUsage := SysUsage{
		config.GetSid(),
		util.Round(loadStat.Load1, 2),
		util.Round(cpuStat[0], 2),
		util.Round(memStat.UsedPercent, 2),
		util.Round(swapStat.UsedPercent, 2),
		util.Round((float64(partitionUsed)/float64(partitionTotal))*100, 2),
		netRx - initNetRx,
		netTx - initNetTx,
	}

	//记录网络使用量供下次使用
	initNetTx = netTx
	initNetRx = netRx

	msg, _ := jsonutil.Encode(sysUsage)

	return string(msg)
}

type SysUsage struct {
	Sid   string  `json:"sid"`
	Load  float64 `json:"load"`
	Cpu   float64 `json:"cpu"`
	Mem   float64 `json:"mem"`
	Swap  float64 `json:"swap"`
	Disk  float64 `json:"disk"`
	NetRx uint64  `json:"netRx"`
	NetTx uint64  `json:"netTx"`
}
