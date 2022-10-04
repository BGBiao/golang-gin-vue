/**
 * @File Name: routers.go
 * @Author: GoOps
 * @Email:
 * @Create Date: 2018-05-26 14:05:42
 * @Last Modified: 2018-06-02 16:06:19
 * @Description:
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//定义日志装饰器
/*
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}


*/
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)

	})
}

/*
todos json struct
api-json数据输出
*/
type Todo struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	Due       string `json:"nowtime"`
}

type Todos []Todo

/*
handler

*/

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the web server.\nurl:%q\nhost:%q\n", r.URL.Path, r.URL.Host)

}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "xuxuebiao",
			Due: time.Now().Format("2006-01-02 15:04:05")},
		Todo{Name: "xxbandy",
			Completed: true,
			Due:       time.Now().Format("2006-01-02 15:04:05")},
		Todo{Name: "xxbandy.github.io",
			Due: time.Now().Format("2006-01-02 15:04:05")},
	}

	//设置响应的header和响应码
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	/*
	   在HTTP协议中，时间都是用格林尼治标准时间来表示的，而不是本地时间。
	*/
	w.Header().Add("Date", "Wed, 21 Oct 2015 07:28:00 CST")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
	fmt.Fprintf(w, "")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintf(w, "Todo show:%s \n", todoId)

}

/*
routers

*/
type Router struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routers []Router

//define three router for "/" "/todo" and "/todos/{todoId}"
//使用数组切片定义更加丰富的路由
var routers = Routers{
	Router{
		"Index",
		"GET",
		"/",
		Index,
	},
	Router{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Router{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
}

/*
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _,route := range routers {
        //handler = route.HandlerFunc
        router.
            //Method(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            HandlerFunc(route.HandlerFunc)
    }
    return router
}
*/

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routers {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
