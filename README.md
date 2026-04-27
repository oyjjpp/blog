# 脚手架

## 功能

## 配置
```yaml
go get github.com/spf13/viper
```

### mysql

### redis

### kafka

### mongodb


## 构建文档

```swag
swag fmt
swag init --parseDependency --outputTypes json
swaggerignore:"true"
```

## 代码生成

【项目目录规范】
项目根目录：blog/
配置文件: internal/pkg/config/config.go
存储中间件初始化: internal/pkg/store/store.go
内存缓存：internal/pkg/store/cache/cache.go
redis缓存：internal/pkg/store/redis/redis.go
mysql存储：internal/pkg/store/mysql/mysql.go
接口处理器：internal/pkg/v1/adapter/admin/user.go
接口处理器：internal/pkg/v1/adapter/api/user.go
业务逻辑：internal/pkg/v1/usecase/user.go
数据访问：internal/pkg/v1/usecase/repository/user.go
数据模型：internal/pkg/v1/domain/model/user.go
参数模型：internal/pkg/v1/domain/user.go
工具包：internal/pkg/utils/request.go
工具包：internal/pkg/utils/{工具名}/{工具名}.go
入口文件：cmd/api/main.go
入口文件：cmd/admin/main.go

【技术栈】Go 1.24 + Gin 框架 + GORM v2 + Redis v8 + JWT
【功能】生成用户登录接口（POST /api/v1/user/login）
【目录约束】
1. 登录接口处理器：internal/handler/user/login.go
2. 登录认证业务逻辑：internal/service/user/auth.go
3. 用户数据模型：internal/model/user.go
4. JWT 工具复用 internal/pkg/jwt/jwt.go（无需重复生成）
5. Redis 客户端复用 internal/pkg/redis/client.go（无需重复生成）

【输出要求】
1. 按上述目录结构生成每个文件的完整代码；
2. 每个文件开头注明包名（如 handler/user 包名为 "user"，service/user 包名为 "user"）；
3. 包导入需符合目录结构（如 import "your-project/internal/model"）；
4. 代码中避免硬编码包路径，使用项目实际的模块名（如 module 为 "github.com/your-name/your-project"）；
5. 输出格式：按文件路径分块展示，每块开头标注「// 文件路径：xxx」。