package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
	"time"
	"year_end_project/webcrawler/model"
)

func GetCanteenInfo(c *gin.Context) {
	var user model.User

	err := c.ShouldBind(&user)

	ji, rp := model.NewCCNUClient(user.Number, user.Password)

	if err == nil {
		pt, _ := model.Index(ji, rp)

		//Authorization:
		//Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJFMzJFOTBENEJCODk5MTNCRTA1MDAwMDAwMDAwMjRGNiIsImV4cCI6MTcwNzI5OTkxMX0.76zbJKc1obVTsRu1PtUwAxezvVeCmDdxJ8cBj40ppCA
		autho := "Bearer " + pt

		//Cookie:
		//JSESSIONID=08997B7DB45C90356905AFE484C1184B; routeportal=5fb17c08d7829466342e4fe60bd21884; PORTAL_TOKEN=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJFMzJFOTBENEJCODk5MTNCRTA1MDAwMDAwMDAwMjRGNiIsImV4cCI6MTcwNzI5OTkxMX0.76zbJKc1obVTsRu1PtUwAxezvVeCmDdxJ8cBj40ppCA
		cookie := "JSESSIONID=" + ji + "; routeportal=" + rp + "; PORTAL_TOKEN=" + pt

		// 获取当前时间
		currentTime := time.Now()
		// 将时间转换为字符串
		timeString := currentTime.Format("2006-01-02")

		firstDayOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
		// 将时间转换为字符串
		previousMonthString := firstDayOfMonth.Format("2006-01-02")

		fmt.Println(timeString)
		fmt.Println(previousMonthString)

		body, _ := getCanteenDataWithCookie("http://one.ccnu.edu.cn/ecard_portal/query_trans", autho, cookie, previousMonthString, timeString)
		c.JSON(200, gin.H{
			"message": body,
		})
	}
}

func getCanteenDataWithCookie(dataURL, authorization, cookie, startT, endT string) (string, error) {
	requestBody := "limit=10&page=1&tranType=&start=" + startT + "&end=" + endT

	// 创建包含Cookie的请求
	req, err := http.NewRequest("POST", dataURL, strings.NewReader(requestBody))
	if err != nil {
		return "", err
	}

	// 设置请求头中的Cookie
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Cookie", cookie)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 处理响应...
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
