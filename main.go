package main

import (
	"Jaeger/operation"
	"fmt"
)

// welcome 函数打印欢迎信息
func welcome() {
	fmt.Println("欢迎使用Jaeger应急响应快速检索工具V0.1")
}

func main() {
	welcome() // 调用欢迎函数
	//operation.CheckSystemVersion()
	//operation.CheckLoginRecords()
	//operation.CheckCommandHistory()
	//operation.CheckSystemProcesses()
	//operation.CheckNetworkConnections()
	//operation.CheckRecentFiles()
	//operation.CheckSystemServices()
	operation.CheckSystemUsersAndGroups()
}
