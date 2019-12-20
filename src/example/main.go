package main

import (
	"github.com/ClareChu/collection/src/example/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)


func Login (ctx iris.Context) (form controller.LoginForm) {

	login := controller.LoginForm{}
	ctx.ReadJSON(&login)
	return login
}

func Login1 (ctx iris.Context) (a string) {
	a  = ctx.URLParam("a")
	return a
}

func main()  {
	app := iris.New()

	hero := hero.Register(Login)

	app.Get("/", hero.Handler(controller.User))

	hero1 := hero.Register(Login1)
	app.Get("/user1", hero1.Handler(controller.User1))
	app.Run(iris.Addr(":8080"))
}

