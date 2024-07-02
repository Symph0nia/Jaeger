package operation

import core "Jaeger/module"

// CheckRecentFiles 列出系统中最近更改的500个文件
func CheckRecentFiles() {
	command := "find / -type f -printf '%T@ %p\n' | sort -n | tail -500"
	description := "列出系统中最近更改的500个文件"
	_, err := core.ExecuteCommand("bash", []string{"-c", command}, description)
	if err != nil {
		// 错误处理已在core.ExecuteCommand中完成，这里不再处理
		return
	}
}
