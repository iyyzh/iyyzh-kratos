package service

import (
	"github.com/google/wire"
	v1 "iyyzh-kratos/api/realword/v1"
	"iyyzh-kratos/internal/biz"
)

//类似control

// ProviderSet is 依赖注入 服务提供者
var ProviderSet = wire.NewSet(NewRealwordService)

// RealwordService is 结构体
type RealwordService struct {
	v1.UnimplementedRealwordServer
	uc *biz.RealwordUsecase
}

// NewRealwordService is 构造函数
func NewRealwordService(uc *biz.RealwordUsecase) *RealwordService {
	return &RealwordService{uc: uc}
}
