// Package main pretend - 假装在工作
// 一个在终端持续打印各种日志的命令行工具，让别人以为你在认真工作
package main

import "github.com/teatang/pretend/cmd"

func main() {
	// 解析命令行参数
	done := cmd.Run()
	if done {
		return
	}
	// 开始打印日志
	cmd.StartPrintingLogs(cmd.LogType)
}
