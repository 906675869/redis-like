package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"redis-like/config"
	"redis-like/protocol"
	"redis-like/storage"
	"syscall"
	"time"
)

func main() {
	log.Println("server starting,env init")
	// 环境变量配置
	config.EnvConfigInstance()
	log.Println("env init success,storage init")
	// 存储引擎初始化
	stor := storage.StorageInstance()
	log.Println("storage init success,protocol starting ...")
	// 网络初始化
	server := protocol.Start()
	log.Println("protocol started ...")
	// 监听操作系统信号量，优雅退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("access signal : " + s.String())
			// tcp链接关闭
			server.Stop()
			// 存储引擎关闭
			stor.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
