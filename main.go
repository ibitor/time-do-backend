package main

import (
	"log"
	"net/http"
	"timedo/common"

	"github.com/gin-gonic/gin"
	"timedo/sql"
)

var db = sql.Test()

func main() {
	route := gin.Default()
	route.Use(Cors())
	route.GET("/testing", startPage)
	route.GET("/list", listAllItems)
	route.Run("192.168.2.240:8085")
}

func listAllItems(c *gin.Context) {
	var items []common.Item
	db.Find(&items)
	c.JSON(0, items)
}

func startPage(c *gin.Context) {
	var items common.Item
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	// 查看更多：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
	//item := common.Item{Name: "Jinzhu", Count: 18, ProduceDate: time.Now()}
	//db.Create(&item) // 通过数据的指针来创建
	//fmt.Println(items)
	if c.ShouldBind(&items) == nil {
		log.Println(items.Name)
		log.Println(items.Count)
		log.Println(items.ProduceDate)
		log.Println(items.SafeDay)
	}
	db.Create(&items) // 通过数据的指针来创建s
	c.String(200, "Success")
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
