package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"nodepanels-probe/util"
)

func GetPartitionUsage() []Partition {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("get partition usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	partitionInfo, _ := disk.Partitions(false)

	partitionList := []Partition{}

	for _, val := range partitionInfo {

		usage, _ := disk.Usage(val.Mountpoint)

		partitionInfo := Partition{}
		partitionInfo.Device = val.Device
		partitionInfo.Used = usage.Used

		partitionList = append(partitionList, partitionInfo)
	}

	return partitionList
}
