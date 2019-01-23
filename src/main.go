package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/samoslab/samos-object-system/src/router"
)

var port string

func init() {
	flag.StringVar(&port, "port", ":9086", "监听地址和端口号")
	flag.Parse()
}

func main() {
	router := router.New()
	// 绑定server
	server := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 异步运行server
	go func() {
		server.ListenAndServe()
	}()

	// 监听系统信号
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for _ = range ch {
		server.Shutdown(context.TODO())
		break
	}
}
