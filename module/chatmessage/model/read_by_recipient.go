package modelchatmessage

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type ReadByRecipient struct {
	Id           int        `json:"id" gorm:"column:id;"`
	ReadByUserId int        `json:"read_by_user_id" gorm:"column:read_by_user_id;"`
	ReadAt       *time.Time `json:"read_at,omitempty" gorm:"column:read_at;"`
}

type ReadByRecipients []ReadByRecipient

func (j *ReadByRecipient) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img ReadByRecipient
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *ReadByRecipient) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *ReadByRecipients) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img []ReadByRecipient
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *ReadByRecipients) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
