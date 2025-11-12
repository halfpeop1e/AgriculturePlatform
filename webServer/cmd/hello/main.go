package main

import (
	"fmt"
	"go-agriculture/internal/config"
	httpserver "go-agriculture/internal/http_server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.GetConfig()
	host := conf.MainConfig.Host
	port := conf.MainConfig.Port
	go func() {
		if err := httpserver.GE.Run(fmt.Sprintf("%s:%d", host, port)); err != nil {
			return
		}
	}()
	// 设置信号监听
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 等待信号
	<-quit
}
