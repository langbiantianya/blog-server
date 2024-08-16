package cmd

import (
	"log"
	"net/http"
)

func StartStaticServer() {
	// 设置静态文件目录
	staticDir := "./static" // 假设你的静态文件放在当前目录下的public文件夹中

	// 创建一个静态文件服务器，用于服务public目录下的文件
	fs := http.FileServer(http.Dir(staticDir))

	// 定义一个http处理器，将所有请求代理到静态文件服务器
	http.Handle("/", http.StripPrefix("/", fs))

	// 启动服务器，监听8080端口
	log.Println("Server starting on port 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
