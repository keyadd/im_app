package main

import (
	"app_ws/global"
	"app_ws/initialize"
	"app_ws/router"
	"app_ws/validator"
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	//加载配置文件
	global.GVA_VP = initialize.Viper()
	//初始化日志
	global.GVA_LOG = initialize.Zap()
	//初始化Mysql连接
	//global.SQLX_DB = initialize.Sqlx()
	//db := global.SQLX_DB.DB
	//defer db.Close()
	//初始化Redis连接
	//global.GVA_REDIS = initialize.Redis()
	//defer global.GVA_REDIS.Close()
	//雪花算法生产唯一ID
	initialize.Snowflake()
	//validator参数验证
	global.TRANS = validator.InitTrans("zh")

	// 初使化 minio client对象
	//global.MINIO = initialize.InitMinIO()

	//utils.CreateMinoBuket("userheader")

	//注册路由
	r := router.Setup()

	global.GVA_LOG.Info(fmt.Sprintf("0.0.0.0:%d", global.GVA_CONFIG.System.Addr))
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时n
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		global.GVA_LOG.Fatal("Server Shutdown: ", zap.Error(err))
	}
	global.GVA_LOG.Info("Server exiting")
}
