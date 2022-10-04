package ginuser

import (
	"github.com/gin-gonic/gin"
	"mine-chat/common"
	"mine-chat/common/hasher"
	"mine-chat/common/tokenprovider/jwt"
	appctx "mine-chat/component"
	bizuser "mine-chat/module/user/biz"
	usermodel "mine-chat/module/user/model"
	userstorage "mine-chat/module/user/storage"
	"net/http"
)

func Login(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := bizuser.NewLoginBiz(store, md5, tokenProvider, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
