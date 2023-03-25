package Models

import (
	"time"
)

func (b *ChatCSIfElse) TableName() string {
	return "ChatCSIfElse"
}

type ChatCSIfElse struct {
	Username       string    `json:"username"`
	VisitDateTime  time.Time `json:"visit_date_time"`
	MessageIn      string    `json:"message_in"`
	MessageOut     string    `json:"message_out"`
	MessageInTime  time.Time `json:"message_in_time"`
	MessageOutTime time.Time `json:"message_out_time"`
}

type ChatResult struct {
	Welcome string `json:"welcome"`
	Service string `json:"service"`
}
