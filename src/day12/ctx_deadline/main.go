package main

import (
	"fmt"
	"time"
	"context"
	)

func main(){
	d := time.Now().Add(20000000*time.Microsecond)
	ctx,cancel := context.WithDeadline(context.Background(),d)
	defer  cancel()
	select{
	case <-time.After(time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
