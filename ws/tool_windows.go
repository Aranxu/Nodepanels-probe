package ws

import (
	"github.com/gookit/goutil/fsutil"
	"nodepanels-probe/config"
	"nodepanels-probe/util"
	"os/exec"
	"path/filepath"
	"runtime"
)

// CheckTool 检查工具包
func CheckTool(command Command) bool {
	version := command.Tool.Version

	if fsutil.FileExists(filepath.Join(config.BinPath, "nodepanels-tool.exe")) {
		output, _ := exec.Command("cmd", "/C", filepath.Join(config.BinPath, "nodepanels-tool.exe")+" -version").Output()
		if string(output) != version {
			return DownloadTool(command)
		} else {
			return true
		}
	} else {
		return DownloadTool(command)
	}
	return false
}

// DownloadTool 下载工具包
func DownloadTool(command Command) bool {

	defer func() {
		err := recover()
		if err != nil {
			SendMsg(PrintResult(command.Page, command.Tool.Type, "TOOLUPDATINGERROR"))
		}
	}()

	SendMsg(PrintResult(command.Page, command.Tool.Type, "TOOLUPDATINGBEGIN"))
	var url = "https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/tool/prod/" + command.Tool.Version + "/nodepanels-tool-" + runtime.GOOS + "-" + runtime.GOARCH

	util.Download(url+".exe", filepath.Join(config.BinPath, "nodepanels-tool.exe"))

	SendMsg(PrintResult(command.Page, command.Tool.Type, "TOOLUPDATINGCOMPLETE"))
	return true
}
