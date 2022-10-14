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

func ListChatMessage(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var filter modelchatmessage.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		uid, err := common.FromBase58(c.Param("roomId"))
		//id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		condition := map[string]interface{}{"room_chat_id": int(uid.GetLocalID())}

		paging.Fulfill()

		store := chatmessagestorerage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := chatmessagebiz.NewChatMessageListBiz(store)

		result, err := biz.ListChatMessageList(c.Request.Context(), condition, &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
