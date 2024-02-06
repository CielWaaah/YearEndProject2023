package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Number    int       `gorm:"unique" json:"number" form:"number"` //学号
	Admission time.Time `json:"admission" form:"admission"`         //入学时间
}

// 根据学号查用户
func (user *User) GetUserByNumber(db *gorm.DB, studentID int) (User, error) {
	// 查询
	result := db.Where("number = ?", studentID).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return *user, nil
}
