package admin

import (
	"fmt"
)

import(
	"../../xingyun"
)

func PostListHandler(ctx *xingyun.Context){
	fmt.Println(ctx.GetSession("isLogin"))
}

func PostNewHandler(ctx *xingyun.Context){
	
}

func PostDoNewHandler(ctx *xingyun.Context){
	
}

func PostUpdateHandler(ctx *xingyun.Context){
	
}

func PostDoUpdateHandler(ctx *xingyun.Context){
	
}

