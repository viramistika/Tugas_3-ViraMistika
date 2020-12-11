package main

import (
	"fmt"
	"net/http"

	fn "Tugas3RichyAzura/HtmlPostData/function"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", fn.RouteIndexGet)
	http.HandleFunc("/process", fn.RouteSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
