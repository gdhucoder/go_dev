package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	res := ctx.Value("tace_id")
	fmt.Println(res)
}

func main() {
	ctx := context.WithValue(context.Background(), "tace_id", 123123123)
	process(ctx)
}
