package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"nodepanels-probe/util"
	"runtime"
	"strings"
)

func GetCpuInfo() CpuInfo {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("get cpu info error : " + fmt.Sprintf("%s", err))
		}
	}()

	infoStat, _ := cpu.Info()

	cpuNums := 0
	physicalIds := ""
	if runtime.GOOS == "linux" {
		for _, val := range infoStat {
			if !strings.Contains(physicalIds, val.PhysicalID+",") {
				physicalIds += val.PhysicalID + ","
				cpuNums++
			}
		}
	}
	if runtime.GOOS == "windows" {
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
			util.LogError("get cpu usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	cpuUsage := Cpu{}

	combineCpuUsage, _ := cpu.Percent(0, false)

	for _, ccu := range combineCpuUsage {
		cpuUsage.Total = util.RoundFloat64(ccu, 2)
		JudgeCpuWarning(util.String2int(fmt.Sprintf("%.0f", ccu)))
	}

	logicCore, _ := cpu.Counts(true)
	if logicCore != 1 {
		perCpuUsage, _ := cpu.Percent(0, true)
		perCpuList := []float64{}
		for _, pcu := range perCpuUsage {
			perCpuList = append(perCpuList, util.RoundFloat64(pcu, 1))
		}
		cpuUsage.Per = perCpuList
	}

	return cpuUsage
}
