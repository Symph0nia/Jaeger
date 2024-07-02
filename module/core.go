package core

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ANSI color codes
const (
	Blue   = "\033[34m"
	Green  = "\033[32m"
	Orange = "\033[33m"
	Purple = "\033[35m"
	Reset  = "\033[0m"
)

// ExecuteCommand 执行一个系统命令并返回其输出或错误，同时负责彩色打印输出、错误信息和描述信息
func ExecuteCommand(command string, args []string, description string) (string, error) {
	// 使用紫色打印命令的描述
	fmt.Printf("%s[*] %s%s\n", Purple, description, Reset)

	// 使用蓝色打印正在执行的命令
	cmdLine := command + " " + strings.Join(args, " ")
	fmt.Printf("%s[*] 正在执行命令：%s%s\n", Blue, cmdLine, Reset)

	// 使用bash -c来允许管道和其他shell特性
	cmd := exec.Command("bash", "-c", cmdLine)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		errorMsg := fmt.Sprintf("%s[-] 命令执行失败：%s，错误：%s%s", Orange, err, stderr.String(), Reset)
		fmt.Println(errorMsg) // 使用橙色打印错误信息
		return "", fmt.Errorf(errorMsg)
	}

	output := out.String()
	fmt.Printf("%s[+] 命令输出：\n%s%s", Green, output, Reset) // 使用绿色打印成功的输出

	// 等待用户按回车键
	fmt.Print("按回车键继续...")
	fmt.Scanln()

	return output, nil
}
