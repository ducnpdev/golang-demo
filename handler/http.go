package handler

import (
	"context"
	"net/http"
	"time"

	redislocal "golang-docker-demo/redis"

	"github.com/gin-gonic/gin"
)

type Reponse struct {
	Message string `json:"message"`
}

func Set() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cmdR := redislocal.RedisIn.Set(context.Background(), "set", "value-set", time.Duration(time.Second*199))
		valueResut := []byte("")
		if cmdR.Err() != nil {
			tem := []byte("redis set error \n")
			valueResut = append(valueResut, tem...)
		} else {
			tem := []byte("redis set ok \n")
			valueResut = append(valueResut, tem...)
		}
		ctx.JSON(http.StatusOK, Reponse{
			Message: string(valueResut),
		})
	}
}

func Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cmdR := redislocal.RedisIn.Get(context.Background(), "set")
		valueResut := []byte("Get ")
		if cmdR.Err() != nil {
			tem := []byte("redis get not found \n")
			valueResut = append(valueResut, tem...)
		} else {
			value, err := cmdR.Bytes()
			if err != nil {
				tem := []byte("redis get err, err:")
				tem = append(tem, []byte(err.Error())...)
				valueResut = append(valueResut, tem...)
			} else {
				text := []byte("redis get ok, value:")
				text = append(text, value...)
				valueResut = append(valueResut, text...)
			}
		}
		line := []byte("\n")
		valueResut = append(valueResut, line...)
		ctx.JSON(http.StatusOK, Reponse{
			Message: string(valueResut),
		})
	}
}
