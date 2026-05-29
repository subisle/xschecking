# 授权系统后端基础框架

基于 Go + Gin 的授权系统后端骨架，面向注册、登录、验证码、邮箱验证码、JWT 会话和基础风控流程。

当前仓库更偏向后端框架和开发基线，不是完整业务成品。

## 技术栈

| 模块 | 技术 |
| --- | --- |
| Web 框架 | Gin |
| 配置 | Viper |
| ORM | GORM |
| 数据库 | MySQL 8.x |
| 缓存 | Redis |
| 认证 | JWT |
| 日志 | Zap |
| 密码哈希 | bcrypt |

## 快速开始

### 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis 6.0+

### 安装依赖

```bash
go mod download
```

### 配置

复制或修改 `config.yaml`，填入本地开发环境配置：

```yaml
database:
  host: localhost
  port: 3306
  user: root
  password: your_password
  dbname: auth_system

redis:
  host: localhost
  port: 6379
  password: ""

jwt:
  secret: your_jwt_secret_key_change_in_production
```

生产环境不要使用示例密码、空 Redis 密码或弱 JWT secret。

### 创建数据库

```sql
CREATE DATABASE auth_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 启动服务

```bash
go run cmd/server/main.go
```

服务默认运行在：

```text
http://localhost:8080
```

## API 概览

### 公开接口

- `GET /api/v1/health`：健康检查。
- `POST /api/v1/captcha/generate`：生成验证码。
- `POST /api/v1/captcha/verify`：验证验证码并生成一次性通行证。
- `POST /api/v1/auth/register`：用户注册。
- `POST /api/v1/auth/login`：用户登录。
- `POST /api/v1/email/send-code`：发送邮箱验证码。

### 需要认证

- `GET /api/v1/user/info`：获取用户信息。
- `PUT /api/v1/user/password`：修改密码。

## 项目结构

```text
cmd/server/       # 应用入口
internal/config/  # 配置加载
internal/model/   # 数据模型
internal/pkg/     # 数据库、Redis、JWT、响应和工具函数
internal/middleware/
internal/router/
internal/handler/
docs/             # 架构、规范和实现顺序文档
```

## 核心设计

- 验证码通行证：验证码验证成功后生成一次性 `captcha_pass_token`。
- 认证流程：注册和登录接口必须校验验证码通行证。
- 会话机制：使用 JWT 承载用户身份。
- 安全基线：密码使用 bcrypt 哈希存储，接口统一响应格式。
- 可扩展边界：预留 service、repository、auth、risk 模块。

## 文档

- [架构说明](docs/ARCHITECTURE.md)
- [开发准则](docs/PROJECT-GUIDELINES.md)
- [设计文档](docs/项目文档.md)

## License

MIT License. See [LICENSE](LICENSE).
