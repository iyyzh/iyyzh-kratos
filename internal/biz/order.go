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

func NewOrderUseCase(repo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (s *OrderUseCase) GetOrder(ctx context.Context, in *v1.GetOrderRequest) (*v1.GetOrderReply, error) {
	return nil, nil
}

func (s *OrderUseCase) CreateOrder(ctx context.Context, in *v1.CreateOrderRequest) (*v1.CreateOrderReply, error) {
	return nil, nil
}

func (s *OrderUseCase) UpdateOrder(ctx context.Context, in *v1.UpdateOrderRequest) (*v1.UpdateOrderReply, error) {
	return nil, nil
}

func (s *OrderUseCase) DeleteOrder(ctx context.Context, in *v1.DeleteOrderRequest) (*v1.DeleteOrderReply, error) {
	return nil, nil
}
