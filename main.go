package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/childelins/go-gin-api/pkg/app"

	"github.com/gin-gonic/gin"

	"github.com/childelins/go-gin-api/global"
	"github.com/childelins/go-gin-api/initialize"
)

var (
	isVersion bool   // 是否显示版本信息
	date      string // 编译日期
	version   string // 版本号
	commit    string // git commit ID

	serviceId string // consul服务ID
)

func init() {
	initFlag()

	// 显示版本信息，需要在编译时通过ldflags设置
	if isVersion {
		printVersionInfo()
		return
	}

	var err error
	if err = initialize.InitConfig(); err != nil {
		log.Fatalf("初始化配置中心失败: %v", err)
	}
	if err = initialize.InitLogger(); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}
	serviceId = app.UUID()
	if err = initialize.InitRegistry(serviceId); err != nil {
		log.Fatalf("初始化注册中心失败: %v", err)
	}
	if err = initialize.InitTracer(); err != nil {
		log.Fatalf("初始化链路追踪失败: %v", err)
	}
	if err = initialize.InitGRPCClient(); err != nil {
		log.Fatalf("初始化GRPC客户端失败: %v", err)
	}
	if err = initialize.InitDB(); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	if err = initialize.InitTrans("zh"); err != nil {
		log.Fatalf("初始化翻译器失败: %v", err)
	}
	if err = initialize.InitSentinel(); err != nil {
		log.Fatalf("初始化限流组件失败：%v", err)
	}
}

func main() {
	log.Printf("%#v", global.ServerConfig)
	gin.SetMode(global.ServerConfig.RunMode)

	/* 访问日志设置
	logFile, err := os.OpenFile(
		fmt.Sprintf("storage/logs/access-%s.log", time.Now().Format("2006-01-02")),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("访问日志创建失败, err: %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile)
	*/

	r := initialize.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.ServerConfig.Port),
		Handler:        r,
		ReadTimeout:    global.ServerConfig.ReadTimeout,
		WriteTimeout:   global.ServerConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("http服务异常终止: %v", err)
		}
	}()

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("开始关闭http服务...")
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("关闭http服务失败: ", err)
	}
	log.Println("关闭http服务成功")

	log.Println("开始注销consul服务...")
	if err := global.Registry.DeRegister(serviceId); err != nil {
		log.Fatalf("注销consul服务[%s]失败: %v", serviceId, err)
	}
	log.Printf("注销consul服务[%s]成功", serviceId)

	log.Println("所有服务已正常退出")
}

func initFlag() {
	flag.BoolVar(&isVersion, "version", false, "版本信息")
	flag.Parse()
}

func printVersionInfo() {
	fmt.Printf("date: %s\n", date)
	fmt.Printf("version: %s\n", version)
	fmt.Printf("commit: %s\n", commit)
}
