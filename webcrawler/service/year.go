package service

import (
	"time"
	"year_end_project/webcrawler/model"
)

type YearService struct {
	//
}

// 年级
func (service *YearService) GetGrade(number string) string {
	var user model.User
	var a string

	//通过学号找到用户
	user.GetUserByNumber(model.DB, number)

	sTime := user.Admission //开学时间
	nTime := time.Now()

	//年份差
	yDiff := nTime.Year() - sTime.Year()

	if yDiff == 0 {
		a = "从大一到大一"
	} else if yDiff == 1 {
		a = "从大一到大二"
	} else if yDiff == 2 {
		a = "从大二到大三"
	} else if yDiff == 3 {
		a = "从大三到大四"
	} else if yDiff == 4 {
		a = "从大四到大四"
	}

	return a
}

// 天数
func (service *YearService) GetDays(number string) int {
	var user model.User
	//通过学号找到用户
	user.GetUserByNumber(model.DB, number)

	sTime := user.Admission //开学时间
	nTime := time.Now()

	//计算差值
	duration := nTime.Sub(sTime)
	//天数
	days := int(duration.Hours() / 24)

	return days
}

// 冬天
func (service *YearService) GetWinters(number string) int {
	var user model.User
	var d int

	//通过学号找到用户
	user.GetUserByNumber(model.DB, number)

	sTime := user.Admission //开学时间
	nTime := time.Now()

	//年份差
	yDiff := nTime.Year() - sTime.Year()

	if yDiff == 0 {
		d = 1
	} else if yDiff == 1 {
		d = 2
	} else if yDiff == 2 {
		d = 3
	} else if yDiff == 3 {
		d = 4
	} else if yDiff == 4 {
		d = 5
	}

	return d
}
