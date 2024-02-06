package api

import (
	"github.com/gin-gonic/gin"
	"year_end_project/webcrawler/service"
)

func Member(c *gin.Context) {
	var m service.WorkBenchService
	info, err := m.GetMemberInfo()
	if err == nil {
		c.JSON(200, gin.H{
			"message": info,
		})
	}
}

func WYX(c *gin.Context) {
	var m service.WorkBenchService
	uid := "272"
	articles, err := m.GetUserArticles(uid)
	if err == nil {
		c.JSON(200, gin.H{
			"message": articles,
		})
	}

}
