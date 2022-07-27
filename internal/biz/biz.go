package biz

import "github.com/google/wire"

//3 biz层 业务逻辑的组装层 定义操作数据库的 repo 接口，使用依赖倒置的原则

// ProviderSet is 依赖注入
var ProviderSet = wire.NewSet(NewUserUsecase, NewOrderUsecase)
