package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"iyyzh-kratos/internal/biz"
)

//UserRepo is 实现 biz 层定义的 repo 接口
type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo is 结构体 注入数据库连接
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

//以下是具体实现 biz 层定义的 repo 接口 访问数据库封装数据返回给 biz 层

func (r *UserRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}
