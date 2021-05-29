package routers

import (
	"gblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("hello", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	r.Run(utils.HttpPort)

}
