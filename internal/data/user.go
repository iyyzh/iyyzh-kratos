package data

import (
	"context"
	"gorm.io/gorm"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"iyyzh-kratos/internal/biz"
)

//UserRepo is 实现 biz 层定义的 repo 接口
type UserRepo struct {
	data *Data
	log  *log.Helper
}

//User is 定义PO 关系型数据库映射属性
type User struct {
	gorm.Model
	Id        int64
	GroupId   int64
	Username  string
	Account   string
	Password  string
	Sex       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUserRepo is 结构体 注入数据库连接
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

//以下是具体实现 biz 层定义的 repo 接口 访问数据库封装数据返回给 biz 层

func (r *UserRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	return u, nil
}

func (r *UserRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	return true, nil
}

func (r *UserRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	return &biz.User{}, nil
}
