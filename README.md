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
# 编辑 config.yaml 文件配置数据库连接
```

4. 运行服务：
```bash
go run cmd/main.go
```

## API 文档

服务启动后访问：
- Swagger UI: http://localhost:8080/swagger/index.html
- API 文档: http://localhost:8080/docs

## 项目结构

```
api-main/
├── app/
│   ├── controllers/
│   ├── models/
│   ├── services/
│   └── utils/
├── config/
├── cmd/
│   └── main.go
└── README.md
```

## 贡献指南

1. Fork 项目
2. 创建新分支 (`git checkout -b feature/your-feature`)
3. 提交修改 (`git commit -am 'Add some feature'`)
4. 推送分支 (`git push origin feature/your-feature`)
5. 创建 Pull Request
