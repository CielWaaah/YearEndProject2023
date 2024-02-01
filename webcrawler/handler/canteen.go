package handler

import (
	"github.com/gin-gonic/gin"
	"time"
	"year_end_project/webcrawler/model"
)

func GetCanteenInfo(c *gin.Context) {
	var user model.User

	err := c.ShouldBind(&user)

	if err == nil {
		cookie, autho := model.GetCookieAndAuth(user.Number, user.Password)

		// 获取当前时间
		currentTime := time.Now()
		// 将时间转换为字符串
		timeString := currentTime.Format("2006-01-02")

		firstDayOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
		// 将时间转换为字符串
		previousMonthString := firstDayOfMonth.Format("2006-01-02")

		body, _ := model.GetCanteenDataWithCookie("http://one.ccnu.edu.cn/ecard_portal/query_trans", autho, cookie, previousMonthString, timeString)
		c.JSON(200, gin.H{
			"message": body,
		})
	}
}
