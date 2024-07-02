package operation

import core "Jaeger/module"

// CheckSystemProcesses 使用多种命令检查系统的所有进程和打开的文件
func CheckSystemProcesses() {
	commands := []struct {
		Command     string
		Args        []string
		Description string
	}{
		{"ps", []string{"-eo", "user,time,pid,ppid,cmd"}, "检查所有系统进程及其详细信息"},
		{"top", []string{"-b", "-n", "1"}, "使用top命令检查所有系统进程"},
		{"lsof", []string{"-u", "root"}, "使用lsof命令检查root用户打开的文件"},
		{"lsof", []string{"-d", "0-9"}, "使用lsof命令检查文件描述符0-9的打开文件"},
		{"lsof", []string{"+D", "/var/log"}, "使用lsof命令检查/var/log目录下的打开文件"},
		{"lsof", []string{"+D", "/tmp"}, "使用lsof命令检查/tmp目录下的打开文件"},
		{"lsof", []string{"+D", "/var/tmp"}, "使用lsof命令检查/var/tmp目录下的打开文件"},
		{"lsof", []string{"+D", "/dev/shm"}, "使用lsof命令检查/dev/shm目录下的打开文件"},
		{"lsof", []string{"+D", "/usr/local/bin"}, "使用lsof命令检查/usr/local/bin目录下的打开文件"},
		{"lsof", []string{"+D", "/usr/local/sbin"}, "使用lsof命令检查/usr/local/sbin目录下的打开文件"},
	}

	for _, cmd := range commands {
		_, err := core.ExecuteCommand(cmd.Command, cmd.Args, cmd.Description)
		if err != nil {
			continue // 错误处理已在core.ExecuteCommand中完成，这里不再处理
		}
	}
}
