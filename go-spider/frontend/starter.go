package main

import (
	"net/http"
	"oliver/study/go-spider/frontend/controller"
)

func main() {
	http.Handle("/",http.FileServer(http.Dir("/Users/oliver/go/src/oliver/study/go-spider/frontend/view")))
	http.Handle("/search",controller.CreateSearchResultHandler("/Users/oliver/go/src/oliver/study/go-spider/frontend/view/index.html"))
	err := http.ListenAndServe(":8888", nil)
	if err !=nil{
		panic(err)
	}
}
