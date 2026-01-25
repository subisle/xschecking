# 授权系统后端基础框架

基于 Go + Gin 的授权系统后端骨架，遵循安全流程驱动的设计理念。

## 项目结构

```
.
├── cmd/
│   └── server/
│       └── main.go              # 主入口文件
├── internal/
│   ├── config/                  # 配置模块
│   │   └── config.go
│   ├── model/                   # 数据模型
│   │   ├── user.go             # 用户模型
│   │   ├── captcha_token.go    # 验证码通行证
│   │   └── email_code.go       # 邮箱验证码
│   ├── pkg/                     # 公共工具
│   │   ├── database.go         # 数据库连接
│   │   ├── redis.go            # Redis连接
│   │   ├── logger.go           # 日志工具
│   │   ├── jwt.go              # JWT工具
│   │   ├── response.go         # 响应工具
│   │   └── utils.go            # 通用工具
│   ├── middleware/              # 中间件
│   │   ├── request_id.go       # 请求ID
│   │   ├── logger.go           # 日志中间件
│   │   ├── recovery.go         # 崩溃恢复
│   │   ├── auth.go             # JWT认证
│   │   └── cors.go             # 跨域处理
│   ├── router/                  # 路由
│   │   └── router.go
│   ├── handler/                 # 控制器
│   │   └── handler.go
│   ├── service/                 # 业务逻辑层（待实现）
│   ├── repository/              # 数据访问层（待实现）
│   ├── auth/                    # 认证模块（待实现）
│   └── risk/                    # 风控模块（待实现）
├── config.yaml                  # 配置文件
├── go.mod                       # Go模块文件
└── README.md                    # 项目说明

```

## 技术栈

- **Web框架**: Gin
- **配置管理**: Viper
- **ORM**: GORM
- **数据库**: MySQL 8.x
- **缓存**: Redis
- **身份认证**: JWT
- **日志**: Zap
- **密码加密**: bcrypt

## 快速开始

### 1. 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis 6.0+

### 2. 安装依赖

```bash
go mod download
```

### 3. 配置文件

修改 `config.yaml` 中的配置：

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

### 4. 创建数据库

```sql
CREATE DATABASE auth_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 5. 运行项目

```bash
go run cmd/server/main.go
```

服务将在 `http://localhost:8080` 启动

## API 路由

### 公开接口

- `GET /api/v1/health` - 健康检查
- `POST /api/v1/captcha/generate` - 生成验证码
- `POST /api/v1/captcha/verify` - 验证验证码
- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录
- `POST /api/v1/email/send-code` - 发送邮箱验证码

### 需要认证的接口

- `GET /api/v1/user/info` - 获取用户信息
- `PUT /api/v1/user/password` - 修改密码

### 管理员接口

- `GET /api/v1/admin/users` - 获取用户列表
- `PUT /api/v1/admin/users/:id/status` - 更新用户状态

## 核心特性

### 1. 验证码通行证机制（captcha_pass_token）

- 防重放攻击
- 防接口跳过
- 一次性使用

### 2. 安全的用户认证

- JWT Token 认证
- IP 段验证
- UA Hash 验证
- 密码 bcrypt 加密

### 3. 完整的中间件体系

- 请求追踪（Request ID）
- 访问日志
- 崩溃恢复
- JWT 认证
- CORS 跨域

### 4. 数据模型

#### 用户表（users）
- 邮箱唯一索引
- 密码哈希存储
- 用户状态管理（正常/冻结/封禁）
- 登录记录

#### 验证码通行证表（captcha_tokens）
- 一次性使用
- IP 和 UA 绑定
- 过期时间控制

#### 邮箱验证码表（email_codes）
- 类型区分（注册/登录/重置）
- 失败次数限制
- 使用状态追踪

## 后续实现顺序

按照文档建议的顺序实现：

1. ✅ 项目骨架 + 配置
2. ⏳ 验证码 + captcha_token
3. ⏳ 邮箱验证码
4. ⏳ 用户注册
5. ⏳ 用户登录
6. ⏳ JWT 会话
7. ⏳ 风控模块
8. ⏳ 卡密系统
9. ⏳ 一机一码

## 开发规范

### 响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

### 错误码

- `0` - 成功
- `400` - 参数错误
- `401` - 未授权
- `403` - 禁止访问
- `404` - 资源不存在
- `429` - 请求过多
- `500` - 服务器错误
- `1001` - 验证码错误
- `1002` - 邮箱验证码错误
- `1003` - 用户已存在
- `1004` - 用户不存在
- `1005` - 密码错误
- `1006` - 账号冻结
- `1007` - 账号封禁
- `1008` - Token过期
- `1009` - Token无效

## 注意事项

1. **配置文件安全**：生产环境必须修改所有密钥
2. **数据库权限**：建议使用专用数据库用户，不要使用root
3. **Redis安全**：生产环境必须设置密码
4. **JWT密钥**：必须使用强随机密钥
5. **CORS配置**：生产环境需要限制允许的域名

## License

MIT
