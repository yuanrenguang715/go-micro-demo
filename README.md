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
│   ├── gin         gin框架
│   ├── mobile          移动端
│   └── openapi         开放API
├── deploy              部署
├── doc                 文档资源
├── gateway             网关，自定义micro/忽略
├── pkg                 公共资源包
└── srv                 基础服务/go-micro
    ├── user         账户服务
    │   ├── domain              领域
    │   ├── interface           rpc接口实现
    │   └── usecase             应用用例
    │       ├── event           消息事件
    │       ├── service         应用服务
    ├── example         micro srv不同场景示例
    └── pb              基础服务协议统一.proto
```

### GIN-web/api框架介绍
#### 为什么使用gin:
1，快，基于前缀树(radix tree)的路由策略，占用更小的内存，无须反射。
2，支持中间件，对于一个http请求，可以通过一串链式的中间件处理，然后再作出最后的应答。
3，crash-free，不会crash停服，Gin框架可以捕获http请求中的panic，并恢复。
4，routes group。更好的组织你的路由
5，内置的rendering，gin提供了简单的api使用json,xml，html,pb等
6，扩展性，可以简单的创建自己的中间件

安装: ```go get github.com/gin-gonic/gin```

### go-micro微服务架构
#### 特性：
1.务注册/发现
2.负载均衡
3.消息解码，并默认支持json以及protobuf
4.基于rpc的请求响应
5.异步的消息通讯
6.最后也是我觉得最牛的一点，接口可插拔，你可以不管用运行环境到底是使用的etcd或者consul来做服务发现，是使用http或rabbitmq做通讯，使用kafka做订阅还是用rabbitmq或者redis，是的，所有的一些都是可以插拔的，只要在运行时加入对应的启动参数，就可以使用对应的插件
7.独立部署，每个服务独立去部署，独立运行于不同的进程，使得每个服务可以根据负载做部署数量调整
8.服务降级，当系统负载增高，需要对某些关键业务做优先增配时，可以对不关键、低级别的业务做服务降级，临时关闭某些服务来保证关键业务的稳定
9.功能单一，技术选型更广，由于微服务的功能相对单一，可以针对不同的微服务采用更适合的编程语言、数据库等，例如做朋友关系、二度好友功能就考虑使用neo4j的图形数据库，底层稳定性要求高的微服务则可以通过Java来开发，并发比较高的微服务则可以考虑使用Go来开发等等
10.后期维护隔离，对功能的后期开发升级不会影响整体的应用稳定性
11.单独测试，单独开发，可以使每个微服务由不同的团队去开发、去单独测试，增加了整体开发的速度
12.复用性，多个不同的应用可以直接复用某个微服务
