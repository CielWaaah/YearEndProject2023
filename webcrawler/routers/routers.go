package routers

import (
	"github.com/gin-gonic/gin"
	"year_end_project/webcrawler/handler"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//路由
	v1 := r.Group("api/v1")

	//食堂

	a := v1.Group("/canteen")
	{
		a.POST("/get", handler.GetCanteenInfo)
	}
	//宿舍
	b := v1.Group("/dorm")
	{
		b.GET("/get/area", handler.GetDormInfo)
	}
	//图书馆
	c := v1.Group("/library")
	{
		c.POST("/get", handler.Library)
	}

	return r
}
