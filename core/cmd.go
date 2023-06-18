package core

import (
	"fmt"
	"os/exec"
	"runtime"
)

func ExecCommand(command string) string {
	content := ""
	if runtime.GOOS == "windows" {
		content = ExecWin(command)
	} else if runtime.GOOS == "linux" {
		content = ExecLinux(command)
	}
	fmt.Println(content)
	return content
}

func ExecDome(command string) string {
	cmd := exec.Command(command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err.Error()
	}
	return string(output)
}

func ExecWin(command string) string {
	cmd := exec.Command("cmd", "/c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err.Error()
	}
	return string(output)
}

func ExecLinux(command string) string {
	//sh、/bin/bash
	cmd := exec.Command("/bin/bash", "/c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err.Error()
	}
	return string(output)
}
