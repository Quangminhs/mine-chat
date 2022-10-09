package ginchatmessage

import (
	"github.com/gin-gonic/gin"
	"mine-chat/common"
	appctx "mine-chat/component"
	chatmessagebiz "mine-chat/module/chatmessage/biz"
	modelchatmessage "mine-chat/module/chatmessage/model"
	chatmessagestorerage "mine-chat/module/chatmessage/storerage"
	"net/http"
)

func PostChatMessage(ctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		romChatId, err := common.FromBase58(c.Param("roomId"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		var data modelchatmessage.ChatMessageCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		postByUserId, err := common.FromBase58(data.PostByUser)
		if err != nil {
			panic(err)
			return
		}

		dataChatMessage := modelchatmessage.ChatMessage{
			RomChatId:        int(romChatId.GetLocalID()),
			PostByUser:       int(postByUserId.GetLocalID()),
			Message:          data.Message,
			ReadByRecipients: nil,
		}

		db := ctx.GetMainDBConnection()
		store := chatmessagestorerage.NewSqlStore(db)
		biz := chatmessagebiz.NewPostChatMessageBiz(store)

		if err := biz.PostNewChatMessageBiz(c, &dataChatMessage); err != nil {
			panic(err)
			return
		}

		data.Id = common.GenUID(dataChatMessage.Id, common.DB_TYPE_MESSAGE)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
