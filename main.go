package main

import (
	"fmt"
	"golang-docker-demo/pkg/ratelimit"
	redislocal "golang-docker-demo/redis"
	"net/http"
	"time"

	"golang-docker-demo/handler"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	redisRate "github.com/ulule/limiter/v3/drivers/store/redis"
)

func init() {

}

var (
	ServiceName = "golang-demo"
	port        = "8100"
)

type Server struct {
	limiter    *limiter.Limiter
	router     *gin.Engine
	httpServer *http.Server
}

func NewServer() *Server {
	r := gin.New()
	return &Server{
		router: r,
	}
}

func main() {
	redisIn := redislocal.InitRedis()

	server := NewServer()

	// 5rp5s
	limiterRate := limiter.Rate{
		Period: 5 * time.Second,
		Limit:  5,
	}
	store, err := redisRate.NewStoreWithOptions(redisIn, limiter.StoreOptions{
		Prefix: ServiceName,
	})
	if err != nil {
		panic(err)
	}

	server.limiter = limiter.New(store, limiterRate)
	server.Router()

	server.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: server.router,
	}

	err = server.httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (s *Server) Router() {
	grouter := s.router.Group("v1")
	grouter.Use(ratelimit.RetiverRatelimit(s.limiter))
	// demo
	grouter.GET("/set", handler.Set())
	grouter.GET("/get", handler.Get())

}
