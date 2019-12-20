package controller

import (
	"fmt"
	"github.com/kataras/iris"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInterface interface {
}

func User(to LoginForm, ctx iris.Context) (login *LoginForm) {
	ctx.Path()
	fmt.Print(to)
	return &LoginForm{Username: "1", Password: "3"}
}

func User1(a string) string {
	return a
}
