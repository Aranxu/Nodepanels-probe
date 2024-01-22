package ws

import (
	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/fsutil"
	"nodepanels-probe/config"
	"nodepanels-probe/util"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// CheckTool 检查工具包
func CheckTool(command Command) bool {
	version := command.Tool.Version

	if fsutil.FileExists(filepath.Join(config.BinPath, "nodepanels-tool")) {
		output, _ := exec.Command("sh", "-c", filepath.Join(config.BinPath, "nodepanels-tool")+" -version").Output()
		if string(output) != version {
			return DownloadTool(command)
		} else {
			return true
		}
	} else {
		return DownloadTool(command)
	}
}

// DownloadTool 下载工具包
func DownloadTool(command Command) bool {

	defer func() {
		err := recover()
		if err != nil {
			PrintResult(command.Page, command.Tool.Type, "TOOLUPDATINGERROR")
		}
	}()

	PrintResult(command.Page, command.Tool.Type, "TOOLUPDATINGBEGIN")
	var url = "https://file.nodepanels.com/tool/" + command.Tool.Version + "/nodepanels-tool-" + runtime.GOOS + "-" + runtime.GOARCH

	util.Download(url, filepath.Join(config.BinPath, "nodepanels-tool"))

	if envutil.IsLinux() {
		//linux系统赋予执行权限
		os.Chmod(filepath.Join(config.BinPath, "nodepanels-tool"), 0777)
	}
	PrintResult(command.Page, command.Tool.Type, "TOOLUPDATINGCOMPLETE")
	return true
}
