package main

import (
	"github.com/kataras/iris"
)

func main() {

	app := iris.New()
	/**
	 * 路由组请求
	 */
	userParty := app.Party("/users/{date}/", func(context context.Context) {
		app.Logger().Info("request users controller")
		// 处理下一级请求
		context.Next()
	})

	/**
	 * 路由组下面的下一级请求
	 * ../users/login
	 */
	 userParty.Get("/login", func(context context.Context) {
		app.Logger().Info("user login")
	})

	usersRouter := app.Party("/admin", userMiddleware)

	usersRouter.Get("/info", func(context context.Context) {
		app.Logger().Info("users info")
		context.Next()// 手动显示调用
	})

	//Done：
	usersRouter.Done(func(context context.Context) {
		context.Application().Logger().Infof("response sent to " + context.Path())
	})

}

//用户路由中间件
func userMiddleware(context iris.Context) {
	context.Next()
}
