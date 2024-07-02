package operation

import core "Jaeger/module"

// CheckLoginRecords 检查wtmp、btmp和其他服务的登录记录
func CheckLoginRecords() {
	commands := []struct {
		Command     string
		Args        []string
		Description string
	}{
		{"last", nil, "检查所有登录记录"},     // 默认last命令检查wtmp
		{"lastb", nil, "检查所有失败的登录尝试"}, // 默认lastb命令检查btmp
		{"lastlog", nil, "显示系统中所有用户的最后登录时间"},
		{"faillog", nil, "显示用户失败的登录尝试日志"},
		{"cat", []string{"/var/log/auth.log"}, "检查Debian系系统的认证日志"},
		{"cat", []string{"/var/log/secure"}, "检查Red Hat系系统的安全日志"},
		{"who", nil, "显示当前登录的用户信息"},
		{"w", nil, "显示当前登录用户及其正在执行的命令"},
	}

	for _, cmd := range commands {
		_, err := core.ExecuteCommand(cmd.Command, cmd.Args, cmd.Description)
		if err != nil {
			continue // 错误处理已在core.ExecuteCommand中完成，这里不再处理
		}
		// 输出结果已在core.ExecuteCommand中处理，此处无需进一步操作
	}
}
