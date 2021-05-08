package ginutils

import "github.com/gin-gonic/gin"

func InitGin() *gin.Engine {
	r := gin.Default()
	r.GET("ping", GetHander())
	r.POST("create", postPerson())
	r.GET("art", getTdxData()) //http://127.0.0.1:8080/art?code=000725&parameter=22
	return r
}
