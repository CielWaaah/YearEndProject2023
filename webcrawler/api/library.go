package api

import (
	"github.com/gin-gonic/gin"
	"year_end_project/webcrawler/service"
)

func Library(c *gin.Context) {
	var user service.UserService
	var lib service.LibService
	err := c.ShouldBind(&user)
	si := lib.NewLibrayClient(user.Number, user.Password)

	if err == nil {
		body, _ := lib.GetLibraryRecords(si)
		c.JSON(200, gin.H{
			"message": body,
		})
	}
}
