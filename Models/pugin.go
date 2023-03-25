package Models

import (
	"fmt"
	"time"

	"mvc_63050096_2565_2/Config"

	_ "github.com/go-sql-driver/mysql"
)

//get username
func GetUsername(req ChatCSIfElse) (user ChatCSIfElse, err error) {
	fmt.Println("=============== User ===============")
	if err = Config.DB.Table("Test.ChatCSIfElse").
		Where("username = ?", req.Username).
		Find(&user).Error; err != nil {
		return
	}
	fmt.Println(user)
	return
}

//create non exist username
func CreateUsername(user *ChatCSIfElse) (err error) {
	fmt.Println("date ", user.VisitDateTime)
	fmt.Println("=============== Create =============")
	if err = Config.DB.Create(&user).Error; err != nil {
		return
	}
	return
}

//update History
func UpdateHistory(user *ChatCSIfElse) (err error) {
	fmt.Println("=============== Update =============")
	if err = Config.DB.Table("Test.ChatCSIfElse").
		Where("username = ?", user.Username).
		Update(map[string]interface{}{
			"visit_date_time": time.Now(),
			"message_in":      user.MessageIn,
			"message_out":     user.MessageOut,
		}).
		Error; err != nil {
		return
	}
	return
}
