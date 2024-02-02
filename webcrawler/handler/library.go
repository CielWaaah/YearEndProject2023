package handler

import (
	"github.com/gin-gonic/gin"
	"year_end_project/webcrawler/model"
)

func Library(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	si := model.NewLibrayClient(user.Number, user.Password)

	if err == nil {
		body, _ := model.GetLibraryRecords(si)
		c.JSON(200, gin.H{
			"message": body,
		})
	}
}
