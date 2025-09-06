# ForumProject2.0

## 项目简介
ForumProject2.0 是一个基于 Go 语言开发的论坛管理系统，支持用户注册、登录、发帖、举报以及管理员审核举报等功能。项目采用了 Gin 框架作为 HTTP 服务器，GORM 作为 ORM 工具，MySQL 作为数据库。

## 功能模块

### 用户模块
- **注册**：用户可以通过提供用户名和密码进行注册。
- **登录**：用户可以通过用户名和密码登录。

### 帖子模块
- **发帖**：用户可以创建帖子。
- **删除帖子**：用户可以删除自己的帖子。
- **获取帖子**：用户可以查看所有帖子。
- **更新帖子**：用户可以更新自己的帖子。

### 举报模块
- **举报帖子**：用户可以举报违规帖子，需提供举报原因。
- **查询举报进度**：用户可以查看自己举报帖子的处理状态。

### 管理员模块
- **获取未审批举报**：管理员可以查看所有未审批的举报。
- **审核举报**：管理员可以通过或拒绝举报。

## 技术栈
- **后端框架**：Gin
- **数据库**：MySQL
- **ORM**：GORM
- **依赖管理**：Go Modules

## 项目结构
```
ForumProject2.0/
├── cmd/                # 应用入口
├── internal/           # 内部模块
│   ├── admin/          # 管理员模块
│   │   ├── controller/ # 控制器层
│   │   ├── service/    # 服务层
│   │   └── repository/ # 数据库操作层
│   ├── student/        # 学生模块
│   │   ├── controller/ # 控制器层
│   │   ├── service/    # 服务层
│   │   └── repository/ # 数据库操作层
│   └── user/           # 用户模块
│       ├── controller/ # 控制器层
│       ├── service/    # 服务层
│       └── repository/ # 数据库操作层
├── go.mod              # Go Modules 配置文件
├── go.sum              # Go Modules 依赖文件
└── README.md           # 项目说明文件
```

## 环境配置

### 数据库配置
1. 确保已安装 MySQL。
2. 创建名为 `items` 的数据库：
   ```sql
   CREATE DATABASE items;
   ```
3. 在 `internal/admin/repository/db_connection.go` 中配置数据库连接信息：
   ```go
   dsn := "root:password@tcp(127.0.0.1:3306)/items?charset=utf8mb4&parseTime=True&loc=Local"
   ```

### 启动项目
1. 安装依赖：
   ```bash
   go mod tidy
   ```
2. 运行项目：
   ```bash
   go run cmd/main.go
   ```

## API 文档

### 用户模块
- **注册**
  - URL: `/api/user/reg`
  - 方法: `POST`
  - 请求体:
    ```json
    {
      "username": "testuser",
      "password": "password123"
    }
    ```

- **登录**
  - URL: `/api/user/login`
  - 方法: `POST`
  - 请求体:
    ```json
    {
      "username": "testuser",
      "password": "password123"
    }
    ```

### 帖子模块
- **发帖**
  - URL: `/api/student/post`
  - 方法: `POST`
  - 请求体:
    ```json
    {
      "content": "This is a new post."
    }
    ```

- **删除帖子**
  - URL: `/api/student/post`
  - 方法: `DELETE`
  - 请求体:
    ```json
    {
      "post_id": 1
    }
    ```

- **获取帖子**
  - URL: `/api/student/post`
  - 方法: `GET`

- **更新帖子**
  - URL: `/api/student/post`
  - 方法: `PUT`
  - 请求体:
    ```json
    {
      "post_id": 1,
      "content": "Updated content."
    }
    ```

### 举报模块
- **举报帖子**
  - URL: `/api/student/report`
  - 方法: `POST`
  - 请求体:
    ```json
    {
      "post_id": 1,
      "reason": "Inappropriate content."
    }
    ```

- **查询举报进度**
  - URL: `/api/student/report-status`
  - 方法: `GET`
  - 参数: `user_id`

### 管理员模块
- **获取未审批举报**
  - URL: `/api/admin/reports`
  - 方法: `GET`

- **审核举报**
  - URL: `/api/admin/reports`
  - 方法: `POST`
  - 请求体:
    ```json
    {
      "report_id": 1,
      "approval": 1
    }
    ```

