package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/process"
	"nodepanels-probe/config"
	"nodepanels-probe/log"
	"nodepanels-probe/util"
)

func GetProcessUsage() []ProcessUsage {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("Get process usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	if initProcessUsageMap == nil {
		initProcessUsageMap = make(map[string]ProcessUsage)
	}

	processCmdList := config.GetConfig().Monitor.Rule.Process
	processCmdMap := make(map[string]int)
	for _, val := range processCmdList {
		processCmdMap[val] = 1
	}

	//临时map
	initProcessUsageTempMap := make(map[string]ProcessUsage)

	processUsageList := []ProcessUsage{}

	processes, _ := process.Processes()
	for _, val := range processes {

		cmd, _ := val.Cmdline()

		//判断是否需要监控
		if _, ok := processCmdMap[cmd]; ok {

			pid := val.Pid
			name, _ := val.Name()
			cpuPercent, _ := val.CPUPercent()
			memPercent, _ := val.MemoryPercent()
			ioData, _ := val.IOCounters()

			processUsage := ProcessUsage{}
			processUsage.Pid = pid
			processUsage.Name = name
			processUsage.Cmd = cmd
			processUsage.CpuPercent = util.Round(cpuPercent, 1)
			processUsage.MemPercent = util.Round(float64(memPercent), 1)

			if _, ok := initProcessUsageMap[cmd]; ok {
				if ioData != nil {
					if ioData.ReadBytes > initProcessUsageMap[cmd].DiskRead {
						processUsage.DiskRead = ioData.ReadBytes - initProcessUsageMap[cmd].DiskRead
					} else {
						processUsage.DiskRead = 0
					}
					if ioData.WriteBytes > initProcessUsageMap[cmd].DiskWrite {
						processUsage.DiskWrite = ioData.WriteBytes - initProcessUsageMap[cmd].DiskWrite
					} else {
						processUsage.DiskWrite = 0
					}
				}
			} else {
				if ioData != nil {
					processUsage.DiskRead = ioData.ReadBytes
					processUsage.DiskWrite = ioData.WriteBytes
				}
			}
			processUsageList = append(processUsageList, processUsage)

			if ioData != nil {
				initProcessUsageTempMap[cmd] = ProcessUsage{
					DiskRead:  ioData.ReadBytes,
					DiskWrite: ioData.WriteBytes,
				}
			}
		}
	}

	initProcessUsageMap = initProcessUsageTempMap

	return processUsageList
}

func GetProcessNum() uint64 {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("Get process num error : " + fmt.Sprintf("%s", err))
		}
	}()

	infoStat, _ := host.Info()

	return infoStat.Procs
}

type ProcessUsage struct {
	Name       string  `json:"name"`
	CpuPercent float64 `json:"cpu"`
	MemPercent float64 `json:"mem"`
	DiskWrite  uint64  `json:"write"`
	DiskRead   uint64  `json:"read"`
	Cmd        string  `json:"cmd"`
	Pid        int32   `json:"pid"`
}

type ProcessUsageSlice []ProcessUsage

func (p ProcessUsageSlice) Len() int {
	return len(p)
}

func (p ProcessUsageSlice) Less(i, j int) bool {
	if p[i].CpuPercent == p[j].CpuPercent {
		return p[i].MemPercent > p[j].MemPercent
	} else {
		return p[i].CpuPercent > p[j].CpuPercent
	}
}

func (p ProcessUsageSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
