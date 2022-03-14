package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"nodepanels-probe/util"
	"runtime"
	"strings"
)

func GetDiskInfo() []DiskInfo {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("get disk info error : " + fmt.Sprintf("%s", err))
		}
	}()

	partitionInfo, _ := disk.Partitions(false)

	diskInfoList := []DiskInfo{}

	for _, val := range partitionInfo {

		usage, _ := disk.Usage(val.Mountpoint)

		diskInfo := DiskInfo{}
		diskInfo.Device = val.Device
		diskInfo.Mountpoint = val.Mountpoint
		diskInfo.Fstype = val.Fstype
		diskInfo.Total = usage.Total

		diskInfoList = append(diskInfoList, diskInfo)
	}

	//记录初始读写数据
	ioData := getDiskIOCounters()

	diskInitReadBytes = make(map[string]uint64)
	diskInitWriteBytes = make(map[string]uint64)

	for key, val := range ioData {
		diskInitReadBytes[key] = val.ReadBytes
		diskInitWriteBytes[key] = val.WriteBytes
	}

	return diskInfoList
}

func GetDiskUsage() []Disk {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("get disk io usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	ioData := getDiskIOCounters()

	diskList := []Disk{}

	for key, val := range ioData {

		read := uint64(0)
		write := uint64(0)

		if val.ReadBytes >= diskInitReadBytes[key] {
			read = val.ReadBytes - diskInitReadBytes[key]
			write = val.WriteBytes - diskInitWriteBytes[key]
		}

		disk := Disk{
			Name:  key,
			Read:  read,
			Write: write,
		}

		diskInitReadBytes[key] = val.ReadBytes
		diskInitWriteBytes[key] = val.WriteBytes

		diskList = append(diskList, disk)
	}

	return diskList
}

func getDiskIOCounters() map[string]disk.IOCountersStat {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get disk IO counters error : " + fmt.Sprintf("%s", err))
		}
	}()

	partitionInfo, _ := disk.IOCounters()
	if runtime.GOOS == "linux" {
		disks := make([]string, 0)
		diskMap := make(map[string]string)
		for _, val := range partitionInfo {
			if strings.Index(val.Name, "sd") == 0 || strings.Index(val.Name, "vd") == 0 {
				diskMap[val.Name[0:3]] = "1"
			}
		}
		for key, _ := range diskMap {
			disks = append(disks, key)
		}
		partitionInfo, _ = disk.IOCounters(disks...)
	}
	return partitionInfo
}
