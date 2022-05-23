package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {

}

var (
	ServiceName = "golang-docker-demo"
	port        = "8100"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf("ping ok %s", ServiceName)))
	})
	fmt.Println("start service with port: ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
