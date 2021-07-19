package router

import (
	v1 "starlight-wallpaper/api/v1"

	"github.com/gin-gonic/gin"
)

func InitializationBingRouter(r *gin.RouterGroup) {
	r.GET("/today", v1.BingToday)
	r.GET("/today/:index", v1.BingTodayByIndex)
	r.GET("/history/rand", v1.BingHistoryRand)
}

func InitializationStarlightRouter(r *gin.RouterGroup) {
	r.GET("/:type", v1.RandByType)
}
