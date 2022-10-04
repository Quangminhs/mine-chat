package ginuser

import (
	"github.com/gin-gonic/gin"
	"mine-chat/common"
	"mine-chat/common/hasher"
	appctx "mine-chat/component"
	bizuser "mine-chat/module/user/biz"
	usermodel "mine-chat/module/user/model"
	userstorage "mine-chat/module/user/storage"
	"net/http"
)

func Register(appctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appctx.GetMainDBConnection()

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := bizuser.NewRegisterBiz(store, md5)

		if err := biz.RegisterUser(c, &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
