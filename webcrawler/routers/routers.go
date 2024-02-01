package routers

import (
	"github.com/gin-gonic/gin"
	"year_end_project/webcrawler/handler"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//路由
	v1 := r.Group("api/v1")

	a := v1.Group("/canteen")
	{
		a.POST("/get", handler.GetCanteenInfo)
	}
	b := v1.Group("/dorm")
	{
		b.GET("/get/area", handler.GetDormInfo)
	}
	return r
}
