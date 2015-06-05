package admin

import (
	"strings"
)

import (
	"../../xingyun"
)

func LoginHandler(ctx *xingyun.Context) {
	ctx.Config.Layout = ""
	noticeStr := ctx.GetFlash().Notice
	ctx.Render("admin/login", map[string]interface{}{
		"noticeStr": noticeStr,
	})
}


func DoLoginHandler(ctx *xingyun.Context) {
	name := ctx.Params["name"]
	password := ctx.Params["password"]

	name = strings.TrimSpace(name)
	password = strings.TrimSpace(password)

	if name == "" || password == "" {
		ctx.SetFlashNotice("用户名或密码不能为空")
		ctx.Redirect("/admin/login")
	}

	//简单验证
	if name == "admin" && password == "123" {
		ctx.SetSession("isLogin", []byte("yes"))
		ctx.Redirect("/admin/post/new")
	} else {
		ctx.SetFlashNotice("用户名或密码错误")
		ctx.Redirect("/admin/login")
	}
}
