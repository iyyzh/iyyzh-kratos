package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"iyyzh-kratos/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewRealwordRepo .
func NewRealwordRepo(data *Data, logger log.Logger) biz.RealwordRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Realword) (*biz.Realword, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Realword) (*biz.Realword, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Realword, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Realword, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Realword, error) {
	return nil, nil
}
