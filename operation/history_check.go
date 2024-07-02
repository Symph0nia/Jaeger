package operation

import (
	core "Jaeger/module"
	"fmt"
	"os"
	"path/filepath"
)

// CheckCommandHistory 检查所有用户的命令历史
func CheckCommandHistory() {
	// 定义常见的shell历史文件名
	historyFiles := []string{
		".bash_history",
		".zsh_history",
		".fish_history",
	}

	// 获取系统上的所有用户
	users, err := os.ReadDir("/home")
	if err != nil {
		fmt.Println("无法读取用户目录：", err)
		return
	}

	for _, userDir := range users {
		if !userDir.IsDir() {
			continue // 只处理目录
		}

		username := userDir.Name()
		userHome := filepath.Join("/home", username)

		for _, historyFile := range historyFiles {
			fullPath := filepath.Join(userHome, historyFile)
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				continue // 跳过不存在历史文件的用户
			}

			cmd := "cat"
			args := []string{fullPath}
			description := "检查用户 " + username + " 的命令历史文件：" + historyFile
			_, err = core.ExecuteCommand(cmd, args, description)
			if err != nil {
				continue // 错误处理已在core.ExecuteCommand中完成，这里不再处理
			}
		}
	}

	// 处理root用户的历史文件
	rootHistoryFiles := []string{
		"/root/.bash_history",
		"/root/.zsh_history",
		"/root/.fish_history",
	}

	for _, rootHistoryFile := range rootHistoryFiles {
		if _, err := os.Stat(rootHistoryFile); err == nil {
			cmd := "cat"
			args := []string{rootHistoryFile}
			description := "检查root用户的命令历史文件：" + rootHistoryFile
			_, err = core.ExecuteCommand(cmd, args, description)
			if err != nil {
				// 错误处理已在core.ExecuteCommand中完成，这里不再处理
			}
		}
	}
}
