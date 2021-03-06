package main

import "github.com/gohttp/rpc-logger"
import "github.com/gohttp/app"
import "net/http"
import "io"
import "fmt"

func main() {
	a := app.New()

	a.Use(logger.New())

	a.Get("/", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello"))
		res.Write([]byte(" world"))
	}))

	a.Get("/yahoo", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		page, _ := http.Get("http://yahoo.com")
		defer page.Body.Close()
		io.Copy(res, page.Body)
	}))

	a.Get("/error", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(500)
		res.Write([]byte("boom"))
	}))

	fmt.Printf("Listening on :3000")

	a.Listen(":3000")
}
