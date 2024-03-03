package main

/*
$ curl http://localhost:9999
Hello World!
$ curl http://localhost:9999/abc
Hello World!
*/
import (
	"log"
	"net/http"
)

type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("hello"))
}

func main() {
	var s server
	log.Fatal(http.ListenAndServe("localhost:9999", &s))
}
