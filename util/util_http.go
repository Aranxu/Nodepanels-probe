package util

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"nodepanels-probe/log"
	"os"
	"unsafe"
)

func PostJson(url string, jsonParam []byte) string {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("Post json error : " + fmt.Sprintf("%s", err))
		}
	}()

	request, err := http.NewRequest("POST", url, bytes.NewReader(jsonParam))
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	return *str
}

func Download(url string, target string) {

	defer func() {
		err := recover()
		if err != nil {
			log.Error("Download file error : " + fmt.Sprintf("%s", err))
		}
	}()

	res, _ := http.Get(url)
	newFile, _ := os.Create(target)
	io.Copy(newFile, res.Body)
	defer res.Body.Close()
	defer newFile.Close()
}
