package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	appctx "mine-chat/component"
	"mine-chat/middleware"
	"mine-chat/module/room/transport/ginchatroom"
	userstorage "mine-chat/module/user/storage"
	"mine-chat/module/user/transport/ginuser"
	"net/http"
)

func main() {
	dsn := "mine-chat:123456@tcp(127.0.0.1:3306)/mine-chat?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := os.Getenv("MYSQL_CONN_STRING")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.Debug()

	appCtx := appctx.NewAppContext(db, "1234")

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	v1 := r.Group("/v1")
	v1.POST("/login", ginuser.Login(appCtx))
	userStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	v1.GET("/profile", middleware.RequiredAuth(appCtx, userStore), ginuser.Profile(appCtx))

	users := v1.Group("/users")
	{
		users.POST("", ginuser.Register(appCtx))
		users.GET("", ginuser.ListUserChat(appCtx))
	}

	chat := v1.Group("/chat")
	{
		chat.POST("/initiate", ginchatroom.InitiateChat(appCtx))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(); err != nil {
		return
	}
}
