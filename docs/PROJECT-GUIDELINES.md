# 授权系统项目开发准则

> 本文档包含授权系统项目的具体开发准则和技术规范

## 重要说明

⚠️ **技术参考约束**

在进行项目的创建、优化、修改、添加等任何开发工作时，必须：
1. **优先参考** `docs/项目文档.md` 中定义的技术栈和设计方案
2. **严格使用** 项目文档中指定的技术或其周边生态技术
3. **禁止引入** 未在项目文档中提及的新技术栈
4. **如需变更** 技术选型，必须先与用户讨论并更新项目文档

## 项目概述

- **项目名称**：授权网站后端系统
- **技术栈**：Go 1.21 + Gin + GORM + MySQL + Redis + JWT + Zap
- **设计理念**：安全流程驱动
- **参考文档**：`docs/项目文档.md`

## 技术栈规范

### 核心技术
- **Web框架**：Gin
- **ORM**：GORM
- **数据库**：MySQL 8.x
- **缓存**：Redis
- **身份认证**：JWT (golang-jwt/jwt/v5)
- **日志**：Zap
- **配置管理**：Viper
- **密码加密**：bcrypt

### 项目结构
```
cmd/server/main.go          # 启动入口
internal/
  ├── config/               # 配置管理
  ├── model/                # 数据模型
  ├── pkg/                  # 公共工具
  ├── middleware/           # 中间件
  ├── router/               # 路由
  ├── handler/              # 控制器
  ├── service/              # 业务逻辑
  ├── repository/           # 数据访问
  ├── auth/                 # 认证模块
  └── risk/                 # 风控模块
```

## 代码规范

### 中间件顺序
```
RequestID → Logger → Recovery → CORS → Auth（按需）
```

### 响应工具
- 成功响应：`pkg.Success(c, data)`
- 错误响应：`pkg.Error(c, code, message)`

### 数据库操作
- 获取数据库实例：`pkg.GetDB()`
- 使用 GORM 的 AutoMigrate 进行迁移
- 不直接修改数据库表结构

### 日志记录
- 获取日志实例：`pkg.GetLogger()`
- 关键操作必须记录日志

### 配置管理
- 所有配置项放在 `config.yaml`
- 不在代码中硬编码密钥和密码

## 功能实现顺序

按照以下顺序实现功能，确保每个阶段完成后再进入下一阶段：

1. ✅ **项目骨架 + 配置**
   - 目录结构
   - 配置文件
   - 数据库连接
   - Redis连接
   - 日志系统

2. **验证码 + captcha_token**
   - 对接 LuckyCola
   - 生成验证码
   - 验证验证码
   - 生成 captcha_pass_token
   - 校验 captcha_pass_token

3. **邮箱验证码**
   - 发送邮箱验证码
   - 验证邮箱验证码
   - 类型区分（注册/登录/重置）
   - 失败次数限制

4. **用户注册**
   - 校验 captcha_pass_token
   - 校验邮箱验证码
   - 校验卡密（可选）
   - 风控判断
   - 创建用户

5. **用户登录**
   - 校验 captcha_pass_token
   - 验证用户名密码
   - 风控检查
   - 生成 JWT Token

6. **JWT 会话**
   - Token 生成
   - Token 验证
   - Token 刷新
   - IP/UA 验证

7. **风控模块**
   - 登录失败计数
   - 注册频率控制
   - 邮箱验证码频控
   - 异常行为检测

8. **卡密系统**
   - 卡密生成
   - 卡密验证
   - 卡密管理

9. **一机一码**
   - 设备指纹
   - 设备绑定
   - 设备管理

## API 路由规范

### 路由前缀
- 所有 API 使用 `/api/v1` 前缀
- 重大变更使用新版本号

### 路由分组
```
/api/v1
  ├── /health              # 健康检查
  ├── /captcha             # 验证码相关
  ├── /auth                # 认证相关
  ├── /email               # 邮箱相关
  ├── /user                # 用户相关（需认证）
  └── /admin               # 管理后台（需认证+权限）
```

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

## 数据模型

### 用户表 (users)
- `id` - 主键
- `email` - 邮箱（唯一索引）
- `password_hash` - 密码哈希
- `status` - 状态（1=正常，2=冻结，3=封禁）
- `last_login_at` - 最后登录时间
- `last_login_ip` - 最后登录IP
- `created_at` - 创建时间
- `updated_at` - 更新时间

### 验证码通行证表 (captcha_tokens)
- `id` - 主键
- `token` - 通行证（唯一索引）
- `captcha_id` - 验证码ID
- `ip` - IP地址
- `ua_hash` - UA哈希
- `used` - 是否已使用
- `expire_at` - 过期时间
- `created_at` - 创建时间

### 邮箱验证码表 (email_codes)
- `id` - 主键
- `email` - 邮箱
- `code` - 验证码
- `type` - 类型（register/login/reset）
- `ip` - IP地址
- `used` - 是否已使用
- `fail_count` - 失败次数
- `expire_at` - 过期时间
- `created_at` - 创建时间

## 安全机制

### 验证码通行证流程
1. 前端请求验证码
2. 用户完成验证码
3. 后端验证成功后生成 `captcha_pass_token`
4. 前端携带 token 调用注册/登录接口
5. 后端验证 token 有效性
6. token 使用后立即失效

### JWT 会话管理
- Token 包含：用户ID、登录时间、IP段、UA哈希
- 每次请求验证 IP 段和 UA
- Token 过期时间可配置
- 支持 Token 刷新机制

### 密码安全
- 使用 bcrypt 加密
- Cost 值为 10
- 只存储哈希值

## 配置文件示例

```yaml
server:
  port: 8080
  mode: debug

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
  db: 0

jwt:
  secret: your_jwt_secret_key
  expire_hours: 24

captcha:
  luckycola_api_url: https://api.luckycola.com.cn
  luckycola_app_id: your_app_id
  luckycola_app_secret: your_app_secret

email:
  smtp_host: smtp.example.com
  smtp_port: 587
  username: noreply@example.com
  password: your_email_password
```

## 开发注意事项

1. **配置文件安全**：生产环境必须修改所有密钥
2. **数据库权限**：使用专用数据库用户，不使用root
3. **Redis安全**：生产环境必须设置密码
4. **JWT密钥**：使用强随机密钥
5. **CORS配置**：生产环境限制允许的域名

## 相关文档

### 核心文档
- **`docs/项目文档.md`** - 项目需求和技术方案（技术选型的唯一参考）
- **`docs/ARCHITECTURE.md`** - 架构设计文档
- **`README.md`** - 项目说明和快速开始

### 开发规范
- **`.clinerules/auth-system-development.md`** - Cline 开发规范（通用原则）
- **`docs/PROJECT-GUIDELINES.md`** - 项目开发准则（本文档）

### 文档优先级
1. **技术选型** → 参考 `docs/项目文档.md`
2. **开发规范** → 参考 `.clinerules/auth-system-development.md`
3. **项目规范** → 参考 `docs/PROJECT-GUIDELINES.md`
4. **架构设计** → 参考 `docs/ARCHITECTURE.md`

## 文档维护

- 代码变更后及时更新 README.md
- 架构变更后及时更新 docs/ARCHITECTURE.md
- API 变更后及时更新 API 文档
- 技术选型变更后必须更新 docs/项目文档.md
- 本文档随项目演进持续更新
