package controller

import(
    "../xingyun"
	"fmt"
)

func IndexHandler(ctx *xingyun.Context){
    ctx.Render("index",map[string]interface{}{})
}

func DetailHandler(ctx *xingyun.Context){
	id,ok:=ctx.Params["id"]
	if !ok{
		ctx.NotFound()
	}
	fmt.Println(id)
}