// Package docker 提供Docker日志生成功能
package docker

import (
	"math/rand"

	"github.com/teatang/pretend/pkg/style"
)

// 日志组件定义
var (
	Images   = []string{"nginx", "node", "python", "redis", "mysql", "postgres", "mongo", "rabbitmq", "elasticsearch", "ubuntu", "alpine"}
	Actions  = []string{"Pulling", "Pull complete", "Extracting", "Download complete", "Status"}
	Tags     = []string{"latest", "18", "20", "22", "1.0", "latest", "stable", "alpine"}
)

// PrintColored 打印彩色Docker日志
// 配对符号内使用相同颜色
func PrintColored() {
	colors := style.GetDockerColors()
	action := Actions[rand.Intn(len(Actions))]
	image := Images[rand.Intn(len(Images))]
	tag := Tags[rand.Intn(len(Tags))]

	switch action {
	case "Pulling":
		layer := rand.Intn(5)
		if layer == 0 {
			colors.Action.Print("Pulling from library/")
			colors.Image.Println(image)
		} else {
			colors.Action.Println("Pulling fs layer")
		}
	case "Pull complete":
		colors.Action.Println("Pull complete")
	case "Extracting":
		colors.Action.Println("Extracting")
	case "Download complete":
		colors.Action.Println("Download complete")
	case "Status":
		colors.Status.Print("Status: Downloaded newer image for ")
		colors.Image.Print(image + ":")
		colors.Tag.Println(tag)
	default:
		colors.Image.Print("docker.io/library/" + image + ":")
		colors.Tag.Println(tag)
	}
}

// GenerateText 生成纯文本Docker日志
func GenerateText() string {
	action := Actions[rand.Intn(len(Actions))]
	image := Images[rand.Intn(len(Images))]
	tag := Tags[rand.Intn(len(Tags))]

	switch action {
	case "Pulling":
		layer := rand.Intn(5)
		if layer == 0 {
			return "Pulling from library/" + image
		}
		return "Pulling fs layer"
	case "Pull complete":
		return "Pull complete"
	case "Extracting":
		return "Extracting"
	case "Download complete":
		return "Download complete"
	case "Status":
		return "Status: Downloaded newer image for " + image + ":" + tag
	default:
		return "docker.io/library/" + image + ":" + tag
	}
}
