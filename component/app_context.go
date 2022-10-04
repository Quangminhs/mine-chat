package appctx

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	//UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	//GetPubSub() pubsub.PubSub
}

type appCtx struct {
	db *gorm.DB
	//uploadProvider uploadprovider.UploadProvider
	secretKey string
	//pb             pubsub.PubSub
}

func NewAppContext(db *gorm.DB, secretKey string) *appCtx {
	return &appCtx{db: db, secretKey: secretKey}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string { return ctx.secretKey }

//func (ctx *appCtx) GetPubSub() pubsub.PubSub { return ctx.pb }
