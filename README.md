# pretend - 假装在工作

一个有趣的命令行工具，可以在终端持续打印各种类型的日志，让别人以为你在认真工作。

## 功能特点

- 支持多种日志类型：Nginx、pip、npm、Docker
- 彩色输出，视觉效果更好
- 随机生成日志，每次都不同，更加真实
- 简单易用的命令行参数
- 基于标准库 `flag` 开发
- 完整的单元测试覆盖

## 安装

```bash
# 克隆项目
git clone https://github.com/yourusername/pretend.git
cd pretend

# 构建
go build -o pretend

# 安装到PATH（可选）
sudo mv pretend /usr/local/bin/
```

## 使用方法

```bash
# 显示帮助信息
./pretend -h

# 显示版本号
./pretend -v

# 显示支持的日志类型
./pretend -l

# 打印Nginx日志（默认）
./pretend

# 打印pip安装日志
./pretend -t pip

# 打印npm安装日志
./pretend -t npm

# 打印Docker拉取日志
./pretend -t docker
```

## 命令行参数

| 参数 | 说明 |
|------|------|
| `-v` | 显示版本号 |
| `-h` | 显示帮助信息 |
| `-l` | 显示可以打印的日志类型 |
| `-t <类型>` | 选择需要打印的日志类型（默认: nginx） |

## 支持的日志类型

| 类型 | 说明 | 颜色主题 |
|------|------|----------|
| `nginx` | Nginx访问日志 | IP黄、时间洋红、方法绿、路径青、状态红 |
| `pip` | Python pip安装日志 | 操作青、包绿、版本黄、大小蓝 |
| `npm` | Node.js npm安装日志 | npm红、包绿、数字黄、脚本青 |
| `docker` | Docker镜像拉取日志 | 操作蓝、镜像绿、标签黄 |

## 运行测试

```bash
go test -v ./...
```

## 项目结构

```
pretend/
├── main.go                           # 主程序入口
├── go.mod                           # Go模块文件
├── .gitignore                       # Git忽略文件
├── LICENSE                          # 许可证文件
├── README.md                        # 项目说明文档
├── CLAUDE.md                        # Claude AI辅助开发文档
├── cmd/
│   ├── cmd.go                      # 命令行参数解析和打印功能
│   └── cmd_test.go                 # 命令行功能单元测试
└── pkg/
    ├── style/
    │   └── style.go               # 颜色样式定义
    └── logs/
        ├── logs.go                # 日志生成核心逻辑（分发器）
        ├── logs_test.go           # 日志生成单元测试
        ├── nginx/
        │   └── nginx.go          # Nginx日志生成
        ├── pip/
        │   └── pip.go            # Pip日志生成
        ├── npm/
        │   └── npm.go             # Npm日志生成
        └── docker/
            └── docker.go          # Docker日志生成
```

## 技术栈

- Go 1.25+
- [flag](https://pkg.go.dev/flag) - 命令行参数解析
- [fatih/color](https://github.com/fatih/color) - 终端彩色输出
- [stretchr/testify](https://github.com/stretchr/testify) - 单元测试框架

## 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件
