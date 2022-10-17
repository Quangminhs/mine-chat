package ginchatmessage

import (
	"github.com/gin-gonic/gin"
	"mine-chat/common"
	appctx "mine-chat/component"
	chatmessagebiz "mine-chat/module/chatmessage/biz"
	modelchatmessage "mine-chat/module/chatmessage/model"
	chatmessagerepo "mine-chat/module/chatmessage/repo"
	chatmessagestorerage "mine-chat/module/chatmessage/storerage"
	roomstorage "mine-chat/module/room/storage"
	"net/http"
)

func ListMessageRecentConversation(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		var filter modelchatmessage.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		chatMessageStore := chatmessagestorerage.NewSqlStore(appCtx.GetMainDBConnection())
		chatRoomStore := roomstorage.NewSqlStore(appCtx.GetMainDBConnection())
		repo := chatmessagerepo.NewRecentConversationRepo(chatRoomStore, chatMessageStore)
		biz := chatmessagebiz.NewChatMessageRecentConversationBiz(repo)

		result, err := biz.GetListMessageRecentConversation(c.Request.Context(), user.GetUserId(), &filter, &paging)
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
