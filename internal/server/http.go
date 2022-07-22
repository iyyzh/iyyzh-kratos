package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	v1 "iyyzh-kratos/api/realword/v1"
	"iyyzh-kratos/internal/conf"
	"iyyzh-kratos/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.RealwordService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	//swagger
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)

	//http
	v1.RegisterRealwordHTTPServer(srv, greeter)
	return srv
}
