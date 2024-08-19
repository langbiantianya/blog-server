package cmd

import (
	"fmt"
	"log"
	"net/http"
)

func StartStaticServer(staticDir string, port int) {

	// 创建一个静态文件服务器，用于服务public目录下的文件
	fs := http.FileServer(http.Dir(staticDir))

	// 定义一个http处理器，将所有请求代理到静态文件服务器
	http.Handle("/", http.StripPrefix("/", fs))

	// 启动服务器，监听8080端口
	log.Printf("Server starting on port %d...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
