package modelchatmessage

import (
	"mine-chat/common"
	"time"
)

const EntityName = "chat_message"

type ChatMessage struct {
	common.EntityModel `json:",inline"`
	RoomChatId         int                `json:"-" gorm:"column:room_chat_id;"`
	RoomChat           string             `json:"room_chat_id" gorm:"-"`
	Message            string             `json:"message" gorm:"column:message;"`
	Type               *string            `json:"type,omitempty" gorm:"column:type;"`
	PostByUserId       int                `json:"-" gorm:"column:post_by_user;"`
	PostByUser         *common.SimpleUser `json:"post_by_user" gorm:"preload:false;""`
	Status             int                `json:"status" gorm:"column:status;"`
	ReadByRecipients   *ReadByRecipients  `json:"read_by_recipients" gorm:"read_by_recipients;"`
	CreatedAt          *time.Time         `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt          *time.Time         `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (data *ChatMessage) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DB_TYPE_MESSAGE)
	data.RoomChat = common.GenUID(data.RoomChatId, common.DB_TYPE_ROOM)
	if u := data.PostByUser; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

func (ChatMessage) TableName() string {
	return "chat_message"
}

type ChatMessageRequest struct {
	Id         string `json:"id,omitempty"`
	Message    string `json:"message" gorm:"column:message;"`
	Type       string `json:"type,omitempty" gorm:"column:type;"`
	PostByUser string `json:"post_by_user" gorm:"column:post_by_user;"`
}
