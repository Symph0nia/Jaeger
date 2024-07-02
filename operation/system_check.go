package operation

import (
	"Jaeger/module" // 根据实际路径导入core包
)

// CheckSystemVersion 使用多种方式检查Linux系统版本，不负责打印输出
func CheckSystemVersion() {
	commands := []struct {
		Command     string
		Args        []string
		Description string // 添加描述以理解每个命令的目的
	}{
		{"uname", []string{"-a"}, "获取所有系统信息"},
		{"lsb_release", []string{"-a"}, "获取Linux标准基础版本信息"},
		{"cat", []string{"/etc/os-release"}, "获取系统发行版信息"},
		{"cat", []string{"/proc/version"}, "获取内核版本信息"},
		{"hostnamectl", nil, "使用hostnamectl获取系统信息，主要针对systemd系统"},
		{"dmesg", []string{"| grep 'Linux'"}, "获取内核相关的启动信息"},
	}

	for _, cmd := range commands {
		_, err := core.ExecuteCommand(cmd.Command, cmd.Args, cmd.Description)
		if err != nil {
			continue // 错误处理已在core.ExecuteCommand中完成，这里不再处理
		}
	}
}
