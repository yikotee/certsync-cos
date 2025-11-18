<h1 align="center">Welcome to CertSync-COS</h1>

<p align="center">
    <a href="https://goreportcard.com/report/github.com/yikotee/certsync-cos">
    <img src="https://goreportcard.com/badge/github.com/yikotee/certsync-cos" /></a>
    <a href="https://github.com/yikotee/certsync-cos/blob/main/README.md">
    <img src="https://img.shields.io/badge/Docs-使用文档-blue?style=flat-square&logo=readthedocs" alt="Docs" /></a>
</p>

<h3 align="center">一个轻量级的证书同步工具 🔒</h3>

<p align="center">
从腾讯云 COS（对象存储）自动下载 SSL/TLS 证书，并定期更新到本地服务器
</p>

## ✨ 功能特性

- ✅ **自动同步**：定期从腾讯云 COS 拉取最新的 SSL/TLS 证书
- ✅ **自动重载**：证书更新后自动执行自定义命令（如 `systemctl reload nginx`）
- ✅ **跨平台编译**：支持 Linux、macOS、Windows 多个平台和架构（x86_64、ARM64）
- ✅ **Docker 支持**：提供 Docker 和 Docker Compose 配置，便于容器化部署
- ✅ **灵活配置**：基于 YAML 配置文件，易于定制和部署

## 💡 项目背景

该项目用于解决以下场景：
- 将 SSL/TLS 证书存储在腾讯云 COS 中进行统一管理
- 在多个服务器上自动同步和更新证书
- 证书更新后自动触发服务重载，无需手动干预

## 🚀 快速开始

### 前置要求

- Go 1.23+
- Docker & Docker Compose（可选，仅用于容器部署）
- 腾讯云 COS 账户及 API 凭证

### 从源码编译

1. **克隆仓库**
   ```bash
   git clone https://github.com/yourusername/certsync-cos.git
   cd certsync-cos
   ```

2. **编译项目**

   ```bash
   # 编译当前平台
   go build -o certsync-cos
   
   # 或编译多个平台（使用 Makefile）
   make build-all
   ```

3. **配置应用**

   复制配置文件模板并编辑：

   ```bash
   cp config.example.yaml config.yaml
   # 编辑 config.yaml，填入腾讯云 COS 凭证
   ```

4. **运行应用**

   ```bash
   ./certsync-cos
   ```

### 使用 Docker Compose 部署

```bash
# 1. 编辑配置文件
cp config.example.yaml config.yaml

# 2. 修改 config.yaml 中的 COS 凭证和证书路径

# 3. 启动服务
docker-compose up -d

# 4. 查看日志
docker-compose logs -f certsync-cos

# 5. 停止服务
docker-compose down
```

## ⚙️ 配置说明

编辑 `config.yaml` 文件，配置以下参数：

```yaml
cos:
  bucket: "your-bucket-name"             # COS 存储桶名称
  region: "ap-beijing"                   # COS 地域（如 ap-beijing, ap-shanghai 等）
  secret_id: "your-secret-id"            # 腾讯云 API Secret ID
  secret_key: "your-secret-key"          # 腾讯云 API Secret Key
  prefix: "prefix"                       # COS 中的前缀路径
  cert_path: "your/cert/folder/xxx.cert" # COS 中证书文件的路径
  key_path: "your/cert/folder/xxx.key"   # COS 中私钥文件的路径

local:
  cert_path: "/etc/nginx/certs"          # 本地证书保存目录
  key_path: "/etc/nginx/certs"           # 本地私钥保存目录
  reload_cmd: "systemctl reload nginx"   # 证书更新后执行的命令（可选）

domain: "example.com"                    # 域名，用于生成证书文件名
schedule: "@every 12h"                   # 同步周期（使用 cron 表达式或简单时间间隔）
```

### 配置参数详解

| 参数 | 说明 | 示例 | 必填 |
|------|------|------|------|
| `bucket` | COS 存储桶名称 | `cert-1234567890` | ✅ |
| `region` | 腾讯云地域 | `ap-beijing` | ✅ |
| `secret_id` | 腾讯云 API Secret ID | - | ✅ |
| `secret_key` | 腾讯云 API Secret Key | - | ✅ |
| `prefix` | COS 中的前缀路径 | `certs` | ❌ |
| `cert_path` | COS 中证书文件的路径 | `xxx.crt` | ✅ |
| `key_path` | COS 中私钥文件的路径 | `xxx.key` | ✅ |
| `local.cert_path` | 本地证书保存目录 | `/etc/nginx/certs` | ✅ |
| `local.key_path` | 本地私钥保存目录 | `/etc/nginx/certs` | ✅ |
| `local.reload_cmd` | 证书更新后执行的命令 | `systemctl reload nginx` | ❌ |
| `domain` | 域名 | `example.com` | ✅ |
| `schedule` | 同步周期 | `@every 12h` | ✅ |

### 配置说明

