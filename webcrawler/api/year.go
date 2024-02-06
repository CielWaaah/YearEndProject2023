package api

import (
	"github.com/gin-gonic/gin"
	"year_end_project/webcrawler/response"
	"year_end_project/webcrawler/service"
)

// 获取过去一年的年级（大一/二/三）到（大二/三/四）
func Year(c *gin.Context) {
	var user service.UserService
	var year service.YearService

	err := c.ShouldBind(&user)
	if err == nil {
		a := year.GetGrade(user.Number)
		c.JSON(200, response.Response{
			Status: 200,
			Data:   a,
			Msg:    "年级获取成功",
		})
	}
}

// 入学到现在过了多少天
func Days(c *gin.Context) {
	var user service.UserService
	var year service.YearService

	err := c.ShouldBind(&user)
	if err == nil {
		days := year.GetDays(user.Number)
		c.JSON(200, response.Response{
			Status: 200,
			Data:   days,
			Msg:    "天数计算成功",
		})
	}
}

// 入学到现在度过了多少个冬天
func Winters(c *gin.Context) {
	var user service.UserService
	var year service.YearService

	err := c.ShouldBind(&user)
	if err == nil {
		w := year.GetWinters(user.Number)
		c.JSON(200, response.Response{
			Status: 200,
			Data:   w,
			Msg:    "冬天计算成功",
		})
	}
}
