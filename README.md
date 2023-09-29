# COMP90018 Backend

### Directory Introduction

##### Directory Structure
```
/COMP90018_Backend
  /api
    /routesHandler
      userRoute.go
      memoRoute.go
  /models
    user.go
  /repository
    user.go
  /services
    user.go
  /server
    main.go
  Dockerfile
  README.md
```
##### Explanation
- /cmd/myapp/main.go: 项目的入口点。main.go文件通常相对较小，并包含了项目的启动代码。
- /pkg: 这里放的是可以被外部应用使用的库代码（即，其他人可能会想要导入你的代码并运行它）
- /api: 这里放的是公共的 API 定义和协议，例如数据格式、JSON schemas 等。
- /web: 这里放的是与 HTTP handler、路由 相关的代码。
- /models: 这个目录存放的是数据模型定义，例如 user.go。
- /repository: 这个目录存放的是数据库交互相关的代码，如 user.go 包含了所有与用户相关的数据库交互函数。
- /services: 这个目录存放的是核心业务逻辑，或者说是应用的“服务”层。在这个目录下包含了与用户相关的业务逻辑，例如注册新用户、验证用户等。
- Dockerfile: 用于 Docker 的配置文件。


# MySQL

### Users Table
```
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    username VARCHAR(20) NOT NULL UNIQUE,
    email VARCHAR(40) NOT NULL UNIQUE,
    password CHAR(64) NOT NULL
);
```

### Memos Table
```
CREATE TABLE memos (
    memoId CHAR(36) PRIMARY KEY,
    userId CHAR(36) NOT NULL,
    title VARCHAR(20) NOT NULL,
    content VARCHAR(200) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```
