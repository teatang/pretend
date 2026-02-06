package logs

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestIsValidType 测试日志类型验证
func TestIsValidType(t *testing.T) {
	tests := []struct {
		name     string
		logType  LogType
		expected bool
	}{
		{"有效的nginx类型", TypeNginx, true},
		{"有效的pip类型", TypePip, true},
		{"有效的npm类型", TypeNpm, true},
		{"有效的docker类型", TypeDocker, true},
		{"无效的空类型", "", false},
		{"无效的类型", "apache", false},
		{"无效的类型", "mysql", false},
		{"大小写敏感", "Nginx", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidType(tt.logType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestGenerateNginxLog 测试Nginx日志生成
func TestGenerateNginxLog(t *testing.T) {
	// 生成多次日志确保随机性
	logs := make(map[string]int)
	for i := range 100 {
		_ = i
		log := GenerateLog(TypeNginx)
		logs[log]++
		// 验证日志格式
		assert.Contains(t, log, "HTTP/", "日志应包含HTTP协议")
		assert.Contains(t, log, " - - [", "日志应包含时间戳")
		// 验证有User-Agent（可以是curl, Python, Mozilla等）
		assert.Regexp(t, `(Mozilla|curl|Python|Go-http-client)`, log, "日志应有有效的User-Agent")
	}
	// 多次生成应该产生不同的日志
	assert.Greater(t, len(logs), 1, "多次生成应产生不同的日志")
}

// TestGeneratePipLog 测试pip日志生成
func TestGeneratePipLog(t *testing.T) {
	// 生成多次日志确保随机性
	logs := make(map[string]int)
	actions := make(map[string]bool)

	// pipActions包含: Collecting, Downloading, Installing, Successfully installed, Requirement already satisfied
	expectedActions := []string{"Collecting", "Downloading", "Installing", "Successfully installed", "Requirement already satisfied"}

	for i := range 100 {
		_ = i
		log := GenerateLog(TypePip)
		logs[log]++
		// 验证日志以有效的pip操作开头
		validStart := false
		for _, action := range expectedActions {
			if strings.HasPrefix(log, action) || strings.HasPrefix(log, "  "+action) {
				actions[action] = true
				validStart = true
				break
			}
		}
		assert.True(t, validStart, "日志应以有效的pip操作开头, got: %s", log)
	}

	// 验证所有pip操作类型都有可能出现
	assert.Greater(t, len(actions), 0, "应该生成各种类型的pip日志")
}

// TestGenerateNpmLog 测试npm日志生成
func TestGenerateNpmLog(t *testing.T) {
	// 生成多次日志确保随机性
	logs := make(map[string]int)

	for i := range 100 {
		_ = i
		log := GenerateLog(TypeNpm)
		logs[log]++
		// 验证日志不为空
		assert.NotEmpty(t, log, "npm日志不应为空")
	}

	// 多次生成应该产生不同的日志
	assert.Greater(t, len(logs), 1, "多次生成应产生不同的npm日志")
}

// TestGenerateDockerLog 测试Docker日志生成
func TestGenerateDockerLog(t *testing.T) {
	// 生成多次日志确保随机性
	logs := make(map[string]int)

	for i := range 100 {
		_ = i
		log := GenerateLog(TypeDocker)
		logs[log]++
		// 验证日志不为空
		assert.NotEmpty(t, log, "Docker日志不应为空")
	}

	// 多次生成应该产生不同的日志
	assert.Greater(t, len(logs), 1, "多次生成应产生不同的Docker日志")
}

// TestGenerateLog 测试通用日志生成函数
func TestGenerateLog(t *testing.T) {
	tests := []struct {
		name    string
		logType LogType
	}{
		{"nginx类型", TypeNginx},
		{"pip类型", TypePip},
		{"npm类型", TypeNpm},
		{"docker类型", TypeDocker},
		{"未知类型默认nginx", "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateLog(tt.logType)
			require.NotEmpty(t, result, "生成的日志不应为空")
		})
	}
}

// TestSupportedTypes 测试支持的类型列表
func TestSupportedTypes(t *testing.T) {
	require.Len(t, SupportedTypes, 4, "应该支持4种日志类型")

	typesMap := make(map[LogType]bool)
	for _, t := range SupportedTypes {
		typesMap[t] = true
	}

	assert.True(t, typesMap[TypeNginx], "应包含nginx")
	assert.True(t, typesMap[TypePip], "应包含pip")
	assert.True(t, typesMap[TypeNpm], "应包含npm")
	assert.True(t, typesMap[TypeDocker], "应包含docker")
}

// TestPrintColoredLog 测试彩色日志打印（确保不panic）
func TestPrintColoredLog(t *testing.T) {
	types := []LogType{TypeNginx, TypePip, TypeNpm, TypeDocker}

	for _, lt := range types {
		t.Run(string(lt), func(t *testing.T) {
			assert.NotPanics(t, func() {
				PrintColoredLog(lt)
			})
		})
	}
}
