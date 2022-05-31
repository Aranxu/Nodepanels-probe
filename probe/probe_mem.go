package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"nodepanels-probe/log"
	"nodepanels-probe/util"
)

func GetMemInfo() MemInfo {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("Get mem info error : " + fmt.Sprintf("%s", err))
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
			log.Error("Get mem usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	virtualMemoryStat, _ := mem.VirtualMemory()

	mem := Mem{}

	mem.Usage = util.Round(virtualMemoryStat.UsedPercent, 2)

	return mem
}

func GetSwapUsage() Swap {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("Get swap usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	swap := Swap{}

	swapMemory, _ := mem.SwapMemory()

	swap.Usage = util.Round(swapMemory.UsedPercent, 2)

	return swap
}
