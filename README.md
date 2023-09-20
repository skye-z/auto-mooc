# Auto Mooc - 自动化慕课工具

请注意: 本项目仅供交流学习自动化测试技术, 滥用造成的后果请自行承担!

## 前言

本项目涵盖了网页自动化操作、网页监控、Cookie及Storage持久化存储、Signal信号处理、基本Http服务等技术.

项目结构简单, 适合Go语言初学者或自动化技术初学者.

### 目录结构

```
- /global           公共组件目录
    - config.go     配置文件工具
    - response.go   响应工具
- /service          Http服务目录
    - mooc.go       慕课服务
    - status.go     状态检测服务
- /webkit           WebKit操作目录
    - basic.go      基础操作
    - init.go       初始化
- go.mod            项目依赖
- go.sum            ~
- main.go           项目入口
```

### 引用组件

1. Playground: 微软出品的自动化工具
3. Viper: Go语言配置文件读写库
2. Gin: Go语言主流Http服务库

## 安装

可直接下载二进制文件, 或拉取项目代码, 然后在项目目录下执行`go mod download`

## 使用

双击下载的二进制文件即可运行, 如果是拉取项目代码的, 请在项目目录下执行`go run main.go`

1. 首先访问`/login`登录账户
2. 登录后访问`/class/select`进行选课
2. 选完后访问`/class/start`开始上课

### 配置文件

在应用程序首次运行后, 应用会在所在目录下生成`config.ini`配置文件, 如果不了解配置含义请勿随意修改!!!!

```ini
[basic]
; 是否已安装, 用于标定初始化状态
install=true
; 应用监听端口, 将在此端口上提供Http服务
port=80
; 应用工作空间, 应用所有资源所在的目录
workspace=/Users/root/Workspace/auto-mooc

[mooc]
; 选课编号
class=123456
; 是否已登录, 由应用自行判断后填充
login=false
; 慕课服务地址, 请已登录后跳转的域名地址为准
path=http://mooc.com
; 持久化存储位置, 提供Cookie和本地存储的持久化
storage=/Users/root/Workspace/auto-mooc/storage.db
```

## 编译

这里提供两种编译方法.

第一种是编译本机环境的二进制文件: `go build -ldflags '-s -w'`

第二种则是交叉编译: `CGO_ENABLED=0 GOOS=平台 GOARCH=架构 go build -ldflags '-s -w'`