package ws

import (
	"bufio"
	"fmt"
	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/jsonutil"
	"io"
	"io/fs"
	"io/ioutil"
	"nodepanels-probe/config"
	"nodepanels-probe/log"
	"nodepanels-probe/usage"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func CallTool(command Command) {

	toolType := command.Tool.Type

	if toolType == "usage" {
		usage.InitUsageNet()
		config.InitConfigUsage()
	} else {
		//工具类不存在或者版本号不一致的话，下载工具类
		if CheckTool(command) {

			//生成入参临时文件
			tempNo := CreateTempFile(command)

			var cmd *exec.Cmd
			if envutil.IsWin() {
				cmd = exec.Command("cmd", "/C", filepath.Join(config.BinPath, "nodepanels-tool.exe")+" "+tempNo)
			}
			if envutil.IsLinux() {
				cmd = exec.Command("sh", "-c", filepath.Join(config.BinPath, "nodepanels-tool")+" "+tempNo)
			}

			stdout, _ := cmd.StdoutPipe()
			stderr, _ := cmd.StderrPipe()

			if err := cmd.Start(); err != nil {
				log.Error("Error starting command: " + fmt.Sprintf("%s", err))
			}

			go asyncLog(stdout)
			go asyncLog(stderr)

			if err := cmd.Wait(); err != nil {
				log.Error("Error waiting for command execution: " + fmt.Sprintf("%s", err))
			}

			DelCommandTempFile(tempNo)
		}
	}

}

func asyncLog(std io.ReadCloser) error {
	reader := bufio.NewReader(std)
	for {
		readString, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			return err
		}
		SendMsg(string(readString))
	}
}

// CreateTempFile 生成临时入参文件
func CreateTempFile(command Command) string {
	os.MkdirAll(filepath.Join(config.BinPath, "temp"), fs.ModeDir)
	tempNo := strconv.FormatInt(time.Now().UnixNano(), 10)
	json, _ := jsonutil.Pretty(command)
	ioutil.WriteFile(filepath.Join(config.BinPath, "temp", tempNo+".temp"), []byte(json), 0666)
	return tempNo
}

func DelCommandTempFile(tempNo string) {
	fsutil.QuietRemove(filepath.Join(config.BinPath, "temp", tempNo+".temp"))
}
