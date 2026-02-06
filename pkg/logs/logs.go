// Package logs 提供各种日志类型的生成和打印功能
// 用于模拟真实的工作日志输出
package logs

import (
	"github.com/teatang/pretend/pkg/logs/docker"
	"github.com/teatang/pretend/pkg/logs/nginx"
	"github.com/teatang/pretend/pkg/logs/npm"
	"github.com/teatang/pretend/pkg/logs/pip"
)

// LogType 日志类型定义
type LogType string

// 支持的日志类型常量
const (
	TypeNginx  LogType = "nginx"
	TypePip    LogType = "pip"
	TypeNpm    LogType = "npm"
	TypeDocker LogType = "docker"
)

// SupportedTypes 所有支持的日志类型列表
var SupportedTypes = []LogType{TypeNginx, TypePip, TypeNpm, TypeDocker}

// PrintColoredLog 根据日志类型打印彩色日志
func PrintColoredLog(logType LogType) {
	switch logType {
	case TypePip:
		pip.PrintColored()
	case TypeNpm:
		npm.PrintColored()
	case TypeDocker:
		docker.PrintColored()
	case TypeNginx:
		fallthrough
	default:
		nginx.PrintColored()
	}
}

// GenerateLog 根据日志类型生成对应的随机日志（纯文本）
func GenerateLog(logType LogType) string {
	switch logType {
	case TypePip:
		return pip.GenerateText()
	case TypeNpm:
		return npm.GenerateText()
	case TypeDocker:
		return docker.GenerateText()
	case TypeNginx:
		return nginx.GenerateText()
	default:
		return nginx.GenerateText()
	}
}

// IsValidType 验证日志类型是否有效
func IsValidType(logType LogType) bool {
	for _, t := range SupportedTypes {
		if t == logType {
			return true
		}
	}
	return false
}
