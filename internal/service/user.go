package service

import (
	"context"
	v1 "iyyzh-kratos/api/realword/v1"
)

func (s *RealwordService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
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

func (s *RealwordService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	return &v1.CreateUserReply{}, nil
}

func (s *RealwordService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	return &v1.UpdateUserReply{}, nil
}

func (s *RealwordService) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	return &v1.DeleteUserReply{}, nil
}

func (s *RealwordService) GetUserList(ctx context.Context, in *v1.GetUserListRequest) (*v1.GetUserListReply, error) {
	return &v1.GetUserListReply{}, nil
}
