# Gofar

[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

**gofar**是一个基于protobuf的web微服务框架，“gofar，you are a mature framework, you should learn to coding by yourself”。
## 特点
 - 支持proto协议go代码的自动生成。  
 - 支持ORM数据层持久化。
## 使用
1.目前需要先编译命令行工具  
```
 go build -o gofar gofar-cli/main.go  // 编译命令行工具
```  
2.修改gofar的权限或者以`./`执行
```
 chmod +x gofar //或者 ./gofar
``` 
3.例子
```
 >_: gofar version
 Gofar CLI Tool: 0.0.1
```  

4.proto 外部以来安装
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
vim ~/.bash_profile
source ~/.bash_profile