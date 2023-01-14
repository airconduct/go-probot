package main

import (
	"fmt"
	"net/http"

	"github.com/airconduct/go-probot/web"
)

func main() {
	web.RegisterHandler(http.DefaultServeMux)
	fmt.Println("http://127.0.0.1:7771")
	http.ListenAndServe("127.0.0.1:7771", nil)
}
