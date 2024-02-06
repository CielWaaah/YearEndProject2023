package api

import (
	"github.com/gin-gonic/gin"
	"time"
	"year_end_project/webcrawler/service"
)

func GetCanteenInfo(c *gin.Context) {
	var user service.UserService
	var canteen *service.CanteenService

	err := c.ShouldBind(&user)

	if err == nil {
		cookie, autho := user.GetCookieAndAuth(user.Number, user.Password)

		// 获取当前时间
		currentTime := time.Now()
		// 将时间转换为字符串
		timeString := currentTime.Format("2006-01-02")

		// 获取今年的第一天
		firstDay := time.Date(currentTime.Year(), 1, 1, 0, 0, 0, 0, currentTime.Location())
		// 将时间转换为字符串
		firstString := firstDay.Format("2006-01-02")

		body, _ := canteen.GetCanteenDataWithCookie("http://one.ccnu.edu.cn/ecard_portal/query_trans", autho, cookie, firstString, timeString)

		c.JSON(200, gin.H{
			"message": body,
		})
	}
}
