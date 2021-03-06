package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"html/template"
	"path/filepath"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprint(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		abs, _ := filepath.Abs("web/sample/login.html")
		t, err := template.ParseFiles(abs)
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayHelloName) //设置访问的路由
	http.HandleFunc("/login", login) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
