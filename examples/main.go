package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lack-io/dingding-sdk"
	"github.com/lack-io/dingding-sdk/api"
)

func main() {
	appKey, _ := os.LookupEnv("ddAppKey")
	appSecret, _ := os.LookupEnv("ddAppSecret")

	cli := dingding.NewClient(
		dingding.WithAppKey(appKey),
		dingding.WithAppSecret(appSecret),
	)

	req := &api.ListDepartmentRequest{
		DeptId:   1,
		Language: "zh_CN",
	}

	ctx := context.TODO()
	tok, _ := cli.GetAccessToken(ctx)
	fmt.Println(tok)

	r, err := cli.GetAuthInfos(ctx, &api.GetAuthInfosRequest{CorpId: "ding7ab5b138eeb38cd7"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r)

	rsp, err := cli.ListDepartment(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(rsp)

	users, hasMore, next, err := cli.ListUser(ctx, &api.ListUserRequest{
		DeptId:   rsp[0].DeptId,
		Size:     5,
		Language: "zh_CN",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(users, hasMore, next)
}
