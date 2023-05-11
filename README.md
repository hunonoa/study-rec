### 注意：Go语言里面都是值传递
#### 一、GOPATH简单说明
##### Linux 默认GOPATH路劲在 ~/go目录，Windows默认在%USERPROFILE%\go目录
##### 官方推荐所有项目和第三方库都放在同一个GOPATH目录下（注意：也可以将每个项目放在不同的GOPATH下）


#### 三、拉取第三方依赖库简要说明（下面的示例是本项目需要的第三方依赖）
##### 1、可以使用命令 go get "第三方库连接地址" 拉取第三方依赖
##### 2、如果 go get 拉不下来第三方库，可以使用gopm工具去拉
```bash
$ go get -v golang.org/x/text
$ go get -v golang.org/x/net/html
```
