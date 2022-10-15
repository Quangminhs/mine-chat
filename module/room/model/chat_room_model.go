package roommodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type ChatRoom struct {
	Id            int        `json:"id" gorm:"column:id;"`
	UserIds       *ChatUsers `json:"user_ids" gorm:"user_ids,type:json"`
	Type          string     `json:"type,omitempty" gorm:"type,omitempty;"`
	ChatInitiator string     `json:"chat_initiator" gorm:"column:chat_initiator;"`
}

type ChatRoomRequest struct {
	Id            string `json:"id" gorm:"-"`
	UserIds       string `json:"user_ids" gorm:"user_ids"`
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

//type UserChat struct {
//	Id       string `json:"id"`
//	FullName string `json:"full_name"`
//}
//
//func (j *UserChat) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
//	}
//
//	var img UserChat
//	if err := json.Unmarshal(bytes, &img); err != nil {
//		return err
//	}
//
//	*j = img
//	return nil
//}

//// Value return json value, implement driver.Valuer interface
//func (j *UserChat) Value() (driver.Value, error) {
//	if j == nil {
//		return nil, nil
//	}
//	return json.Marshal(j)
//}

type UserId int

func (j *UserId) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img UserId
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *UserId) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type ChatUsers []int

func (j *ChatUsers) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img ChatUsers
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *ChatUsers) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
