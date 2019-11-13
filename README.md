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

### 目录结构

```bash
├── app                 应用，API聚合、Web应用
│   ├── console         控制台
│   │   ├── api         go.micro.api.*，API
│   │   └── web         gin框架
│   ├── mobile          移动端
│   └── openapi         开放API
├── deploy              部署
├── doc                 文档资源
├── gateway             网关，自定义micro
├── pkg                 公共资源包
└── srv                 基础服务
    ├── account         账户服务，领域模型整洁架构示例
    │   ├── domain              领域
    │   │   ├── model           模型
    │   │   ├── repository      存储接口
    │   │   │   └── persistence ①存储接口实现   
    │   │   └── service         领域服务
    │   ├── interface           接口
    │   │   ├── handler         micro handler接口
    │   │   └── persistence     ②存储接口实现
    │   └── usecase             应用用例
    │       ├── event           消息事件
    │       ├── service         应用服务
    ├── example         micro srv不同场景示例
    └── pb              基础服务协议统一.proto
```
