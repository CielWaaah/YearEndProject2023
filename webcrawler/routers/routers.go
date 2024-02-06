package routers

import (
	"github.com/gin-gonic/gin"
	"year_end_project/webcrawler/api"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//路由
	v1 := r.Group("api/v1")

	//食堂
	a := v1.Group("/canteen")
	{
		a.POST("/get", api.GetCanteenInfo)
	}
	//宿舍
	b := v1.Group("/dorm")
	{
		b.GET("/get/area", api.GetDormInfo)
	}
	//图书馆
	c := v1.Group("/library")
	{
		c.POST("/get", api.Library)
	}
	//工作台
	d := v1.Group("/work_bench")
	{
		d.GET("/get", api.Member)
		d.GET("/get/wyx", api.WYX)
	}
	//学年
	e := v1.Group("/year")
	{
		e.POST("/get", api.Year)
		e.POST("/days", api.Days)
		e.POST("/winters", api.Winters)
	}

	return r
}
