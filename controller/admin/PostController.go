package admin

import (
	"strings"
//	"fmt"
	//"../../service"
	"../../xingyun"
)

func PostListHandler(ctx *xingyun.Context) {

}

func PostNewHandler(ctx *xingyun.Context) {
	noticeStr := ctx.GetFlash().Notice
	ctx.Render("admin/post_new", map[string]interface{}{
		"noticeStr":noticeStr,
	})
	
}

func PostDoNewHandler(ctx *xingyun.Context) {
	//err:=service.WritePost("标题测试", "内容测试", "tag1,tag2", 1433494374, 1, 1)
	title:=ctx.Params["title"]
	content:=ctx.Params["text"]
	//tags:=ctx.Params["tags"]
	//pubdate:=ctx.Params["date"]	
	
	title=strings.TrimSpace(title)
	content=strings.TrimSpace(content)
	if title=="" || content==""{
		ctx.SetFlashNotice("标题和内容必填")
		ctx.Redirect("/admin/post/new")
		return 
	}
	
}

func PostUpdateHandler(ctx *xingyun.Context) {

}

func PostDoUpdateHandler(ctx *xingyun.Context) {

}
