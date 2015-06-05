package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"

	"./controller"
	"./controller/admin"
	"./model"
	"./xingyun"
)

func init() {
	model.InitDb()
	model.Dbmap.TraceOn("[gorp]", log.New(os.Stdout, "go-blog:", log.Lmicroseconds))
}

func main() {
	cfg := &xingyun.Config{
		StaticHost: "/",
		ViewPath:   "view",
		Layout:     "layout",
	}
	server := xingyun.NewServer(cfg)
	logger := server.Logger()

	//pipe setting
	pipe := server.NewPipe("normal", xingyun.PipeHandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		cfg.ViewPath = "view"
		cfg.Layout = "layout"
		next(w, r)
	}))

	//route mapping
	server.Get("/", pipe.Wrap(controller.IndexHandler))
	server.Get("/archives/{id:[0-9]+}", pipe.Wrap(controller.DetailHandler))

	//验证登入权限
	AuthPipe := server.NewPipe("auth", xingyun.PipeHandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		ctx := xingyun.GetContext(r)
		ctx.Config.Layout = "admin/layout"
		loginState := ctx.GetSession("isLogin")
		//fmt.Println(loginState)
		if string(loginState) != "yes" {
			//跳转至登入页面
			ctx.Redirect("/admin/login")
			return
		}
		next(w, r)
	}))

	//admin route
	//登入
	server.Get("/admin/login", admin.LoginHandler)
	server.Post("/admin/login", admin.DoLoginHandler)

	//文章列表
	server.Get("/admin/post", AuthPipe.Wrap(admin.PostListHandler))

	//新增文章
	server.Get("/admin/post/new", AuthPipe.Wrap(admin.PostNewHandler))
	server.Post("/admin/post/new", AuthPipe.Wrap(admin.PostDoNewHandler))

	//更新文章
	server.Get("/admin/post/update/{id:[0-9]+}", AuthPipe.Wrap(admin.PostUpdateHandler))
	server.Post("/admin/post/update/{id:[0-9]+}", AuthPipe.Wrap(admin.PostDoUpdateHandler))

	err := server.ListenAndServe(":8080")
	logger.Errorf("%s", err)
}
