# 操作日志

## 项目初始化

生成go.mod

```
go mod init "github/com/mumingv/gin-blog"
```

创建目录和文件

```
mkdir conf controller dao logger models routers settings static templates util
touch conf/config.yaml
touch main.go
```

## 配置文件定义、读取、变更监控

配置文件&对应数据结构定义

```
vim config.yaml
touch settings/settings.go settings/settings_test.go
```

引入第三方库viper读取配置文件

```
go get github.com/spf13/viper
```

引入第三方库fsnotify监控配置文件变更

```
go get github.com/fsnotify/fsnotify
```

引入第三方库gin框架处理Http请求

```
go get github.com/gin-gonic/gin
```

验证配置文件读取&变更

```
go mod tidy
go run main.go
http://127.0.0.1:8080/hello
```