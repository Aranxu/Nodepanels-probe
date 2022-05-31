package main

import (
	_ "embed"
	"fmt"
	"github.com/gookit/goutil/envutil"
	"github.com/kardianos/service"
	"nodepanels-probe/config"
	"nodepanels-probe/log"
	"os"
)

//go:embed logo
var logo string

// Entry 程序入口。根据入参执行安装，卸载，版本查看，帮助等功能
func Entry() {
	serviceName := ""
	if envutil.IsWin() {
		serviceName = "Nodepanels-probe"
	}
	if envutil.IsLinux() {
		serviceName = "nodepanels"
	}

	serConfig := &service.Config{
		Name:        serviceName,
		DisplayName: "Nodepanels-probe",
		Description: "Nodepanels探针进程",
	}

	pro := &Program{}
	s, err := service.New(pro, serConfig)
	if err != nil {
		fmt.Println(err, "Create service failed")
	}

	if len(os.Args) > 1 {

		if os.Args[1] == "-install" {
			err = s.Install()
			if err != nil {
				fmt.Println("Install failed", err)
			} else {
				fmt.Println("Install success")
			}
			return
		}

		if os.Args[1] == "-uninstall" {
			err = s.Uninstall()
			if err != nil {
				fmt.Println("Uninstall failed", err)
			} else {
				fmt.Println("Uninstall success")
			}
			return
		}

		if os.Args[1] == "-version" {

			fmt.Println(logo)
			fmt.Println("====================================")
			fmt.Println("App name    : nodepanels-probe")
			fmt.Println("Version     : " + config.Version)
			fmt.Println("Update Time : 20220525")
			fmt.Println("")
			fmt.Println("Made by     : https://nodepanels.com")
			fmt.Println("====================================")
			return
		}

		if os.Args[1] == "-help" {
			fmt.Println("Available option is :")
			fmt.Println("1) -install    :  Install nodepanels-probe as a service to system.")
			fmt.Println("2) -uninstall  :  Remove nodepanels-probe service from system.")
			fmt.Println("3) -version    :  Check nodepanels-probe version info.")
			fmt.Println("4) -help       :  How to use nodepanels-probe.")
			return
		}
	}

	err = s.Run()
	if err != nil {
		log.Error("Run nodepanels-probe failed" + fmt.Sprintf("%s", err))
	}
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *Program) run() {
	log.Info(logo)
	StartProbe()
}

func (p *Program) Stop(s service.Service) error {
	return nil
}
