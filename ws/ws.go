package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/timex"
	"github.com/gorilla/websocket"
	"nodepanels-probe/config"
	"nodepanels-probe/log"
	"nodepanels-probe/usage"
	"strings"
	"time"
)

var WebsocketConn *websocket.Conn = nil

func Connect() {
	defer func() {
		err := recover()
		if err != nil {
			log.Error("Abnormal connection with proxy program：" + fmt.Sprintf("%s", err))
		}
	}()

	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

TryConn:

	log.Info("Try to establish a connection with the proxy server...")
	wsConnect, _, err := dialer.Dial(config.WsUrl+"/ws/v1/"+config.GetSid(), nil)
	WebsocketConn = wsConnect
	if nil != err {
		fmt.Println(err)
		log.Info("Failed to connect to the proxy server, will retry after 10 seconds ")
		time.Sleep(time.Second * 10)
		goto TryConn
	}
	log.Info("Successfully connected to the proxy server ")

	go readMsg()

	go heartBeat()

	return
}

func heartBeat() {
	defer func() {
		err := recover()
		if err != nil {
			log.Error("Send ws heartbeat error :" + fmt.Sprintf("%s", err))
		}
	}()

	for {
		if WebsocketConn == nil {
			return
		}

		SendMsg("ping")
		time.Sleep(30000 * time.Millisecond)
	}
}

func readMsg() {
	defer func() {
		err := recover()
		if err != nil {
			log.Error("Receive agent message error :" + fmt.Sprintf("%s", err))
		}
	}()

	for {
		if WebsocketConn == nil {
			log.Info("Disconnect from proxy server, will retry after 10 seconds ")
			time.Sleep(time.Second * 10)
			go Connect()
			return
		}

		messageType, messageData, err := WebsocketConn.ReadMessage()
		if nil != err {
			WebsocketConn = nil
		}

		switch messageType {
		case websocket.TextMessage:
			go handleMsg(string(messageData))

		default:
			SendMsg("bad request")
		}
	}
}

func SendMsg(msg string) {
	defer func() {
		err := recover()
		if err != nil {
			log.Error("SendMsg error :" + fmt.Sprintf("%s", err))
		}
	}()

	if WebsocketConn != nil {
		WebsocketConn.WriteMessage(websocket.TextMessage, []byte(msg))
	}
}

func handleMsg(message string) {
	defer func() {
		err := recover()
		if err != nil {
			log.Error("HandleMsg error :" + fmt.Sprintf("%s", err))
		}
	}()

	if message != "pong" {
		command := Command{}
		json.Unmarshal([]byte(message), &command)

		CallTool(command)
	}
}

// SendUsage 新开线程循环获取系统使用率(2秒粒度)
func SendUsage() {
	for {
		if WebsocketConn != nil && config.C.Usage > timex.NowUnix() {
			go SendMsg(PrintResult(config.GetSid(), "usage", usage.Usage()))
		}
		time.Sleep(2 * time.Second)
	}
}

func PrintResult(pid string, toolType string, msg string) string {
	msg = strings.ReplaceAll(msg, "\\", "\\\\")
	msg = strings.ReplaceAll(msg, "\n", "\\n")
	msg = strings.ReplaceAll(msg, "\"", "\\\"")
	return "{\"pid\":\"" + pid + "\"," + "\"toolType\":\"" + toolType + "\",\"serverId\":\"" + config.GetSid() + "\",\"msg\":\"" + msg + "\"}"
}

type Command struct {
	ServerId string      `json:"serverId"`
	Page     string      `json:"page"`
	Tool     CommandTool `json:"tool"`
}

type CommandTool struct {
	Version string `json:"version"`
	Type    string `json:"type"`
	Param   string `json:"param"`
}
