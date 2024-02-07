package api

import (
	"github.com/gin-gonic/gin"
	"year_end_project/webcrawler/service"
)

func GetDormInfo(c *gin.Context) {
	var dorm service.DomitoryService
	body, err := dorm.GetAreaNameAndID()
	if err == nil {
		c.JSON(200, gin.H{
			"message": body,
		})
	}
}

//func Money(c *gin.Context) {
//	var dorm service.DomitoryService
//	var user service.UserService
//
//	err := c.ShouldBind(&user)
//
//	if err == nil {
//		//dorm.GetMoney()
//	}
//}
