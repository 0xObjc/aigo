# AIGO - AI项目结构整理工具

AIGO（AI Project Structure Organizer）是一个用于将项目整理成文本格式的工具，旨在方便用户向AI提问和获取帮助。通过AIGO，你可以快速生成项目的目录结构和关键文件内容，并将其复制到剪贴板，以便在与AI交流时粘贴使用。

## 功能特点

- **项目结构生成**：自动生成项目的目录结构，包括文件夹和文件。
- **关键文件内容提取**：提取项目中的关键文件内容，如 `.go` 和 `go.mod` 文件。
- **支持所有文件**：通过 `-all` 参数，可以包含项目中的所有文件。
- **剪贴板复制**：生成的项目结构和文件内容可以直接复制到剪贴板，方便粘贴使用。

## 安装

1. 确保你已经安装了Go环境。
2. 下载并安装AIGO：

```sh
go get github.com/yourusername/aigo
```

3. 编译并安装AIGO：

```sh
cd $GOPATH/src/github.com/yourusername/aigo
go install
```

## 使用方法

### 基本用法

```sh
aigo <directory>
```

这将生成指定目录的项目结构和关键文件内容，并将其复制到剪贴板。

### 包含所有文件

```sh
aigo -all <directory>
```

这将生成指定目录的项目结构和所有文件内容，并将其复制到剪贴板。

## 示例

假设你有一个名为 `TeleNotify` 的项目目录，你可以使用以下命令生成项目结构：

```sh
aigo TeleNotify
```

生成的项目结构将类似于：

```markdown
### 项目结构

```
TeleNotify/
├── bot/
│   └── bot.go
├── config/
│   └── config.go
├── db/
│   └── db.go
├── handlers/
│   └── handlers.go
├── utils/
│   └── utils.go
├── main.go
├── go.mod
```

## 贡献

我们欢迎任何形式的贡献，包括但不限于代码改进、功能建议、文档完善等。请通过GitHub的Issue和Pull Request系统进行贡献。

## 许可证

本项目采用MIT许可证，详情请参见 [LICENSE](LICENSE) 文件。

---

感谢使用AIGO，希望它能帮助你更高效地与AI进行项目交流！