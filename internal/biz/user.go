package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "iyyzh-kratos/api/realword/v1"
)

// User is 定义操作数据库的结构体参数
type User struct {
	Hello string
}

// UserRepo is 定义操作数据库的 repo 接口 由data层实现具体数据库操作的业务逻辑
type UserRepo interface {
	Save(ctx context.Context, u *User) (*User, error)
}

// UserUseCase is 和 service 层进行交互 把封装好的内容传递给 service
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase is 构造函数 注入与 data 层交互的 repo 接口
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

//以下是具体业务封装代码

func (s *UserUseCase) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{
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

func (s *UserUseCase) CreateUser(ctx context.Context, u *User) (*v1.StateReply, error) {
	s.repo.Save(ctx, u)
	return &v1.StateReply{}, nil
}

func (s *UserUseCase) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.StateReply, error) {
	return &v1.StateReply{}, nil
}

func (s *UserUseCase) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.StateReply, error) {
	return &v1.StateReply{}, nil
}

func (s *UserUseCase) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return &v1.LoginReply{
		Token: fmt.Sprintf("Token %s%s ", in.Account, in.Password),
	}, nil
}
