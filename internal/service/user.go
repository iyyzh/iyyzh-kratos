package service

import (
	"context"
	v1 "iyyzh-kratos/api/realword/v1"
	"iyyzh-kratos/internal/biz"
)

//以下实现 DTO to DO 的简单业务逻辑

func (s *RealwordService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.UserReply, error) {
	return s.uc.GetUser(ctx, in)
}

func (s *RealwordService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.StateReply, error) {
	return s.uc.CreateUser(ctx, &biz.User{})
}

func (s *RealwordService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.StateReply, error) {
	return s.uc.UpdateUser(ctx, in)
}

func (s *RealwordService) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.StateReply, error) {
	return s.uc.DeleteUser(ctx, in)
}

func (s *RealwordService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return s.uc.Login(ctx, in)
}
