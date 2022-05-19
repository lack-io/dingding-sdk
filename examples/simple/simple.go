package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lack-io/dingding-sdk"
	"github.com/lack-io/dingding-sdk/api"
)

func main() {

	cli := dingding.NewClient()

	ctx := context.TODO()
	rsp, err := cli.GetContactUser(ctx, &api.GetContactUserRequest{
		Token:   "7f6fb6a7713d35368eb58c95c90b0df4",
		UnionId: "md",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rsp)
}
