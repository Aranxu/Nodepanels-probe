package probe

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"io/ioutil"
	"net/http"
	"nodepanels-probe/util"
)

func GetNetInfo() NetInfo {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get net info error : " + fmt.Sprintf("%s", err))
		}
	}()

	ioCounters, _ := net.IOCounters(false)
	//初始化流量信息
	netInitRxBytes = ioCounters[0].BytesRecv
	netInitTxBytes = ioCounters[0].BytesSent

	url := "http://ip-api.com/json?fields=status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,isp,org,as,asname,query&lang=zh-CN"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	detailInfo := IpInfo{}
	json.Unmarshal(body, &detailInfo)

	privateIp := util.GetPrivateIp()

	netInfo := NetInfo{}
	netInfo.PublicIp = detailInfo.Query
	netInfo.PrivateIp = privateIp
	netInfo.DetailInfo = detailInfo

	netInfo.AgentIp = util.GetAgentIp()
	netInfo.ApiIp = util.GetApiIp()

	return netInfo
}

func GetNetUsage() Net {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("get net usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	netUsage := Net{}

	ioCounters, _ := net.IOCounters(false)

	if ioCounters[0].BytesRecv > netInitRxBytes {
		netUsage.Rx = ioCounters[0].BytesRecv - netInitRxBytes
	} else {
		netUsage.Rx = 0
	}
	if ioCounters[0].BytesSent > netInitTxBytes {
		netUsage.Tx = ioCounters[0].BytesSent - netInitTxBytes
	} else {
		netUsage.Tx = 0
	}

	netInitRxBytes = ioCounters[0].BytesRecv
	netInitTxBytes = ioCounters[0].BytesSent

	return netUsage

}
