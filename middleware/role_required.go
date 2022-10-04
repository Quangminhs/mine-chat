package middleware

import (
	"github.com/gin-gonic/gin"
	"mine-chat/common"
	appctx "mine-chat/component"
)

func RoleRequired(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		for _, item := range allowRoles {
			if item == user.GetRole() {
				c.Next()
				return
			}
		}

		panic(common.ErrNoPermission(nil))
	}
}
