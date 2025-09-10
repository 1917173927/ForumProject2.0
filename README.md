# API-Main 项目-B

## 项目简介

API-Main 是一个基于 Gin 框架的 RESTful API 服务，提供用户管理、帖子管理和举报管理功能。

## 功能特性

- 用户注册、登录
- 帖子创建、查询、更新、删除
- 举报功能
- 管理员审核举报

## 技术栈

- Go 1.18+
- Gin Web 框架
- GORM ORM
- MySQL 数据库

## 安装与运行

### 前置要求

1. 安装 Go 1.18+
2. 安装 MySQL 5.7+
3. 配置数据库连接信息

### 安装步骤

1. 克隆项目：
```bash
git clone https://github.com/your-repo/api-main.git
cd api-main
```

2. 安装依赖：
```bash
go mod download
```

3. 配置数据库：
```bash
cp config.example.yaml config.yaml
# 编辑 config/config.yaml 文件配置数据库连接
```

4. 运行服务：
```bash
go run cmd/main.go
```

## API 文档

访问 [Swagger UI](http://localhost:8080/swagger/index.html) 查看完整的 API 文档。
- 请求路径: /register
- 请求头: Content-Type: application/json
- 请求体示例:
```json
{
  "username": "newuser",
  "password": "newpassword123",
  "user_type": 1,
  "name": "New User"
}
```
- 成功响应示例:
```json
{
  "code": 200,
  "msg": "OK",
  "data": {
    "username": "newuser",
    "password": "newpassword123",
    "user_type": 1,
    "name": "New User"
  }
}
```
- 错误响应示例:
```json
{
  "code": 400,
  "msg": "missing or invalid fields",
  "data": null
}
```

### 帖子管理接口

#### 创建帖子
- 请求方法: POST
- 请求路径: /posts
- 请求头: Content-Type: application/json
- 请求体示例:
```json
{
  "content": "这是一个测试帖子",
  "user_id": 123
}
```
- 成功响应示例:
```json
{
  "code": 200,
  "msg": "OK",
  "data": {
    "id": 456,
    "content": "这是一个测试帖子",
    "user_id": 123
  }
}
```
- 错误响应示例:
```json
{
  "code": 400,
  "msg": "content and user_id are required",
  "data": null
}
```

#### 获取帖子列表
- 请求方法: GET
- 请求路径: /posts
- 成功响应示例:
```json
{
  "code": 200,
  "msg": "OK",
  "data": {
    "posts": [
      {
        "post_id": 456,
        "content": "这是一个测试帖子",
        "user_id": 123
      },
      {
        "post_id": 789,
        "content": "另一个帖子",
        "user_id": 456
      }
    ]
  }
}
```

#### 删除帖子
- 请求方法: DELETE
- 请求路径: /posts?post_id=456&user_id=123
- 成功响应示例:
```json
{
  "code": 200,
  "msg": "OK",
  "data": null
}
```
- 错误响应示例:
```json
{
  "code": 404,
  "msg": "post not found or unauthorized",
  "data": null
}
```

#### 更新帖子
- 请求方法: PUT
- 请求路径: /posts
- 请求头: Content-Type: application/json
- 请求体示例:
```json
{
  "post_id": 456,
  "user_id": 123,
  "content": "更新后的帖子内容"
}
```
- 成功响应示例:
```json
{
  "code": 200,
  "msg": "OK",
  "data": {
    "id": 456,
    "content": "更新后的帖子内容",
    "user_id": 123
  }
}
```

### 举报管理接口

#### 创建举报
- 请求方法: POST
- 请求路径: /reports
- 请求头: Content-Type: application/json
- 请求体示例:
```json
{
  "post_id": 123,
  "user_id": 456,
  "reason": "不当内容",
  "type": "spam"
}
```
- 成功响应示例:
```json
{
  "code": 200,
  "msg": "OK",
  "data": {
    "id": 789,
    "post_id": 123,
    "user_id": 456,
    "reason": "不当内容",
    "status": 0,
    "type": "spam"
  }
}
```
- 错误响应示例:
```json
{
  "code": 400,
  "msg": "post_id and user_id are required",
  "data": null
}
```

#### 获取举报列表
- 请求方法: GET
- 请求路径: /reports
- 成功响应示例:
```json
{
  "code": 200,
  "msg": "OK",
  "data": {
    "reports": [
      {
        "id": 789,
        "post_id": 123,
        "user_id": 456,
        "reason": "不当内容",
        "status": 0,
        "type": "spam"
      },
      {
        "id": 790,
        "post_id": 124,
        "user_id": 457,
        "reason": "违规内容",
        "status": 1,
        "type": "violence"
      }
    ]
  }
}
```

#### 获取待处理举报
- 请求方法: GET
- 请求路径: /reports/pending
- 成功响应示例:
```json
{
  "code": 200,
  "msg": "OK",
  "data": {
    "reports": [
      {
        "id": 789,
        "post_id": 123,
        "user_id": 456,
        "reason": "不当内容",
        "status": 0,
        "type": "spam"
      }
    ]
  }
}
```

#### 审核举报
- 请求方法: POST
- 请求路径: /reports/review
- 请求头: Content-Type: application/json
- 请求体示例:
```json
{
  "report_id": 789,
  "approval": 1,
  "user_id": 1
}
```
- 成功响应示例:
```json
{
  "code": 200,
  "msg": "OK",
  "data": null
}
```
- 错误响应示例:
```json
{
  "code": 403,
  "msg": "Only admin can review reports",
  "data": null
}
```

### 数据库配置
在`config/database/config.yaml`中配置数据库连接：
```yaml
database:
  host: "127.0.0.1"
  port: 3306
  username: "root"
  password: "your_password"
```




## 项目结构

```
api-main/
├── app/                  # 应用核心代码
│   ├── controllers/      # 控制器层
│   ├── models/           # 数据模型
│   ├── services/         # 业务逻辑层
│   └── utils/            # 工具函数
├── config/               # 配置文件
│   └── database/         # 数据库相关配置
│       ├── database.go    # 数据库连接
│       └── migrations/   # 数据库迁移文件
├── main.go               # 应用入口
└── README.md             # 项目文档
```

主要文件说明：
- `main.go`: 应用入口，初始化数据库和服务
- `app/controllers/`: 处理HTTP请求
- `app/models/`: 定义数据模型
- `app/services/`: 业务逻辑实现
- `config/database/`: 数据库连接和迁移配置

## 贡献指南

1. Fork 项目
2. 创建新分支 (`git checkout -b feature/your-feature`)
3. 提交修改 (`git commit -am 'Add some feature'`)
4. 推送分支 (`git push origin feature/your-feature`)
5. 创建 Pull Request
