# Books Manage System

## 项目简介

Books Manage System 是一个图书管理系统，提供用户注册、登录、图书查询、借阅归还等功能，并提供管理员进行图书和用户管理的接口。项目使用 Golang 开发，主要依赖 Gin、Gorm 和 Swagger。

## 项目结构

```
.
├── main.go                    # 主程序入口
├── config
│   └── config.go              # 数据库连接配置
├── controllers
│   ├── auth_controller.go     # 认证相关的控制器
│   ├── book_controller.go     # 图书相关的控制器
│   └── user_controller.go     # 用户相关的控制器
├── docs
│   ├── docs.go                # Swagger 文档生成
│   └── swagger.json           # Swagger 文档
├── middleware
│   └── auth_middleware.go     # 认证中间件
├── models
│   ├── book.go                # 图书模型
│   ├── user.go                # 用户模型
│   └── borrow.go              # 借阅记录模型
├── repositories
│   ├── book_repository.go     # 图书数据仓库
│   ├── user_repository.go     # 用户数据仓库
│   └── borrow_repository.go   # 借阅记录数据仓库
├── routes
│   └── routes.go              # 路由注册
├── services
│   ├── auth_service.go        # 认证服务
│   ├── book_service.go        # 图书服务
│   └── user_service.go        # 用户服务
└── utils
    └── utils.go               # 工具类
```

## 环境依赖

- Go 1.19+
- MySQL 数据库

## 安装与运行

### 克隆项目

```sh
git clone https://github.com/yourusername/books-manage-system.git
cd books-manage-system
```

### 安装依赖

```sh
go mod tidy
```

### 配置数据库

在项目根目录下创建 `.env` 文件，并添加以下内容：

```
DB_USER=<your_database_user>
DB_PASSWORD=<your_database_password>
DB_NAME=<your_database_name>
DB_HOST=<your_database_host>
DB_PORT=<your_database_port>
JWT_SECRET=<your_jwt_secret>
```

### 生成Swagger文档

```sh
swag init
```

### 启动项目

```sh
go run cmd/main.go
```

### 查看API文档

启动项目后，在浏览器中访问 `http://localhost:8080/swagger/index.html` 查看Swagger UI。

## 使用说明

### 用户注册

```
POST /api/v1/register
{
    "username": "example",
    "password": "password",
    "role": "user" // or "admin"
}
```

### 用户登录

```
POST /api/v1/login
{
    "username": "example",
    "password": "password"
}
```

### 添加图书（管理员权限）

```
POST /api/v1/books
Authorization: Bearer <token>
{
    "title": "Book Title",
    "author": "Author Name",
    "category": "Category",
    "quantity": 10,
    "available": 10
}
```

### 查询图书

```
GET /api/v1/books/:id
Authorization: Bearer <token>
```

### 更新图书（管理员权限）

```
PUT /api/v1/books
Authorization: Bearer <token>
{
    "id": 1,
    "title": "Updated Book Title",
    "author": "Updated Author Name",
    "category": "Updated Category",
    "quantity": 5,
    "available": 5
}
```

### 删除图书（管理员权限）

```
DELETE /api/v1/books/:id
Authorization: Bearer <token>
```

## 贡献

如果你有好的意见或建议，欢迎提出 Issue 或提交 Pull Request。

## 许可证

本项目为个人练习，遵循 Apache 2.0 许可证。
