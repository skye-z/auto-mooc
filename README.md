# Auto Mooc - 自动化慕课工具

[![](https://img.shields.io/badge/Go-1.21+-%2300ADD8?style=flat&logo=go)](go.work)
[![](https://img.shields.io/badge/License-GPL%20v3.0-orange)](LICENSE)

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
    - work.go       任务管理
    - script.js     页面注入脚本
- /page             控制页面
- go.mod            项目依赖
- go.sum            ~
- main.go           项目入口
```

### 引用组件

1. Playwright: 微软出品的自动化工具
3. Viper: Go语言配置文件读写库
2. Gin: Go语言主流Http服务库

## 使用

### 在 Docker 使用

```shell
docker run -d -p 80:80 -n auto-mooc skyezhang/auto-mooc:1.0.1
```

最后在浏览器输入 Docker IP (默认80端口), 将自动导航到控制页面

### 在本机使用

#### 下载二进制文件

点击右侧[Releases](https://github.com/skye-z/auto-mooc/releases)下载二进制文件, Windows需在CMD中启动, 其他系统可双击启动.

最后在浏览器输入localhost(默认80端口), 将自动导航到控制页面

#### 从源代码启动

```shell
# 拉取项目
git clone https://github.com/skye-z/auto-mooc.git auto-mooc
# 进入目录
cd auto-mooc
# 下载依赖
go mod download
# 启动项目
go run main.go
```

## 服务接口

1. `/login` 登录账户
2. `/class/list` 获取课程列表
3. `/class/select?id=课程id` 选课
4. `/class/start` 开始上课
5. `/class/stop` 停止上课
6. `/status` 查询状态
7. `/screenshot` 任务截图

## 应用配置

在应用程序首次运行后, 应用会在所在目录下生成`config.ini`配置文件, 如果不了解配置含义请勿随意修改!!!!

> 配置文件中如需填入特殊字符, 请使用``包裹整个值

```ini
[basic]
; 是否已安装, 用于标定初始化状态
install=true
; 调试模式, 开启后将输出路由日志并关闭无头模式
debug=false
; 用户代理信息, 建议用自己常用浏览器的
user-agent=test
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

[push]
; 启用推送
enable=false
; 推送地址(下面的是示例,只要符合标准的地址都可以)
url=`https://api2.pushdeer.com/message/push?pushkey={替换你自己的令牌}&text=`
```

## 自行编译

这里提供三种编译方法.

第一种是编译本机环境的二进制文件: `go build -ldflags '-s -w'`

第二种则是交叉编译: `CGO_ENABLED=0 GOOS=平台 GOARCH=架构 go build -ldflags '-s -w'`

第三种是脚本编译: `bash ./build.sh`