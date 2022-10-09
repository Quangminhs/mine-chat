package common

const (
	SHARD_ID        = 1
	DB_TYPE_ROOM    = 2
	DB_TYPE_MESSAGE = 3
)

func GenUID(id int, dpType int) string {
	uid := NewUID(uint32(id), dpType, SHARD_ID)
	return uid.String()
}
