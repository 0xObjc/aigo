# Aigo - 为大模型设计的项目结构和代码枚举工具

Aigo 是一个专为大模型设计的命令行工具，旨在通过简单的指令快速枚举项目目录结构及代码，一键提供给 AI 进行理解和分析。Aigo 能够高效地生成项目结构，收集相关代码文件，并将其格式化为易于 AI 处理的数据，极大地提升了开发者和 AI 模型之间的协作效率。

## 特点

- **快速枚举**：通过简单的命令，快速生成项目目录结构和收集代码文件。
- **AI 友好**：生成的数据格式专为 AI 设计，便于模型理解和处理。
- **灵活配置**：支持选择性收集所有文件或仅特定文件（如 `.go` 文件和 `go.mod`）。
- **一键复制**：生成的项目结构和代码可以直接复制到剪贴板，方便快捷。

## 安装

要安装 Aigo，你需要在机器上安装 Go。然后，你可以克隆仓库并构建项目：

```sh
git clone https://github.com/0xObjc/aigo.git
cd aigo
go build -o aigo ./cmd/
```

## 使用

Aigo 可以通过命令行运行，语法如下：

```sh
aigo <directory> [-all]
```

- `<directory>`：你想要分析的目录。
- `[-all]`：一个可选标志，用于包含所有文件。如果省略，则只会收集 `.go` 文件和 `go.mod`。

### 示例

```sh
aigo ./myproject
```

这个命令将为 `./myproject` 生成项目结构，仅收集 `.go` 文件和 `go.mod`，渲染模板并复制到剪贴板。

```sh
aigo ./myproject -all
```

这个命令将为 `./myproject` 生成项目结构，收集所有文件，渲染模板并复制到剪贴板。

## 项目结构

Aigo 的项目结构如下：

```
github.com/0xObjc/aigo
├── cmd/
│   └── main.go
├── internal/
│   ├── collector/
│   │   └── collector.go
│   ├── config/
│   │   └── config.go
│   ├── generator/
│   │   └── generator.go
│   ├── model/
│   │   └── model.go
│   └── renderer/
│       └── renderer.go
└── go.mod
```

## 依赖

Aigo 使用以下外部依赖：

- `github.com/atotto/clipboard`：用于将渲染的模板复制到剪贴板。

## 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## 贡献

欢迎贡献！请随时提交拉取请求或打开问题以报告任何错误或功能请求。

## 作者

- [0xObjc](https://github.com/0xObjc)

## 致谢

- 感谢 Go 社区提供的优秀工具和库。