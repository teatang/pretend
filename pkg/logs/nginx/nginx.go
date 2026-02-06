// Package nginx 提供Nginx日志生成功能
package nginx

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/teatang/pretend/pkg/style"
)

// 日志组件定义
var (
	IPs       = []string{"192.168.1.100", "192.168.1.101", "10.0.0.55", "172.16.0.23", "192.168.1.102", "10.10.10.15", "192.168.50.100", "8.8.8.8"}
	Methods   = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"}
	Paths     = []string{"/api/users", "/api/login", "/api/orders", "/static/css/main.css", "/images/logo.png", "/api/products", "/api/checkout", "/health", "/api/v1/users", "/api/v2/orders", "/dashboard", "/api/search", "/api/notifications", "/api/settings", "/favicon.ico", "/robots.txt"}
	Protocols = []string{"HTTP/1.1", "HTTP/1.0", "HTTP/2.0"}
	Statuses  = []string{"200", "201", "204", "301", "302", "304", "400", "401", "403", "404", "500", "502", "503"}
	Referers  = []string{"-", "https://google.com", "https://github.com", "https://baidu.com", "https://bing.com"}
	UserAgents = []string{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X)",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"curl/7.88.1", "Python-urllib/3.11", "Go-http-client/1.1",
	}
)

// PrintColored 打印彩色Nginx日志
// 格式: IP - - [时间] "方法 路径 协议" 状态码 字节数 Referer "User-Agent"
// 配对符号内的内容使用相同颜色
func PrintColored() {
	colors := style.GetNginxColors()
	ip := IPs[rand.Intn(len(IPs))]
	method := Methods[rand.Intn(len(Methods))]
	path := Paths[rand.Intn(len(Paths))]
	protocol := Protocols[rand.Intn(len(Protocols))]
	status := Statuses[rand.Intn(len(Statuses))]
	bodyBytes := rand.Intn(10000)
	referer := Referers[rand.Intn(len(Referers))]
	userAgent := UserAgents[rand.Intn(len(UserAgents))]
	now := time.Now().Format("02/Jan/2006:15:04:05 +0800")

	// 打印日志，配对符号内使用相同颜色
	colors.IP.Print(ip)
	colors.IP.Print(` - - `)
	colors.Time.Print(`[`)
	colors.Time.Print(now)
	colors.Time.Print(`] `)
	colors.Method.Print(`"`)
	colors.Method.Print(method + ` `)
	colors.Path.Print(path + ` `)
	colors.Method.Print(protocol + `" `)
	colors.Status.Print(status + ` `)
	colors.Bytes.Printf("%04d", bodyBytes)
	colors.IP.Print(` ` + referer + ` "`)
	colors.Agent.Print(userAgent)
	colors.Agent.Println(`"`)
}

// GenerateText 生成纯文本Nginx日志
func GenerateText() string {
	ip := IPs[rand.Intn(len(IPs))]
	method := Methods[rand.Intn(len(Methods))]
	path := Paths[rand.Intn(len(Paths))]
	protocol := Protocols[rand.Intn(len(Protocols))]
	status := Statuses[rand.Intn(len(Statuses))]
	bodyBytes := rand.Intn(10000)
	referer := Referers[rand.Intn(len(Referers))]
	userAgent := UserAgents[rand.Intn(len(UserAgents))]
	now := time.Now().Format("02/Jan/2006:15:04:05 +0800")

	return ip + ` - - [` + now + `] "` + method + ` ` + path + ` ` + protocol + `" ` + status + ` ` +
		fmt.Sprintf("%04d", bodyBytes) + ` ` + referer + ` "` + userAgent + `"`
}
