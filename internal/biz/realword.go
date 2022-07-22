package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "iyyzh-kratos/api/realword/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Realword is a Realword model.
type Realword struct {
	Hello string
}

// RealwordRepo is a Greater repo.
type RealwordRepo interface {
	Save(context.Context, *Realword) (*Realword, error)
	Update(context.Context, *Realword) (*Realword, error)
	FindByID(context.Context, int64) (*Realword, error)
	ListByHello(context.Context, string) ([]*Realword, error)
	ListAll(context.Context) ([]*Realword, error)
}

// RealwordUsecase is a Realword usecase.
type RealwordUsecase struct {
	repo RealwordRepo
	log  *log.Helper
}

// NewRealwordUsecase new a Realword usecase.
func NewRealwordUsecase(repo RealwordRepo, logger log.Logger) *RealwordUsecase {
	return &RealwordUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateRealword creates a Realword, and returns the new Realword.
func (uc *RealwordUsecase) CreateRealword(ctx context.Context, g *Realword) (*Realword, error) {
	uc.log.WithContext(ctx).Infof("CreateRealword: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
