// Package cmd 提供命令行参数解析和打印功能
package cmd

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"time"

	"github.com/teatang/pretend/pkg/logs"
)

// Version 版本号
const Version = "0.6.24"

// 命令行参数变量
var (
	versionFlag    bool // -v 显示版本号
	helpFlag       bool // -h 显示帮助信息
	listTypesFlag  bool // -l 显示可以打印的日志类型
	LogType        string // -t 选择需要打印的日志类型
)

func init() {
	flag.BoolVar(&versionFlag, "v", false, "显示版本号")
	flag.BoolVar(&helpFlag, "h", false, "显示帮助信息")
	flag.BoolVar(&listTypesFlag, "l", false, "显示可以打印的日志类型")
	flag.StringVar(&LogType, "t", string(logs.TypeNginx), "选择需要打印的日志类型 (nginx|pip|npm|docker)")
}

// PrintVersion 打印版本号
func PrintVersion() {
	fmt.Printf("pretend version: %s\n", Version)
}

// PrintHelp 打印帮助信息
func PrintHelp() {
	fmt.Println(`用法: pretend [选项]

选项:
  -v          显示版本号
  -h          显示帮助信息
  -l          显示可以打印的日志类型
  -t <类型>   选择需要打印的日志类型 (默认: nginx)

可用日志类型:
  nginx   - Nginx访问日志
  pip     - Python pip安装日志
  npm     - Node.js npm安装日志
  docker  - Docker镜像拉取日志

示例:
  pretend                    # 打印nginx日志
  pretend -t pip            # 打印pip安装日志
  pretend -t npm -l         # 显示npm类型并列出日志类型`)
}

// PrintLogTypes 显示支持的日志类型
func PrintLogTypes() {
	fmt.Println("支持的日志类型:")
	for _, t := range logs.SupportedTypes {
		fmt.Printf("  - %s\n", t)
	}
}

// ValidateLogType 验证日志类型是否支持
func ValidateLogType(logType string) bool {
	return slices.Contains(logs.SupportedTypes, logs.LogType(logType))
}

// Run 解析参数并执行相应操作
// 返回 true 表示已处理完（如 -v, -h, -l），false 表示需要继续运行
func Run() bool {
	flag.Parse()

	if versionFlag {
		PrintVersion()
		return true
	}

	if helpFlag {
		PrintHelp()
		return true
	}

	if listTypesFlag {
		PrintLogTypes()
		return true
	}

	if !ValidateLogType(LogType) {
		fmt.Printf("错误: 不支持的日志类型 '%s'\n", LogType)
		fmt.Println("使用 -l 参数查看支持的日志类型")
		os.Exit(1)
	}

	return false
}

// StartPrintingLogs 开始持续打印日志
func StartPrintingLogs(logType string) {
	fmt.Printf("\n开始打印 %s 日志 (按 Ctrl+C 停止):\n\n", logType)

	for {
		// 打印带颜色的日志（不带序号）
		logs.PrintColoredLog(logs.LogType(logType))
		// 随机延迟，让日志看起来更真实
		time.Sleep(time.Duration(200+rand.Intn(800)) * time.Millisecond)
	}
}
