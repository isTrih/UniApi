package cmd

import (
	"UniApi/internal/controller/sms"
	"UniApi/internal/controller/user"
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			//启动WEB服务
			s := g.Server()

			//添加规范路由-标准返回中间件
			s.Use(ghttp.MiddlewareHandlerResponse)

			//API路由注册
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					//用户相关
					user.NewV1(),
					sms.NewV1(),
				)
				group.GET("/api", func(req *ghttp.Request) {
					req.Response.WriteTpl("/api.html")
				})
			})

			//s.Group("/", func(group *ghttp.RouterGroup) {
			//	group.Middleware(ghttp.MiddlewareHandlerResponse)
			//	group.Bind(
			//		hello.NewV1(),
			//	)
			//})
			s.Run()
			return nil
		},
	}
)
