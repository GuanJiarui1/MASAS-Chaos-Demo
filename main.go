package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 获取环境变量中的端口，默认为8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 定义一个简单的HTTP处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		message := fmt.Sprintf("[%s] Hello from MASAS-Chaos Demo App! Your request has been processed.\n", hostname)
		w.Write([]byte(message))
	})

	// 定义一个健康检查端点
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// 启动HTTP服务器
	address := ":" + port
	log.Printf("Starting server on %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}
