package log

import (
	_ "embed"
	log "github.com/cihub/seelog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var BinPath = ExePath()

//go:embed log.xml
var config string

var logger, _ = log.LoggerFromConfigAsString(strings.ReplaceAll(config, "filename=\"\"", "filename=\""+filepath.Join(BinPath, "log")+"\""))

func Debug(msg string) {
	logger.Debug(msg)
}

func Info(msg string) {
	logger.Info(msg)
}

func Warn(msg string) {
	logger.Warn(msg)
}

func Error(msg string) {
	logger.Error(msg)
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
