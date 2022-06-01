package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

func showDownloadPage(w http.ResponseWriter, r *http.Request) { // http.ResponseWriter表示返回给客户端的响应对象，r表示客户端发给服务器的请求对象
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, nil)
}
func download(w http.ResponseWriter, r *http.Request) {
	//获取请求参数
	fn := r.FormValue("filename")
	//设置响应头
	header := w.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("Content-Disposition", "attachment;filename="+fn)
	//使用ioutil包读取文件
	b, _ := ioutil.ReadFile("file/" + fn)
	//写入到响应流中
	w.Write(b)
}

func main() {
	server := http.Server{Addr: ":8899"}

	http.HandleFunc("/", showDownloadPage)     //  在实际部署的时候，前后端分离，使用Nginx部署index.html页面，后端服务器只处理API接口
	http.HandleFunc("/api/download", download) // 路绝大多数路由都是默认使用前缀匹配规则，如果部分前缀相同，则按最长前缀匹配原则(即匹配的前缀越长，优先级越高，/和/download的前缀都是/，但是/download的前缀更长，/dowbload/xxx优先匹配/download)

	server.ListenAndServe()
}
