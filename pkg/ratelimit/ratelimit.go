package ratelimit

import (
	"golang-docker-demo/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
)

func RetiverRatelimit(rateLimiter *limiter.Limiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ipClient := ctx.ClientIP()
		limiterCtx, err := rateLimiter.Get(ctx, ipClient)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, handler.Reponse{
				Message: "ratelimiter faile",
			})
			return
		}
		if limiterCtx.Reached {
			ctx.JSON(http.StatusBadRequest, handler.Reponse{
				Message: "Yêu cầu quá nhiều lần. Vui lòng thử lại sau!",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
