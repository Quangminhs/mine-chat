package roommodel

type ChatRoom struct {
	Id            int    `json:"id" gorm:"column:id;"`
	UserIds       string `json:"user_ids" gorm:"user_ids"`
	Type          string `json:"type" gorm:"type"`
	ChatInitiator string `json:"chat_initiator" gorm:"chat_initiator"`
}

type ChatRoomCreate struct {
	Id            string `json:"id" gorm:"-"`
	UserIds       string `json:"user_ids" gorm:"user_ids"`
	Type          string `json:"type" gorm:"type"`
	ChatInitiator string `json:"chat_initiator" gorm:"chat_initiator"`
}

func (ChatRoom) TableName() string {
	return "chat_room"
}
