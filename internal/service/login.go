package service

import (
	"context"
	"fmt"
	v1 "iyyzh-kratos/api/realword/v1"
)

////  SayHello implements realword.RealwordServer.
//func (s *RealwordService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
//	g, err := s.uc.CreateRealword(ctx, &biz.Realword{Hello: in.Name})
//	if err != nil {
//		return nil, err
//	}
//	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
//}

func (s *RealwordService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return &v1.LoginReply{
		Token: fmt.Sprintf("Token %s%s ", in.Account, in.Password),
	}, nil
}
