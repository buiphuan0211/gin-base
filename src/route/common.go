package route

import (
	"gin-base/src/handler"
	"github.com/gin-gonic/gin"
)

func common(r *gin.RouterGroup) {
	var (
		g = r.Group("/ping")
		h = handler.Common{}
	)

	g.GET("", h.Ping)
}
