package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "iyyzh-kratos/api/realword/v1"
)

type Order struct {
	Hello string
}

type OrderRepo interface {
}

type OrderUseCase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUsecase(repo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (s *OrderUseCase) GetOrder(ctx context.Context, in *v1.GetOrderRequest) (*v1.OrderReply, error) {
	return nil, nil
}
