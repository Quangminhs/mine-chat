package common

type EntityModel struct {
	Id     int  `json:"-" gorm:"column:id;"`
	FakeId *UID `json:"id" gorm:"-"`
}

func (m *EntityModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)

	m.FakeId = &uid
}
