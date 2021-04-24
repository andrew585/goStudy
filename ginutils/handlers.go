package ginutils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"goStudy/gormodel"
	"goStudy/tdx"
	"io/ioutil"
	"log"
	"strconv"
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

func getTdxData() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		n := c.Query("parameter")

		dayM, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(405, gin.H{
				"status":  "err",
				"message": err.Error(),
			})
		}
		//计算atr
		atr, dataArr, err := tdx.CalculationAtr(code, dayM)
		if err != nil {
			c.JSON(405, gin.H{
				"status":  "err",
				"message": err.Error(),
			})
		}

		//将获取的数组倒叙排序
		c.JSON(405, gin.H{
			"status": "ok",
			"data":   dataArr,
			"atr":    atr,
		})
	}
}
