package ginutils

import "github.com/gin-gonic/gin"

func InitGin() *gin.Engine {
	r := gin.Default()
	r.GET("ping", GetHander())
	r.POST("create", postPerson())
	return r
}
