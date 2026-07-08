package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	//type HandlerFunc func(ResponseWriter, *Request)适配器类型，将方法值转化为接口类型
	/*func (f HandlerFunc) ServeHTTP(w ResponseWriter, r*Request){
	f(w, r)
	 }   */
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}

type dollars float32
type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok { //不存在条目
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d) //两位小数
}

//`第一版 手动分发`
// func main() {
// 	db := database{"shoes": 50, "socks": 5}
// 	log.Fatal(http.ListenAndServe("localhost:8000", db))
// }
// func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	switch req.URL.Path {
// 	case "/list":
// 		for item, price := range db {
// 			fmt.Fprintf(w, "%s: %s\n", item, price)
// 		}
// 	case "/price":
// 		item := req.URL.Query().Get("item")
// 		price, ok := db[item]
// 		if !ok {//不存在条目
// 			w.WriteHeader(http.StatusNotFound) // 404
// 			fmt.Fprintf(w, "no such item: %q\n", item)
// 			return
// 		}
// 		fmt.Fprintf(w,"%s\n",price)
// 	default:
// 		w.WriteHeader(http.StatusNotFound) // 404
// 		fmt.Fprintf(w,"no such page: %s\n",req.URL)
// 	}
// }
