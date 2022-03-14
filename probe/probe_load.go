package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/load"
	"nodepanels-probe/util"
)

func GetLoadUsage() float64 {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get load usage error : " + fmt.Sprintf("%s", err))
		}
	}()

	avg, _ := load.Avg()

	return avg.Load1

}
