package operation

import core "Jaeger/module"

// CheckSystemUsersAndGroups 检查系统中的用户和组
func CheckSystemUsersAndGroups() {
	commands := []struct {
		Command     string
		Args        []string
		Description string
	}{
		{"cat", []string{"/etc/passwd"}, "列出系统中所有用户"},
		{"cat", []string{"/etc/group"}, "列出系统中所有组"},
		{"cat", []string{"/etc/shadow"}, "查看用户密码及过期信息"},
		{"awk", []string{`-F: '{ print $1,$3,$6,$7 }' /etc/passwd`}, "列出所有用户及其ID、家目录和shell"},
		{"lslogins", nil, "列出所有用户的详细信息，包括最后登录时间和失败尝试次数"},
		{"chage", []string{"-l", "root"}, "显示root用户的密码过期信息"}, // 可以扩展到所有用户
		{"cat", []string{"/etc/sudoers"}, "列出sudoers文件内容，检查哪些用户具有sudo权限"},
		{"sudo", []string{"-l"}, "列出当前用户的sudo权限"},
		{"members", []string{"sudo"}, "列出所有属于sudo组的用户"},
		{"cut", []string{"-d:", "-f1", "/etc/passwd"}, "列出所有用户名"},
		{"grep", []string{"-v", "nologin", "/etc/passwd"}, "列出所有允许登录的用户"},
	}

	for _, cmd := range commands {
		_, err := core.ExecuteCommand(cmd.Command, cmd.Args, cmd.Description)
		if err != nil {
			continue // 错误处理已在core.ExecuteCommand中完成，这里不再处理
		}
	}
}
