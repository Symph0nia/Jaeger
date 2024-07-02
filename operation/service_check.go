package operation

import core "Jaeger/module"

// CheckSystemServices 检查系统中正在运行的服务和开机自启动的服务
func CheckSystemServices() {
	commands := []struct {
		Command     string
		Args        []string
		Description string
	}{
		{"systemctl", []string{"list-units", "--type=service", "--state=running"}, "列出系统中正在运行的服务"},
		{"systemctl", []string{"list-unit-files", "--type=service"}, "列出系统中所有服务及其启动状态"},
		{"chkconfig", []string{"--list"}, "列出所有开机自启动服务（适用于SysVinit系统）"},
		{"service", []string{"--status-all"}, "列出所有服务及其状态（适用于SysVinit系统）"},
		{"initctl", []string{"list"}, "列出所有服务及其状态（适用于Upstart系统）"},
		{"systemctl", []string{"list-timers", "--all"}, "列出所有定时器（适用于systemd）"},
		{"systemctl", []string{"--failed"}, "列出启动过程中失败的服务"},
	}

	for _, cmd := range commands {
		_, err := core.ExecuteCommand(cmd.Command, cmd.Args, cmd.Description)
		if err != nil {
			continue // 错误处理已在core.ExecuteCommand中完成，这里不再处理
		}
	}
}
