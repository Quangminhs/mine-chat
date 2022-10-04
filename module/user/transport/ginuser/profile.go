package ginuser

import (
	"github.com/gin-gonic/gin"
	"mine-chat/common"
	appctx "mine-chat/component"
	"net/http"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)
		//newPass := "kdsjkdfsjkdjfksdf"
		//type update struct {
		//	NewPass *string
		//}
		//
		//log.Println( update{ NewPass: &newPass})

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
