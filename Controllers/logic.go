package Controllers

import (
	"fmt"
	"net/http"
	"time"

	"mvc_63050096_2565_2/Models"

	"github.com/gin-gonic/gin"
)

//welcome user
func VisitChatCSIfElse(c *gin.Context) {

	var req Models.ChatCSIfElse
	c.ShouldBindJSON(&req)

	incoming_user := Models.ChatCSIfElse{
		Username:      req.Username,
		VisitDateTime: time.Now(),
		MessageIn:     req.MessageIn,
	}
	fmt.Println("================ Req ===============")
	fmt.Println(incoming_user)

	//variable for checing datetime
	time_avaible_01 := "08:00:00"
	time_avaible_02 := "17:00:00"
	formatTime := incoming_user.VisitDateTime.Format("15:04:05")

	var service string
	var welcome string

	//call pugin
	fmt.Println("=========== Get Username ===========")
	user, er := Models.GetUsername(incoming_user)
	if er != nil {
		fmt.Println("ไม่พบข้อมูล")
	}

	//check id user exist
	if user.Username == "" { //username not exist

		//create new user
		fmt.Println("========== Create Username =========")
		er = Models.CreateUsername(&incoming_user)
		if er != nil {
			c.JSON(http.StatusBadRequest, "เพิ่มข้อมูลไม่สำเร็จ")
			return
		} else {
			fmt.Println("เพิ่มข้อมูลสำเร็จ")
			fmt.Println("========== ChatCSIfElse =========")
			welcome = "Welcome " + req.Username + " to ChatCSIfElse, the best chat AI in the world! What can I help you?"

			//add history
			history := Models.ChatCSIfElse{
				Username:      user.Username,
				MessageIn:     req.MessageIn,
				MessageOut:    welcome,
				VisitDateTime: time.Now(),
			}

			err := Models.UpdateHistory(&history)
			if err != nil {
				c.JSON(http.StatusBadRequest, "เพิ่มข้อมูลไม่สำเร็จ")
				return
			}

			user, er = Models.GetUsername(incoming_user)
			if er != nil {
				c.JSON(http.StatusBadRequest, "ไม่พบข้อมูล")
				return
			}

			//check time
			if formatTime > time_avaible_01 && formatTime < time_avaible_02 {
				fmt.Println("time is 08.00-17.00")
				service = "That is interesting " + user.Username + ", that you said " + user.MessageIn + ". I will send this message to someone else very soon. Anything else I can help?"
			} else {
				fmt.Println("time is 17.00-08.00")
				service = "Sorry, we are out of service in this moment"
			}

			//add history
			history = Models.ChatCSIfElse{
				Username:      user.Username,
				MessageIn:     req.MessageIn,
				MessageOut:    service,
				VisitDateTime: time.Now(),
			}

			err = Models.UpdateHistory(&history)
			if err != nil {
				c.JSON(http.StatusBadRequest, "เพิ่มข้อมูลไม่สำเร็จ")
				return
			}

			result := Models.ChatResult{
				Welcome: welcome,
				Service: service,
			}

			//print output
			c.JSON(http.StatusOK, result)
			return
		}
	} else {

		fmt.Println("=========== ChatCSIfElse ===========")
		welcome = "Welcome again " + user.Username + "! Anything else today?"

		//add history
		history := Models.ChatCSIfElse{
			Username:      user.Username,
			MessageIn:     req.MessageIn,
			MessageOut:    welcome,
			VisitDateTime: time.Now(),
		}

		err := Models.UpdateHistory(&history)
		if err != nil {
			c.JSON(http.StatusBadRequest, "เพิ่มข้อมูลไม่สำเร็จ")
			return
		}

		user, er = Models.GetUsername(incoming_user)
		if er != nil {
			c.JSON(http.StatusBadRequest, "ไม่พบข้อมูล")
			return
		}

		//check time
		if formatTime > time_avaible_01 && formatTime < time_avaible_02 {
			fmt.Println("time is 08.00-17.00")
			service = "That is interesting " + user.Username + ", that you said " + user.MessageIn + ". I will send this message to someone else very soon. Anything else I can help?"
		} else {
			fmt.Println("time is 17.00-08.00")
			service = "Sorry, we are out of service in this moment"
		}

		//add history
		history = Models.ChatCSIfElse{
			Username:      user.Username,
			MessageIn:     req.MessageIn,
			MessageOut:    service,
			VisitDateTime: time.Now(),
		}

		err = Models.UpdateHistory(&history)
		if err != nil {
			c.JSON(http.StatusBadRequest, "เพิ่มข้อมูลไม่สำเร็จ")
			return
		}

		result := Models.ChatResult{
			Welcome: welcome,
			Service: service,
		}

		//print output
		c.JSON(http.StatusOK, result)
		return
	}

}
