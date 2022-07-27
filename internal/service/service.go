package service

import (
	"github.com/google/wire"
	v1 "iyyzh-kratos/api/realword/v1"
	"iyyzh-kratos/internal/biz"
)

//2 service层 实现了 api 定义的服务层 处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
//DTO -> DO 接口接收参数将数据传输参数 转化为 数据库实体操作数据库使用

//ProviderSet is 依赖注入
var ProviderSet = wire.NewSet(NewRealwordService)

// RealwordService is 管理与biz交互的 结构体
type RealwordService struct {
	v1.UnimplementedRealwordServer

	uc *biz.UserUseCase
	oc *biz.OrderUseCase
}

// NewRealwordService is 构造函数 注入交互biz的实例
func NewRealwordService(uc *biz.UserUseCase, oc *biz.OrderUseCase) *RealwordService {
	return &RealwordService{
		uc: uc,
		oc: oc,
	}
}
