package initialize

import (
	"context"
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/childelins/go-gin-api/global"
	myTracer "github.com/childelins/go-gin-api/pkg/tracer"
	"github.com/childelins/go-gin-api/proto"
)

func InitGRPCClient() error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(OpenTracingClientInterceptor(global.Tracer)))
	opts = append(opts, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))

	//conn, err := grpc.Dial("127.0.0.1:50051", opts...)
	conn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s",
		global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port,
		global.ServerConfig.LecturerSrvInfo.Name),
		opts...)

	if err != nil {
		return err
	}
	//defer conn.Close()

	global.LecturerSrvClient = proto.NewLecturerClient(conn)
	return nil
}

func OpenTracingClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		span := opentracing.SpanFromContext(ctx)
		/*
			span, _ := opentracing.StartSpanFromContext(ctx,
				"call gRPC",
				opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
				ext.SpanKindRPCClient)
			defer span.Finish()
		*/

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		err := tracer.Inject(span.Context(), opentracing.TextMap, myTracer.MDReaderWriter{MD: md})
		if err != nil {
			span.LogFields(log.String("inject-error", err.Error()))
			return err
		}

		newCtx := metadata.NewOutgoingContext(ctx, md)
		err = invoker(newCtx, method, req, reply, cc, opts...)
		if err != nil {
			span.LogFields(log.String("call-error", err.Error()))
		}

		return err
	}
}
