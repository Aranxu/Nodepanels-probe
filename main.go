package main

import (
	"github.com/gookit/goutil/strutil"
	"nodepanels-probe/config"
	"nodepanels-probe/log"
	"nodepanels-probe/probe"
	"nodepanels-probe/util"
	"nodepanels-probe/ws"
	"time"
)

//go:generate goversioninfo -arm -o=resource_windows.syso -icon=favicon.ico

func main() {
	Entry()
}

func init() {
	config.InitConfig()
	config.InitRequestIp()
	probe.InitDiskIO()
	probe.InitNet()
}

func StartProbe() {
	if strutil.IsNotBlank(config.GetSid()) {
		log.Info("Program started successfully")
		go sendServerInfo()
		go sendServerUsage()
		go ws.Connect()
		go ws.SendUsage()
	} else {
		log.Error("The program is not completely installed, please reinstall")
	}
}

//发送服务器信息
func sendServerInfo() {
	for {
		go util.PostJson(config.AgentUrl+"/api/v2", probe.GetServerInfo())

		time.Sleep(10 * time.Minute)
	}
}

//发送使用率信息
func sendServerUsage() {
	for {
		go util.PostJson(config.ApiUrl+"/api/v2", probe.GetServerUsage())

		time.Sleep(time.Minute)
	}
}
