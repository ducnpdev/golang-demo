package main

import (
	"fmt"
	"golang-docker-demo/handler"
	redislocal "golang-docker-demo/redis"
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
	redislocal.InitRedis()
	http.HandleFunc("/set", handler.Set)
	http.HandleFunc("/get", handler.Get)

	fmt.Println("start service with port: ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
