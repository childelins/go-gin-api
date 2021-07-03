package initialize

import (
	"fmt"

	"github.com/childelins/go-gin-api/global"
	"github.com/childelins/go-gin-api/pkg/tracer"
)

func InitTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		global.ServerConfig.JaegerInfo.Name,
		fmt.Sprintf("%s:%d", global.ServerConfig.JaegerInfo.Host, global.ServerConfig.JaegerInfo.Port))
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}
