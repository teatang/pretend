// Package style 提供终端输出的颜色样式定义
package style

import "github.com/fatih/color"

// Nginx日志颜色配置
type NginxColors struct {
	IP      *color.Color // IP地址
	Time    *color.Color // 时间戳 [xxx]
	Path    *color.Color // 请求路径
	Method  *color.Color // 请求方法
	Status  *color.Color // 状态码
	Bytes   *color.Color // 字节数
	Agent   *color.Color // User-Agent "xxx"
}

// Pip日志颜色配置
type PipColors struct {
	Action   *color.Color // 操作类型
	Package  *color.Color // 包名
	Version  *color.Color // 版本号
	Size     *color.Color // 大小
	Path     *color.Color // 路径
}

// Npm日志颜色配置
type NpmColors struct {
	Action   *color.Color // npm标识
	Number   *color.Color // 数字
	Package  *color.Color // 包名
	Script   *color.Color // 脚本名
	Warning  *color.Color // 警告
	Arrow    *color.Color // 箭头 ->
}

// Docker日志颜色配置
type DockerColors struct {
	Action  *color.Color // 操作类型
	Image   *color.Color // 镜像名
	Tag     *color.Color // 标签
	Status  *color.Color // Status行
}

// Nginx 日志颜色 - 配对符号内使用相同颜色
var Nginx = NginxColors{
	IP:     color.New(color.FgYellow),
	Time:   color.New(color.FgMagenta),  // [时间]
	Path:   color.New(color.FgCyan),      // /api/xxx
	Method: color.New(color.FgGreen),     // GET/POST
	Status: color.New(color.FgRed),       // 状态码
	Bytes:  color.New(color.FgBlue),      // 字节数
	Agent:  color.New(color.FgHiWhite),  // "User-Agent"
}

// Pip 日志颜色
var Pip = PipColors{
	Action:  color.New(color.FgCyan),     // Collecting/Downloading
	Package: color.New(color.FgGreen),     // 包名
	Version: color.New(color.FgYellow),   // 版本号
	Size:    color.New(color.FgBlue),     // (xxx kB)
	Path:    color.New(color.FgWhite),    // 路径
}

// Npm 日志颜色
var Npm = NpmColors{
	Action:  color.New(color.FgRed),        // npm
	Number:  color.New(color.FgYellow),    // 数字
	Package: color.New(color.FgGreen),     // 包名
	Script:  color.New(color.FgCyan),      // build/start等
	Warning: color.New(color.FgMagenta),  // WARN
	Arrow:   color.New(color.FgBlue),      // ->
}

// Docker 日志颜色
var Docker = DockerColors{
	Action: color.New(color.FgBlue),    // Pulling/Status
	Image:  color.New(color.FgGreen),   // nginx/node等
	Tag:    color.New(color.FgYellow),  // latest/18等
	Status: color.New(color.FgBlue),    // Status: Downloaded
}

// GetNginxColors 获取Nginx日志颜色配置
func GetNginxColors() NginxColors {
	return Nginx
}

// GetPipColors 获取Pip日志颜色配置
func GetPipColors() PipColors {
	return Pip
}

// GetNpmColors 获取Npm日志颜色配置
func GetNpmColors() NpmColors {
	return Npm
}

// GetDockerColors 获取Docker日志颜色配置
func GetDockerColors() DockerColors {
	return Docker
}
