package handler

import (
	"github.com/gin-gonic/gin"
	"year_end_project/webcrawler/model"
)

func GetDormInfo(c *gin.Context) {
	body, err := model.GetAreaNameAndID()
	if err == nil {
		c.JSON(200, gin.H{
			"message": body,
		})
	}
}
