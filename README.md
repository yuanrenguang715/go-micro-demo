# go-micro-demo

### go-mirco proto文件编译
required: 
- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://github.com/golang/protobuf)
- [protoc-gen-micro](https://github.com/micro/protoc-gen-micro)

protoc --micro_out=. --go_out=. *.proto

### server启动
配置config.ini

### client启动
配置config.ini
