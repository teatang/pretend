// Package pip 提供pip安装日志生成功能
package pip

import (
	"fmt"
	"math/rand"

	"github.com/teatang/pretend/pkg/style"
)

// 日志组件定义
var (
	Packages = []string{"requests", "flask", "django", "numpy", "pandas", "pytest", "black", "flake8", "mypy", "httpx", "aiohttp", "celery", "redis", "sqlalchemy", "pydantic", "fastapi", "uvicorn"}
	Actions  = []string{"Collecting", "Downloading", "Installing", "Successfully installed", "Requirement already satisfied"}
)

// PrintColored 打印彩色pip日志
// 配对符号内使用相同颜色
func PrintColored() {
	colors := style.GetPipColors()
	action := Actions[rand.Intn(len(Actions))]

	if action == "Collecting" {
		pkg := Packages[rand.Intn(len(Packages))]
		version := rand.Intn(20) + 1
		colors.Action.Print("Collecting ")
		colors.Package.Print(pkg)
		colors.Version.Println("==" + fmt.Sprintf("%d.0", version))
		return
	}

	if action == "Downloading" {
		pkg := Packages[rand.Intn(len(Packages))]
		size := rand.Intn(500) + 10
		colors.Action.Print("  Downloading ")
		colors.Package.Print(pkg + "-")
		colors.Version.Print(fmt.Sprintf("%d.%d-py3-none-any.whl ", rand.Intn(10), rand.Intn(10)))
		colors.Size.Print(`(`)
		colors.Size.Print(fmt.Sprintf("%d", size))
		colors.Size.Println(` kB)`)
		return
	}

	if action == "Installing" {
		colors.Action.Print("Installing collected packages: ")
		colors.Package.Println(Packages[rand.Intn(len(Packages))])
		return
	}

	if action == "Successfully installed" {
		colors.Action.Print("Successfully installed ")
		colors.Package.Print(Packages[rand.Intn(len(Packages))] + "-")
		colors.Version.Println(fmt.Sprintf("%d.%d.0", rand.Intn(10), rand.Intn(10)))
		return
	}

	colors.Action.Println("Requirement already satisfied: pip in /usr/local/lib/python3.11/site-packages")
}

// GenerateText 生成纯文本pip日志
func GenerateText() string {
	action := Actions[rand.Intn(len(Actions))]

	if action == "Collecting" {
		pkg := Packages[rand.Intn(len(Packages))]
		version := rand.Intn(20) + 1
		return "Collecting " + pkg + "==" + fmt.Sprintf("%d.0", version)
	}

	if action == "Downloading" {
		pkg := Packages[rand.Intn(len(Packages))]
		size := rand.Intn(500) + 10
		return "  Downloading " + pkg + "-" + fmt.Sprintf("%d.%d-py3-none-any.whl (%d kB)", rand.Intn(10), rand.Intn(10), size)
	}

	if action == "Installing" {
		return "Installing collected packages: " + Packages[rand.Intn(len(Packages))]
	}

	if action == "Successfully installed" {
		return "Successfully installed " + Packages[rand.Intn(len(Packages))] + "-" + fmt.Sprintf("%d.%d.0", rand.Intn(10), rand.Intn(10))
	}

	return "Requirement already satisfied: pip in /usr/local/lib/python3.11/site-packages"
}
