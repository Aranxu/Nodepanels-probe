package probe

import (
	"fmt"
	"github.com/gookit/goutil/envutil"
	"github.com/shirou/gopsutil/v3/cpu"
	"nodepanels-probe/log"
	"nodepanels-probe/util"
	"strings"
)

func GetCpuInfo() CpuInfo {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("get cpu info error : " + fmt.Sprintf("%s", err))
		}
	}()

	infoStat, _ := cpu.Info()

	cpuNums := 0
	physicalIds := ""
	if envutil.IsLinux() {
		for _, val := range infoStat {
			if !strings.Contains(physicalIds, val.PhysicalID+",") {
				physicalIds += val.PhysicalID + ","
				cpuNums++
			}
		}
	}
	if envutil.IsWin() {
		cpuNums = len(infoStat)
	}

	physicalCores, _ := cpu.Counts(false)

	logicCore, _ := cpu.Counts(true)

	cpuInfo := CpuInfo{}
	cpuInfo.CpuNums = cpuNums
	cpuInfo.PhysicalCores = physicalCores
	cpuInfo.LogicCore = logicCore

	cpuInfo.Model = infoStat[0].ModelName
	cpuInfo.VendorID = infoStat[0].VendorID
	cpuInfo.Mhz = infoStat[0].Mhz
	cpuInfo.Cache = infoStat[0].CacheSize

	return cpuInfo
}

func GetCpuUsage() Cpu {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("get cpu usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	cpuState, _ := cpu.Percent(0, false)

	cpuUsage := Cpu{
		util.Round(cpuState[0], 2),
	}

	return cpuUsage
}
