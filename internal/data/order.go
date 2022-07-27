package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"iyyzh-kratos/internal/biz"
)

type OrderRepo struct {
	data *Data
	log  *log.Helper
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &OrderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
