package Models

import (
	"time"
)

func (b *ChatCSIfElse) TableName() string {
	return "ChatCSIfElse"
}

type ChatCSIfElse struct {
	Username      string    `json:"username"`
	VisitDateTime time.Time `json:"visit_date_time"`
	MessageIn     string    `json:"message_in"`
	MessageOut    string    `json:"message_out"`
}

type ChatResult struct {
	Welcome string `json:"welcome"`
	Service string `json:"service"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
type Feedback struct {
	RefId          string    `json:"ref_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Feedback       string    `json:"feedback"`
	FeedbackStatus string    `json:"feedback_status"`
	TimeStamp      time.Time `json:"time_stamp"`
}

type GetFeedBack struct {
	Status   string `json:"status"`
	FeedBack []Feedback
}

type GetAllFeedBack struct {
	OpenEscalateFeedback GetFeedBack
	CloseFeedback        GetFeedBack
	OpenFeedback         GetFeedBack
}
