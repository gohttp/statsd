package main

import stats "github.com/gohttp/statsd"
import "github.com/statsd/client"
import "github.com/gohttp/app"
import "net/http"
import "time"

func main() {
	a := app.New()

	statsd, _ := statsd.Dial(":8125")

	a.Use(stats.New(statsd))

	a.Get("/", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello"))
		res.Write([]byte(" world"))
	}))

	a.Get("/slow", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Second)
		res.Write([]byte("hello"))
		res.Write([]byte(" world"))
	}))

	a.Listen(":3000")
}
