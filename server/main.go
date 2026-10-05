// Metro Launcher 静态文件服务器
// 自包含、无外部依赖，监听指定端口服务 Web 目录。
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	dir := flag.String("dir", ".", "web root directory")
	port := flag.String("port", "5080", "listen port")
	flag.Parse()

	if fi, err := os.Stat(*dir); err != nil || !fi.IsDir() {
		log.Fatalf("web directory not found or not a directory: %s", *dir)
	}

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Printf("Metro Launcher serving %s on :%s", *dir, *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatal(err)
	}
}
