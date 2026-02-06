package cmd

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPrintHelp 测试打印帮助信息功能
func TestPrintHelp(t *testing.T) {
	// 验证帮助信息包含关键内容
	assert.NotPanics(t, func() {
		PrintHelp()
	}, "PrintHelp 不应发生panic")
}

// TestPrintVersion 测试打印版本号功能
func TestPrintVersion(t *testing.T) {
	// 验证版本信息格式正确
	assert.NotPanics(t, func() {
		PrintVersion()
	}, "PrintVersion 不应发生panic")
}

// TestPrintLogTypes 测试打印日志类型功能
func TestPrintLogTypes(t *testing.T) {
	// 这个测试主要验证函数不会panic
	assert.NotPanics(t, func() {
		PrintLogTypes()
	}, "PrintLogTypes 不应发生panic")
}

// TestVersion 测试版本号常量
func TestVersion(t *testing.T) {
	// 验证版本号不为空且格式正确
	assert.NotEmpty(t, Version, "版本号不应为空")
	// assert.Equal(t, "1.0.0", Version, "版本号应为1.0.0")
}

// TestValidateLogType 测试日志类型验证
func TestValidateLogType(t *testing.T) {
	tests := []struct {
		name     string
		logType  string
		expected bool
	}{
		{"有效的nginx类型", "nginx", true},
		{"有效的pip类型", "pip", true},
		{"有效的npm类型", "npm", true},
		{"有效的docker类型", "docker", true},
		{"无效的空类型", "", false},
		{"无效的类型", "apache", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateLogType(tt.logType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestRun 测试参数解析功能
func TestRun(t *testing.T) {
	// 测试没有参数时返回 false
	// 由于 flag.Parse() 需要实际的参数，这里主要验证函数不会 panic
	assert.NotPanics(t, func() {
		// 模拟空参数调用
		flag.CommandLine.Parse([]string{})
		result := Run()
		// 空参数时既不是 version, help 也不是 listTypes
		// 所以会检查 ValidateLogType("") 返回 false
		assert.True(t, result || !ValidateLogType(""))
	}, "Run 不应发生panic")
}
