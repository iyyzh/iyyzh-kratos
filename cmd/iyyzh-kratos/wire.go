//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"iyyzh-kratos/internal/biz"
	"iyyzh-kratos/internal/conf"
	"iyyzh-kratos/internal/data"
	"iyyzh-kratos/internal/server"
	"iyyzh-kratos/internal/service"
)

// wire.go Provider 服务提供者， wire_gen.go Injectors 服务使用者

//使用 Makefile 编写的命令 make wire 命令生成编译期依赖注入代码 wire_gen.go
//wireApp 初始化模块 依赖注入的入口
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	// 构建所有模块中的 ProviderSet，用于生成 wire_gen.go 自动依赖注入文件
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
