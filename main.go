package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 定义指标
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 主接口
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.WithLabelValues("/").Inc()
		hostname, _ := os.Hostname()
		message := fmt.Sprintf("[%s] Hello from MASAS-Chaos Demo App!\n", hostname)
		w.Write([]byte(message))
	})

	// 健康检查
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// 暴露 Prometheus 指标接口
	http.Handle("/metrics", promhttp.Handler())

	// 启动服务
	address := ":" + port
	log.Printf("Server started on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
