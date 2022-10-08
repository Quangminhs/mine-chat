package ginchatroom

import (
	"github.com/gin-gonic/gin"
	"mine-chat/common"
	appctx "mine-chat/component"
	roombiz "mine-chat/module/room/biz"
	roommodel "mine-chat/module/room/model"
	roomstorage "mine-chat/module/room/storage"
	"net/http"
)

func InitiateChat(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var data roommodel.ChatRoomCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		store := roomstorage.NewSqlStore(db)
		biz := roombiz.NewChatRoom(store)
		err := biz.InitiateChat(c, &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
