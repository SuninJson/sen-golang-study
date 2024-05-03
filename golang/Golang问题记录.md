# 1、go get报错：A connection attempt failed

例如尝试`go get -u github.com/go-sql-driver/mysql`命令获取mysql的时候，报错

具体错误如下：

```
go get: module github.com/go-sql-driver/mysql: Get "https://proxy.golang.org/github.com/go-sql-driver/mysql/@v/list": dial tcp 172.217.24.17:443: connectex: A connection attemp
t failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
```

具体原因：默认go的代理网站是`GOPROXY=https://proxy.golang.org,direct`，是一个外网地址，国内访问不到，因此我们需要修改代理网站。

解决：使用命令`go env -w GOPROXY=https://goproxy.cn,direct`更改代理网站，然后再重新执行go get命令即可成功

