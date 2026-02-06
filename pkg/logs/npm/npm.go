// Package npm 提供npm安装日志生成功能
package npm

import (
	"fmt"
	"math/rand"

	"github.com/teatang/pretend/pkg/style"
)

// 日志组件定义
var (
	Packages = []string{"express", "react", "vue", "lodash", "axios", "webpack", "typescript", "eslint", "prettier", "jest", "mocha"}
	Scripts  = []string{"build", "start", "test", "dev", "lint"}
)

// PrintColored 打印彩色npm日志
// 配对符号内使用相同颜色
func PrintColored() {
	colors := style.GetNpmColors()
	action := rand.Intn(10)
	pkg := Packages[rand.Intn(len(Packages))]

	switch {
	case action < 4:
		colors.Number.Printf("added %d", rand.Intn(100))
		colors.Number.Println(" packages")
	case action < 6:
		colors.Number.Print("found ")
		colors.Number.Print("0")
		colors.Number.Println(" vulnerabilities")
	case action < 8:
		script := Scripts[rand.Intn(len(Scripts))]
		colors.Action.Print("> ")
		colors.Package.Print(pkg)
		colors.Number.Print("@1.0.0 ")
		colors.Script.Println(script)
	case action < 9:
		colors.Warning.Print("npm WARN deprecated ")
		colors.Package.Println(pkg + "@1.0.0")
	default:
		colors.Action.Print("npm notice ")
		colors.Action.Print("New major version of npm available! ")
		colors.Number.Printf("%d.%d", rand.Intn(10), rand.Intn(10))
		colors.Arrow.Print(` -> `)
		colors.Number.Printf("%d.%d", rand.Intn(10), rand.Intn(10))
		colors.Number.Println("")
	}
}

// GenerateText 生成纯文本npm日志
func GenerateText() string {
	action := rand.Intn(10)
	pkg := Packages[rand.Intn(len(Packages))]

	switch {
	case action < 4:
		return fmt.Sprintf("added %d packages", rand.Intn(100))
	case action < 6:
		return "found 0 vulnerabilities"
	case action < 8:
		script := Scripts[rand.Intn(len(Scripts))]
		return "> " + pkg + "@1.0.0 " + script
	case action < 9:
		return "npm WARN deprecated " + pkg + "@1.0.0"
	default:
		return fmt.Sprintf("npm notice New major version of npm available! %d.%d -> %d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10))
	}
}
