# CLAUDE.md

本项目是一个用Go语言开发的命令行工具，用于在终端打印各种类型的日志。

## 项目概述

- **项目名称**: pretend
- **用途**: 在终端持续打印日志，模拟认真工作的假象
- **主要功能**: 支持Nginx、pip、npm、Docker四种日志类型的随机彩色输出

## 开发环境

- Go 1.25+
- macOS/Darwin

## 项目结构

```
pretend/
├── main.go                 # 主程序入口（最简单的入口文件）
├── go.mod                  # Go模块文件
├── .gitignore              # Git忽略文件
├── README.md               # 项目说明文档
├── CLAUDE.md               # Claude AI辅助开发文档
├── cmd/
│   ├── cmd.go              # 命令行参数解析和打印功能
│   └── cmd_test.go         # 命令行功能单元测试
├── pkg/
│   ├── logs/
│   │   ├── logs.go         # 日志生成核心逻辑
│   │   └── logs_test.go    # 日志生成单元测试
│   └── style/
│       └── style.go         # 颜色样式定义
```

## 依赖管理

使用Go Modules管理依赖：

```bash
# 添加新依赖
go get <package>

# 更新依赖
go get -u <package>

# 下载所有依赖
go mod download
```

当前依赖：
- `github.com/fatih/color` - 终端彩色输出
- `github.com/stretchr/testify` - 单元测试框架

## 代码规范

### 文件命名
- 主程序文件: `main.go` - 只有几行代码，仅作为入口
- 命令行包: `cmd/cmd.go` - 处理命令行参数、帮助信息等
- 命令行测试: `cmd/cmd_test.go`
- 日志包目录: `pkg/logs/`
- 日志测试文件: `pkg/logs/logs_test.go`
- 样式包: `pkg/style/style.go` - 颜色定义
- 遵循Go标准命名规范

### 注释要求
- 代码中需要加入适量的**中文注释**
- 公共函数和常量需要添加注释说明其用途
- 复杂逻辑需要添加行内注释

### 入口文件 (main.go)

```go
package main

import "github.com/teatang/pretend/cmd"

func main() {
    done := cmd.Run()
    if done {
        return
    }
    cmd.StartPrintingLogs(cmd.LogType)
}
```

### 样式包结构 (pkg/style/style.go)

```go
// Package style 提供终端输出的颜色样式定义
package style

// 颜色配置类型
type NginxColors struct {
    IP      *color.Color // IP地址
    Time    *color.Color // 时间戳
    Method  *color.Color // 请求方法
    Path    *color.Color // 请求路径
    Status  *color.Color // 状态码
    Bytes   *color.Color // 字节数
    Referer *color.Color // Referer
    Agent   *color.Color // User-Agent
}

// 预定义的颜色配置
var Nginx = NginxColors{...}
var Pip = PipColors{...}
var Npm = NpmColors{...}
var Docker = DockerColors{...}

// 获取颜色的函数
func GetNginxColors() NginxColors
func GetPipColors() PipColors
func GetNpmColors() NpmColors
func GetDockerColors() DockerColors
```

### 日志包结构 (pkg/logs/logs.go)

```go
// Package logs 提供各种日志类型的生成和打印功能
package logs

type LogType string

const (
    TypeNginx  LogType = "nginx"
    TypePip    LogType = "pip"
    TypeNpm    LogType = "npm"
    TypeDocker LogType = "docker"
)

var SupportedTypes = []LogType{...}

// 彩色打印函数
func PrintColoredNginxLog()
func PrintColoredPipLog()
func PrintColoredNpmLog()
func PrintColoredDockerLog()
func PrintColoredLog(logType LogType)

// 纯文本生成函数
func GenerateLog(logType LogType) string

func IsValidType(logType LogType) bool
```

## 命令行参数

使用 `flag` 包定义参数：

| 参数 | 变量 | 说明 |
|------|------|------|
| `-v` | `versionFlag` | 显示版本号 |
| `-h` | `helpFlag` | 显示帮助信息 |
| `-l` | `listTypesFlag` | 显示可以打印的日志类型 |
| `-t` | `LogType` | 选择需要打印的日志类型，默认nginx |

## 测试要求

### 测试框架
- 使用 `github.com/stretchr/testify` 框架
- 主要使用 `assert` 和 `require` 包

### 测试文件结构
- 测试文件命名: `<package>_test.go`
- 测试函数命名: `TestFunctionName`
- 每个功能点需要对应测试用例

### 运行测试

```bash
# 运行所有测试
go test -v ./...

# 运行特定测试
go test -v -run TestFunctionName ./...

# 每次代码修改后必须执行所有单元测试
go test ./...
```

### 命令行包测试 (cmd/cmd_test.go)

必须包含以下测试：
- `TestPrintHelp` - 测试打印帮助信息
- `TestPrintVersion` - 测试打印版本号
- `TestPrintLogTypes` - 测试打印日志类型列表
- `TestVersion` - 测试版本号常量
- `TestValidateLogType` - 测试日志类型验证
- `TestRun` - 测试参数解析功能

### 日志包测试 (pkg/logs/logs_test.go)

必须包含以下测试：
- `TestIsValidType` - 测试日志类型验证功能
- `TestGenerateNginxLog` - 测试Nginx日志生成
- `TestGeneratePipLog` - 测试pip日志生成
- `TestGenerateNpmLog` - 测试npm日志生成
- `TestGenerateDockerLog` - 测试Docker日志生成
- `TestGenerateLog` - 测试通用日志生成函数
- `TestSupportedTypes` - 测试支持的日志类型列表
- `TestPrintColoredLog` - 测试彩色日志打印

## 开发流程

1. **添加新功能后**：
   - 命令行相关功能添加到 `cmd/cmd.go`
   - 日志生成逻辑添加到 `pkg/logs/logs.go`
   - 颜色配置添加到 `pkg/style/style.go`
   - 执行 `go test -v ./...` 验证所有测试通过
   - 更新 `README.md` 文档

2. **main.go 是入口文件**，不应包含复杂逻辑

3. **添加新的日志类型**：
   1. 在 `pkg/logs/logs.go` 中定义新的日志组件变量
   2. 添加对应的 `PrintColoredXxxLog()` 和 `generateXxxLogText()` 函数
   3. 在 `pkg/style/style.go` 中添加对应的颜色配置
   4. 在 `SupportedTypes` 切片中添加类型名称
   5. 在 `PrintColoredLog()` 函数中添加 case 分支
   6. 在 `GenerateLog()` 函数中添加 case 分支
   7. 在 `README.md` 中更新日志类型说明
   8. 在 `pkg/logs/logs_test.go` 中添加对应的测试用例

## 常用命令

```bash
# 构建项目
go build -o pretend

# 运行程序
./pretend -t nginx

# 运行测试
go test -v ./...

# 查看测试覆盖率
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

## 日志特点

- **随机性**: 日志内容每次生成都不同，更加真实
- **多样性**: 包含多种状态码、请求方法、URL路径
- **时序性**: 时间戳随当前时间变化
- **随机延迟**: 每次打印后有200-1000ms随机延迟，模拟真实场景
- **多彩输出**: 不同部分使用不同颜色，视觉效果更好

## 注意事项

- 每次修改代码后必须执行 `go test ./...` 确保所有测试通过
- 新增日志类型需要确保随机性和多样性
- main.go 只作为入口文件，核心逻辑应在 cmd 包、pkg/logs 包或 pkg/style 包中实现
- 保持代码简洁，复杂的逻辑应该放到专门的包中
