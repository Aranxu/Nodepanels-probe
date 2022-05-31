package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"nodepanels-probe/log"
)

func InitNet() {
	//初始化流量信息
	netInitTxBytes = 0
	netInitRxBytes = 0
	ioCounters, _ := net.IOCounters(false)
	for _, val := range ioCounters {
		netInitTxBytes += val.BytesSent
		netInitRxBytes += val.BytesRecv
	}
}

func GetNetUsage() Net {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("get net usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	netUsage := Net{}

	//网络使用量
	var netTx uint64 = 0
	var netRx uint64 = 0

	ioCounters, _ := net.IOCounters(false)
	for _, val := range ioCounters {
		netTx += val.BytesSent
		netRx += val.BytesRecv
	}
	netUsage.Tx = netTx - netInitTxBytes
	netUsage.Rx = netRx - netInitRxBytes

	netInitTxBytes = netTx
	netInitRxBytes = netRx

	return netUsage

}
