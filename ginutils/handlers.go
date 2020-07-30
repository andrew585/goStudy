package ginutils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"goStudy/gormodel"
	"io/ioutil"
	"log"
)

func GetHander() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ping",
		})
	}
}

func postPerson() gin.HandlerFunc {
	return func(context *gin.Context) {
		per := gormodel.Person{}
		message := context.PostForm("message")
		data, _ := ioutil.ReadAll(context.Request.Body)
		if err := json.Unmarshal(data, &per); err != nil {
			log.Fatal(err)
		}
		gormodel.CreaTeable(per)
		nick := context.DefaultPostForm("nick", "anonymous")
		context.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	}
}
