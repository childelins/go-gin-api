# 项目依赖库

- [gin](https://github.com/gin-gonic/gin)
- [viper](https://github.com/spf13/viper)
- [zap](https://github.com/uber-go/zap)
- [lumberjack](https://github.com/natefinch/lumberjack)
- [validator](https://github.com/go-playground/validator)
- [gorm](https://github.com/go-gorm/gorm)
- [jwt-go](https://github.com/dgrijalva/jwt-go)
- [ratelimit](https://github.com/juju/ratelimit)
- [opentracing](https://github.com/opentracing/opentracing-go)
- [jeager](https://github.com/uber/jaeger-client-go)
- [consul](https://github.com/hashicorp/consul)
- [grpc-consul-resolver](https://github.com/mbobakov/grpc-consul-resolver)
- [nacos](https://github.com/nacos-group/nacos-sdk-go)
- [uuid](https://github.com/satori/go.uuid)
- [sentinel](https://github.com/alibaba/sentinel-golang)
- [sentinel-go-adapters](https://github.com/sentinel-group/sentinel-go-adapters)

# 编译 proto 文件
```
// 单个proto文件
protoc --go_out=plugins=grpc:. hello.proto

// 目录下所有proto文件
protoc --go_out=plugins=grpc:. *.proto
```

# 使用技巧

- 使用 ldflags 设置编译信息
> go build -ldflags "-X main.version=1.0.1 -X main.author=childelins -X date=`date +%Y-%m-%d %H:%M:%S` -X main.commit=`git rev-parse HEAD`"