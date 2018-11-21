package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/mitchellh/go-ps"
)

func main() {
	for {
		processes, err := ps.Processes()
		if err != nil {
			fmt.Printf("Can not get process list, %v\n", err)
		}
		easyConnectExist := false
		monitorExist := false
		for _, process := range processes {
			if strings.Contains(process.Executable(), "EasyConnect") {
				easyConnectExist = true
			}
			if strings.Contains(process.Executable(), "EasyMonitor") {
				monitorExist = true
			}
		}

		if easyConnectExist && !monitorExist {
			cmd := exec.Command("launchctl", "load", "/Library/LaunchDaemons/com.sangfor.EasyMonitor.plist")
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Failed launching monitor, %v", err)
			}
		} else if !easyConnectExist && monitorExist {
			cmd := exec.Command("launchctl", "unload", "/Library/LaunchDaemons/com.sangfor.EasyMonitor.plist")
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Failed stoping monitor, %v", err)
			}
		}

		time.Sleep(time.Second)
	}
}
