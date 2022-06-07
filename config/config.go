package config

import (
	"encoding/json"
	"github.com/go-ping/ping"
	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/goutil/timex"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const Version = "v1.1.0"

var AgentUrl = "https://agent.nodepanels.com"
var ApiUrl = "https://collect.nodepanels.com"
var WsUrl = "wss://ws.nodepanels.com"

var BinPath = ExePath()
var C = Config{}

// InitConfig 初始化配置文件
func InitConfig() {
	f, _ := ioutil.ReadFile(filepath.Join(BinPath, "config.json"))
	json.Unmarshal(f, &C)

	C.Usage = timex.NowAddMinutes(-1).Unix()
	SetConfig()
}

func InitConfigUsage() {
	C.Usage = timex.NowAddMinutes(1).Unix()
	SetConfig()
}

// InitRequestIp 初始化请求地址
func InitRequestIp() {
	pinger, _ := ping.NewPinger(strings.Split(strings.Split(AgentUrl, "://")[1], ":")[0])
	pinger.Count = 4
	pinger.SetPrivileged(true)
	pinger.Timeout = time.Millisecond * 100
	pinger.Run()
	stats := pinger.Statistics()
	if stats.AvgRtt.Nanoseconds() == 0 {
		//如果默认请求域名延迟太高，则使用备用域名
		AgentUrl = "https://cn.agent.nodepanels.com"
		ApiUrl = "https://cn.collect.nodepanels.com"
		WsUrl = "wss://cn.ws.nodepanels.com"
	}
}

// GetSid 获取服务器id
func GetSid() string {
	return C.ServerId
}

func SetConfig() {
	json, _ := jsonutil.EncodePretty(C)
	ioutil.WriteFile(filepath.Join(BinPath, "config.json"), json, 0666)
}

type Config struct {
	ServerId string  `json:"serverId"`
	Monitor  Monitor `json:"monitor"`
	Usage    int64   `json:"usage"` //循环获取实时使用率的end时间，防止无限重复调用
}

type Monitor struct {
	Rule MonitorRule `json:"rule"`
}

type MonitorRule struct {
	Process []string `json:"process"`
}

func ExePath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return ""
	}
	return path[0 : i+1]
}
