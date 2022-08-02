package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "iyyzh-kratos/api/realword/v1"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserPasswordError = errors.New("user password error")
)

// User is 定义DO 从现实世界中抽象出来的业务实体
type User struct {
	GroupId  int64
	Username string
	Account  string
	Password string
	Sex      string
}

// UserRepo is 定义操作数据库的 repo 接口 由data层实现具体数据库操作的业务逻辑
type UserRepo interface {
	Save(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	VerifyPassword(ctx context.Context, u *User) (bool, error)
}

// UserUseCase is 和 service 层进行交互 把封装好的内容传递给 service
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUseCase is 构造函数 注入与 data 层交互的 repo 接口
func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

//以下是具体业务封装代码

func (s *UserUseCase) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	vp, err := s.repo.VerifyPassword(ctx, &User{
		Account:  in.Account,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	if !vp {
		return nil, ErrUserNotFound
	}
	return &v1.LoginReply{
		Token: fmt.Sprintf("Token %s%s ", in.Account, in.Password),
	}, nil
}

func (s *UserUseCase) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	return &v1.GetUserReply{
		User: &v1.User{
			Id:       in.Id,
			GroupId:  1,
			Username: "ymy",
			Account:  "ymy123",
			Password: "123456",
			Sex:      "女",
		},
	}, nil
}

func (s *UserUseCase) CreateUser(ctx context.Context, u *User) (*User, error) {
	user, err := s.repo.Save(ctx, u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserUseCase) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	return &v1.UpdateUserReply{}, nil
}

func (s *UserUseCase) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	return &v1.DeleteUserReply{}, nil
}
