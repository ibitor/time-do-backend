package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"timedo/core"
	"timedo/dao"
	"timedo/sql"
)

var DB = sql.Test()

func main() {
	route := gin.Default()
	route.Use(Cors())
	route.GET("/testing", startPage)
	route.GET("/list", listAllItems)
	route.GET("/item", deleteItem)
	core.CronPush()

	route.Run(":8085")
}

func listAllItems(c *gin.Context) {
	log.Println("list")
	c.JSON(0, dao.GetAllItem())
}

func deleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err == nil {
		log.Println("delete")
		DB.Delete(&sql.Item{}, id)
	}

}

func startPage(c *gin.Context) {
	var items sql.Item
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	// 查看更多：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
	//item := sql.Item{Name: "Jinzhu", Count: 18, ProduceDate: time.Now()}
	//DB.Create(&item) // 通过数据的指针来创建
	//fmt.Println(items)
	if c.ShouldBind(&items) == nil {
		log.Println(items.Name)
		log.Println(items.Count)
		log.Println(items.ProduceDate)
		log.Println(items.SafeDay)
	}
	DB.Create(&items) // 通过数据的指针来创建s
	c.String(200, "Success")
}

// Cors 处理跨域请求,支持options访问
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
