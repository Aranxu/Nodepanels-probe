package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"nodepanels-probe/util"
)

func GetMemInfo() MemInfo {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get mem info error : " + fmt.Sprintf("%s", err))
		}
	}()

	virtualMemoryStat, _ := mem.VirtualMemory()
	swapMemory, _ := mem.SwapMemory()

	memInfo := MemInfo{}
	memInfo.Mem = virtualMemoryStat.Total
	memInfo.Swap = swapMemory.Total

	//全局变量，给进程计算实际使用内存
	memTotal = uint(virtualMemoryStat.Total / 1024 / 1024)

	return memInfo
}

func GetMemUsage() Mem {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get mem usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	virtualMemoryStat, _ := mem.VirtualMemory()

	JudgeMemWarning(util.String2int(util.Float642string(util.RoundFloat64(virtualMemoryStat.UsedPercent, 0))))

	mem := Mem{}

	mem.Usage = util.RoundFloat64(virtualMemoryStat.UsedPercent, 2)

	return mem
}

func GetSwapUsage() Swap {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get swap usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	swap := Swap{}

	swapMemory, _ := mem.SwapMemory()

	swap.Usage = util.RoundFloat64(swapMemory.UsedPercent, 2)

	return swap
}
