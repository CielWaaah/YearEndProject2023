package handler

import (
	"github.com/gin-gonic/gin"
	"year_end_project/webcrawler/model"
)

func Member(c *gin.Context) {
	info, err := model.GetMemberInfo()
	if err == nil {
		c.JSON(200, gin.H{
			"message": info,
		})
	}
}
