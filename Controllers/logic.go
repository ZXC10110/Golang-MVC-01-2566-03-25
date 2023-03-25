package Controllers

import (
	"fmt"
	"net/http"
	"time"

	"mvc_63050096_2565_2/Models"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
)

//welcome user
func VisitChatCSIfElse(c *gin.Context) {

	var req Models.ChatCSIfElse
	c.ShouldBindJSON(&req)

	incoming_user := Models.ChatCSIfElse{
		Username:      req.Username,
		VisitDateTime: req.VisitDateTime,
	}
	fmt.Println("================ Req ===============")
	fmt.Println(incoming_user)

	time_avaible_01 := "08:00:00"
	time_avaible_02 := "17:00:00"
	formatTime := incoming_user.VisitDateTime.Format("15:04:05")

	var service string
	var welcome string

	//call pugin
	fmt.Println("=========== Get Username ===========")
	user, er := Models.GetUsername(incoming_user)
	if er != nil {
		//if can not create return this
		c.JSON(http.StatusBadRequest, "ข้อมูลผิดพลาดโปรดลองอีกครั้ง")
		return
	} else {

		//check id user exist
		if user.Username != "" { //username not exist

			//create new user
			er = Models.CreateUsername(req)
			if er != nil {
				//if can not create return this
				c.JSON(http.StatusBadRequest, "ข้อมูลผิดพลาดโปรดลองอีกครั้ง")
				return
			} else {
				welcome = "Welcome " + req.Username + " to ChatCSIfElse, the best chat AI in the world! What can I help you?"

				//check time
				if formatTime > time_avaible_01 && formatTime < time_avaible_02 {
					fmt.Println("time is 08.00-17.00")
					service = "That is interesting " + user.Username + ", that you said " + user.MessageOut + ". I will send this message to someone else very soon. Anything else I can help?"
				} else {
					fmt.Println("time is 17.00-08.00")
					service = "Sorry, we are out of service in this moment"
				}

				result := Models.ChatResult{
					Welcome: welcome,
					Service: service,
				}

				c.JSON(http.StatusOK, result)
				return
			}
		} else {
			//print output
			welcome = "Welcome again " + user.Username + "! Anything else today?"

			//check time
			if formatTime > time_avaible_01 && formatTime < time_avaible_02 {
				fmt.Println("time is 08.00-17.00")
				service = "That is interesting " + user.Username + ", that you said " + user.MessageOut + ". I will send this message to someone else very soon. Anything else I can help?"
			} else {
				fmt.Println("time is 17.00-08.00")
				service = "Sorry, we are out of service in this moment"
			}

			result := Models.ChatResult{
				Welcome: welcome,
				Service: service,
			}

			c.JSON(http.StatusOK, result)
			return
		}
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//create unique id
func genShortUUID() (id string) {
	id = shortuuid.New()
	return id
}

func CreateFeedback(c *gin.Context) {
	//call model to request input
	var req Models.Feedback
	c.BindJSON(&req)
	feedback := Models.Feedback{
		RefId:          genShortUUID(),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Feedback:       req.Feedback,
		FeedbackStatus: "open",
		TimeStamp:      time.Now(),
	}
	//call pugin to create feedback
	er := Models.CreateFeed(&feedback)
	if er != nil {
		//if can not create return this
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		//print output
		c.JSON(http.StatusOK, feedback)
		return
	}
}

func UpdateFeedback(c *gin.Context) {
	//call model to request input
	var req Models.Feedback
	c.BindJSON(&req)

	//calculate time
	oldTime := req.TimeStamp
	currentTime := time.Now()
	dateEscalate := currentTime.Sub(oldTime)
	diff := dateEscalate.Hours()

	//condition
	if req.FeedbackStatus == "close" { //if user update by changing feedback to "close"
		//call pugin to update to database
		er := Models.UpdateFeed(&req)
		if er != nil {
			//if can not update return this
			c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
			return
		}
		//print result
		result := Models.Feedback{
			RefId:          req.RefId,
			FirstName:      req.FirstName,
			LastName:       req.LastName,
			Email:          req.Email,
			Feedback:       req.FeedbackStatus,
			FeedbackStatus: "close",
			TimeStamp:      req.TimeStamp,
		}
		c.JSON(http.StatusOK, result)
		return

	} else if diff > 168 && req.FeedbackStatus == "open" { //not modified but date > 7
		//call pugin to update to database
		er := Models.UpdateFeed(&req)
		if er != nil {
			c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
			return
		}
		//print result
		result := Models.Feedback{
			RefId:          req.RefId,
			FirstName:      req.FirstName,
			LastName:       req.LastName,
			Email:          req.Email,
			Feedback:       req.Feedback,
			FeedbackStatus: "close",
			TimeStamp:      req.TimeStamp,
		}
		c.JSON(http.StatusOK, result)
		return

	} else {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	}
}

func AdminUpdate(c *gin.Context) {
	//call model to request input
	var req Models.Feedback
	c.BindJSON(&req)

	//call pugin to update to database
	er := Models.AdminUpdate(&req)
	if er != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		//print result
		result := Models.Feedback{
			RefId:          req.RefId,
			FirstName:      req.FirstName,
			LastName:       req.LastName,
			Email:          req.Email,
			Feedback:       req.Feedback,
			FeedbackStatus: "escalate",
			TimeStamp:      req.TimeStamp,
		}
		c.JSON(http.StatusOK, result)
		return
	}
}

func GetFeedBack(c *gin.Context) {
	//call pugin
	openEscalate, er := Models.GetFeedOpenEscalate()
	if er != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	}

	//call pugin
	close, err := Models.GetFeedClose()
	if err != nil {
		return
	}

	//call pugin
	open, err := Models.GetFeedOpen()
	if err != nil {
		return
	}

	//group output
	openFeed := Models.GetFeedBack{
		Status:   "Open",
		FeedBack: open,
	}

	//group output
	openEscalateFeed := Models.GetFeedBack{
		Status:   "Open And Escalate",
		FeedBack: openEscalate,
	}

	//group output
	closeFeed := Models.GetFeedBack{
		Status:   "Close",
		FeedBack: close,
	}

	//group all output
	groupAll := Models.GetAllFeedBack{
		OpenEscalateFeedback: openEscalateFeed,
		CloseFeedback:        closeFeed,
		OpenFeedback:         openFeed,
	}

	//print output
	c.JSON(http.StatusOK, groupAll)
	return

}
