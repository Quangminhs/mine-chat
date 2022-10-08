package bizuser

import (
	"context"
	"mine-chat/common"
	usermodel "mine-chat/module/user/model"
)

type ListUserStore interface {
	ListUser(context context.Context) ([]usermodel.UserSimple, error)
}

type listUserBiz struct {
	store ListUserStore
}

func NewListUserBiz(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz listUserBiz) ListUserChat(ctx context.Context) ([]usermodel.UserSimple, error) {
	dt, err := biz.store.ListUser(ctx)
	if err != nil {
		return nil, err
	}

	for i := range dt {
		dt[i].GenUID(common.DbTypeUser)
	}

	return dt, err
}
