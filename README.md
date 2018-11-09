# go-micro-demo

### go-mirco proto文件编译
required: 
- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://github.com/golang/protobuf)  (切换分支 git checkout v1.0.0)
- [protoc-gen-micro](https://github.com/micro/protoc-gen-micro)

protoc --micro_out=. --go_out=. *.proto

### server启动
配置config.ini

### client启动
配置config.ini


### consul
http://192.168.0.213:8500
##### consul api
```bash
 curl http://192.168.0.213:8500/v1/health/checks/<service name>
 curl http://192.168.0.213:8500/v1/catalog/services
 curl http://192.168.0.213:8500/v1/catalog/service/<service name>
 ```