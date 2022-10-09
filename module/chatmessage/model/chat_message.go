package modelchatmessage

import "time"

type ChatMessage struct {
	Id               int               `json:"id" gorm:"column:id;"`
	RomChatId        int               `json:"room_chat_id" gorm:"column:room_chat_id;"`
	Message          string            `json:"message" gorm:"column:message;"`
	Type             *string           `json:"type,omitempty" gorm:"column:type;"`
	PostByUser       int               `json:"post_by_user" gorm:"column:post_by_user;"`
	ReadByRecipients *ReadByRecipients `json:"read_by_recipients" gorm:"read_by_recipients;"`
	CreateAt         *time.Time        `json:"created_at,omitempty" gorm:"column:created_at;"`
}

func (ChatMessage) TableName() string {
	return "chat_message"
}

type ChatMessageCreate struct {
	Id         string `json:"id,omitempty"`
	Message    string `json:"message" gorm:"column:message;"`
	Type       string `json:"type,omitempty" gorm:"column:type;"`
	PostByUser string `json:"post_by_user" gorm:"column:post_by_user;"`
}
