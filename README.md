# API-Main 项目-Bianweizheng

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

### 举报相关API变更
- 所有举报API中的`reporter_id`字段已统一改为`user_id`
- 举报请求体示例：
```json
{
  "post_id": 123,
  "user_id": 456,
  "reason": "不当内容",
  "type": "spam"
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

服务启动后访问：
- Swagger UI: http://localhost:8080/swagger/index.html
- API 文档: http://localhost:8080/docs

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
