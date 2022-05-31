package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/load"
	"nodepanels-probe/log"
)

func GetLoadUsage() float64 {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("Get load usage error : " + fmt.Sprintf("%s", err))
		}
	}()

	avg, _ := load.Avg()

	return avg.Load1

}
