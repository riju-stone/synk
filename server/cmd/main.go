package main

import (
	"context"
	"fmt"
	"time"

	"github.com/riju-stone/synk/api/service"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := service.QueryServerList(ctx, false, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