- **COS 凭证**：从腾讯云控制台获取，用于连接和认证到 COS 服务
- **local.cert_path / local.key_path**：证书文件将保存为 `{cert_path}/{domain}.crt` 和 `{key_path}/{domain}.key`
- **reload_cmd**：证书更新后自动执行的命令，例如重新加载 Nginx 配置（可选）
- **schedule**：支持 cron 表达式或简单时间间隔（如 `@every 12h`、`@every 24h`）

## 📋 工作流程

```
┌──────────────────────────────────┐
│  启动应用（加载配置）            │
└──────────────┬───────────────────┘
               │
               ▼
┌──────────────────────────────────┐
│  创建 COS 客户端                  │
└──────────────┬───────────────────┘
               │
               ▼
┌──────────────────────────────────┐
│  从 COS 下载证书和私钥           │
└──────────────┬───────────────────┘
               │
               ▼
┌──────────────────────────────────┐
│  保存到本地指定目录              │
└──────────────┬───────────────────┘
               │
               ▼
┌──────────────────────────────────┐
│  执行重载命令（如有配置）        │
└──────────────┬───────────────────┘
               │
               ▼
┌──────────────────────────────────┐
│  等待下一个同步周期              │
└──────────────────────────────────┘
```

## 📁 项目结构

```
certsync-cos/
├── main.go              # 主程序入口
├── config.go            # 配置文件加载逻辑
├── cos.go               # COS 客户端封装
├── util.go              # 工具函数（命令执行等）
├── config.example.yaml  # 配置文件模板
├── Dockerfile           # Docker 镜像构建配置
├── docker-compose.yaml  # Docker Compose 配置
├── Makefile             # 编译脚本
├── go.mod               # Go 模块依赖
└── README.md            # 本文件
```

## 📚 依赖项

- `github.com/tencentyun/cos-go-sdk-v5`: 腾讯云 COS SDK
- `gopkg.in/yaml.v3`: YAML 配置文件解析

## 🔨 编译选项

### 使用 Makefile 编译多个平台

```bash
# 编译所有平台版本（Linux x86_64、Linux ARM64、macOS x86_64、macOS ARM64、Windows x86_64）
make build-all

# 清理编译产物
make clean
```

编译完成后，可执行文件将位于 `dist/` 目录下：
- `certsync-cos-linux-amd64` (Linux x86_64)
- `certsync-cos-linux-arm64` (Linux ARM64)
- `certsync-cos-darwin-amd64` (macOS Intel)
- `certsync-cos-darwin-arm64` (macOS Apple Silicon)
- `certsync-cos-windows-amd64.exe` (Windows x86_64)

## 🏢 常见使用场景

### 场景 1：Nginx 服务器自动更新证书

```yaml
# config.yaml
local:
  cert_path: "/etc/nginx/certs"
  key_path: "/etc/nginx/certs"
  reload_cmd: "systemctl reload nginx"  # 证书更新后自动重载 Nginx

domain: "example.com"
schedule: "@every 24h"  # 每天同步一次
```

### 场景 2：多域名证书管理

为不同的域名创建多个配置文件和应用实例，分别同步不同的证书：

```bash
# 实例 1：同步 example.com
certsync-cos -config config-example.yaml

# 实例 2：同步 api.example.com  
certsync-cos -config config-api.yaml
```

### 场景 3：Docker 容器部署

```bash
# 使用 Docker Compose 在容器中运行
docker-compose up -d

# 通过卷挂载实现证书实时同步到主机
# - ./certs:/app/certs
```

## 📝 日志输出

应用启动后会输出以下信息：

```
2025/11/17 10:15:30 证书同步完成: example.com
2025/11/17 10:15:30 Cert: /etc/nginx/certs/example.com.crt
2025/11/17 10:15:30 Key: /etc/nginx/certs/example.com.key
2025/11/17 10:15:30 执行reload: systemctl reload nginx
2025/11/17 10:15:31 reload成功
```

## 🔧 故障排查

### 连接 COS 失败

- ✅ 检查 `secret_id` 和 `secret_key` 是否正确
- ✅ 确认 COS bucket 和 region 配置无误
- ✅ 检查网络连接和防火墙设置

### 证书文件不存在

- ✅ 验证 `cert_path` 和 `key_path` 在 COS 中的路径是否正确
- ✅ 确保 COS 中的文件确实存在且名称相同

### 重载命令执行失败

- ✅ 检查 `reload_cmd` 是否正确（如 `systemctl reload nginx`）
- ✅ 确保应用运行的用户有执行该命令的权限（可能需要 sudo）
- ✅ 查看应用日志了解具体错误信息

## 🤝 贡献

欢迎贡献代码、报告问题或提出建议！

## 📌 变更日志

### v1.0.0 (2025-11-17)

- ✨ 初始版本发布
- ✨ 支持从腾讯云 COS 下载证书
- ✨ 支持定期更新和自动重载
- ✨ Docker 和多平台编译支持

---

如果这个项目对你有帮助，欢迎 ⭐ Star 支持！
