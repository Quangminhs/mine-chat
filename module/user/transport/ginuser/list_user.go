package ginuser

import (
	"github.com/gin-gonic/gin"
	"mine-chat/common"
	appctx "mine-chat/component"
	bizuser "mine-chat/module/user/biz"
	userstorage "mine-chat/module/user/storage"
	"net/http"
)

func ListUserChat(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		biz := bizuser.NewListUserBiz(store)
		dt, err := biz.ListUserChat(c)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dt))
	}
}
