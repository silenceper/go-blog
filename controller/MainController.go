package controller

import(
    "../xingyun"
)

var name string

func IndexHandler(ctx *xingyun.Context){
    name="wenzhenlin"
    ctx.Render("index",map[string]interface{}{})
}

