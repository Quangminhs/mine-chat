package modelchatmessage

type Filter struct {
	PostByUser int `json:"post_by_user" gorm:"column:post_by_user;"`
}
