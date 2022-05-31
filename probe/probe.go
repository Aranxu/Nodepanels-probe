package probe

import (
	"encoding/json"
	"fmt"
	"nodepanels-probe/config"
	"nodepanels-probe/log"
)

type ProbeUsage struct {
	Cpu       Cpu         `json:"cpu"`
	Mem       Mem         `json:"mem"`
	Swap      Swap        `json:"swap"`
	Disk      []Disk      `json:"disk"`
	Partition []Partition `json:"partition"`
	Net       Net         `json:"net"`
	Process   Process     `json:"process"`
	Load      Load        `json:"load"`
}

type Cpu struct {
	Total float64 `json:"total"`
}

type Mem struct {
	Usage float64 `json:"usage"`
}

type Swap struct {
	Usage float64 `json:"usage"`
}

type Disk struct {
	Name  string `json:"name"`
	Read  uint64 `json:"read"`
	Write uint64 `json:"write"`
}

type Partition struct {
	Device string `json:"device"`
	Used   uint64 `json:"used"`
}

type Net struct {
	Rx uint64 `json:"rx"`
	Tx uint64 `json:"tx"`
}

type Process struct {
	Num         uint64         `json:"num"`
	ProcessList []ProcessUsage `json:"list"`
}

type Load struct {
	SysLoad float64 `json:"sysLoad"`
}

type ProbeInfo struct {
	Version  string     `json:"version"`
	HostInfo HostInfo   `json:"host"`
	CpuInfo  CpuInfo    `json:"cpu"`
	MemInfo  MemInfo    `json:"mem"`
	DiskInfo []DiskInfo `json:"disk"`
}

type HostInfo struct {
	Hostname        string `json:"hostname"`
	Uptime          uint64 `json:"uptime"`
	KernelArch      string `json:"kernelArch"`
	KernelVersion   string `json:"kernelVersion"`
	Os              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformFamily  string `json:"platformFamily"`
	PlatformVersion string `json:"platformVersion"`
}

type CpuInfo struct {
	CpuNums       int     `json:"num"`
	PhysicalCores int     `json:"physical"`
	LogicCore     int     `json:"logic"`
	Model         string  `json:"model"`
	VendorID      string  `json:"vendor"`
	Mhz           float64 `json:"mhz"`
	Cache         int32   `json:"cache"`
}

type MemInfo struct {
	Mem  uint64 `json:"mem"`
	Swap uint64 `json:"swap"`
}

type DiskInfo struct {
	Device     string `json:"device"`
	Mountpoint string `json:"mountpoint"`
	Fstype     string `json:"fstype"`
	Total      uint64 `json:"total"`
}

func GetServerInfo() []byte {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("Get server info data error : " + fmt.Sprintf("%s", err))
		}
	}()

	probeInfo := ProbeInfo{}
	probeInfo.Version = config.Version
	probeInfo.HostInfo = GetHostInfo()
	probeInfo.CpuInfo = GetCpuInfo()
	probeInfo.MemInfo = GetMemInfo()
	probeInfo.DiskInfo = GetDiskInfo()

	msg, _ := json.Marshal(probeInfo)

	resultMap := make(map[string]string)
	resultMap["serverId"] = config.GetSid()
	resultMap["msg"] = string(msg)
	result, _ := json.Marshal(resultMap)

	return result
}

func GetServerUsage() []byte {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("Sending usage data error : " + fmt.Sprintf("%s", err))
		}
	}()

	probeUsage := ProbeUsage{}
	probeUsage.Cpu = GetCpuUsage()
	probeUsage.Mem = GetMemUsage()
	probeUsage.Swap = GetSwapUsage()
	probeUsage.Disk = GetDiskUsage()
	probeUsage.Partition = GetPartitionUsage()
	probeUsage.Net = GetNetUsage()
	probeUsage.Process.Num = GetProcessNum()
	probeUsage.Process.ProcessList = GetProcessUsage()
	probeUsage.Load.SysLoad = GetLoadUsage()

	msg, _ := json.Marshal(probeUsage)

	resultMap := make(map[string]string)
	resultMap["serverId"] = config.GetSid()
	resultMap["msg"] = string(msg)
	result, _ := json.Marshal(resultMap)

	return result
}
