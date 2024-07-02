package operation

import core "Jaeger/module"

// CheckNetworkConnections 检查系统的所有网络通信
func CheckNetworkConnections() {
	commands := []struct {
		Command     string
		Args        []string
		Description string
	}{
		{"netstat", []string{"-tulnp"}, "检查所有监听的端口和网络连接"},
		{"ss", []string{"-tulnp"}, "使用ss命令检查所有监听的端口和网络连接"},
		{"lsof", []string{"-i"}, "使用lsof命令列出打开的网络端口"},
		{"ip", []string{"a"}, "显示所有网络接口及其状态"},
		{"ip", []string{"route"}, "显示路由表信息"},
		{"ifconfig", nil, "使用ifconfig命令显示所有网络接口及其配置"},
	}

	for _, cmd := range commands {
		_, err := core.ExecuteCommand(cmd.Command, cmd.Args, cmd.Description)
		if err != nil {
			continue // 错误处理已在core.ExecuteCommand中完成，这里不再处理
		}
	}
}
