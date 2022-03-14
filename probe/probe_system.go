package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"nodepanels-probe/util"
)

func GetHostInfo() HostInfo {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get host info error : " + fmt.Sprintf("%s", err))
		}
	}()

	infoStat, _ := host.Info()
	uptime, _ := host.Uptime()
	kernelArch, _ := host.KernelArch()
	kernelVersion, _ := host.KernelVersion()
	platform, platformFamily, platformVersion, _ := host.PlatformInformation()

	hostInfo := HostInfo{}
	hostInfo.Hostname = infoStat.Hostname
	hostInfo.Uptime = uptime
	hostInfo.KernelArch = kernelArch
	hostInfo.KernelVersion = kernelVersion
	hostInfo.Os = infoStat.OS
	hostInfo.Platform = platform
	hostInfo.PlatformFamily = platformFamily
	hostInfo.PlatformVersion = platformVersion

	return hostInfo
}
