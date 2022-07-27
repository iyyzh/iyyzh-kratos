package service

import (
	"context"
	v1 "iyyzh-kratos/api/realword/v1"
)

func (s *RealwordService) GetOrder(ctx context.Context, in *v1.GetOrderRequest) (*v1.OrderReply, error) {
	return s.oc.GetOrder(ctx, in)
}
