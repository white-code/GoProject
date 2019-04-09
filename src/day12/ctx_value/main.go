package main

import (
	"context"
	"fmt"
)


func process(ctx context.Context){
	ret,_ := ctx.Value("trace_id").(int)
	s,_ := ctx.Value("session").(string)
	fmt.Printf("ret : %d\nsession : %v",ret,s)
}

//所有函数都需要用的参数，用ctx来指定
func main(){
	ctx := context.WithValue(context.Background(),"trace_id",13483434)
	ctx = context.WithValue(ctx,"session","my english is very poor")
	process(ctx)
}
