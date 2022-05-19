package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lack-io/dingding-sdk"
	"github.com/lack-io/dingding-sdk/api"
)

func main() {
	//appKey, _ := os.LookupEnv("ddAppKey")
	//appSecret, _ := os.LookupEnv("ddAppSecret")

	cli := dingding.NewClient(
		dingding.WithAppKey("dingn56mdcjgydflqbie"),
		dingding.WithAppSecret("7ACSM55HhNgDaq4yQ2bHU0-M1QPDoZ8zXlxBnlfdxuy1RYT2jMNswOy6LVWGVEus"),
	)

	ctx := context.TODO()
	tok, _ := cli.GetAccessToken(ctx)
	fmt.Println(tok)

	//r, err := cli.GetAuthInfos(ctx, &api.GetAuthInfosRequest{CorpId: "ding7ab5b138eeb38cd7"})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(r)

	req := &api.ListDepartmentRequest{
		DeptId:   1,
		Language: "zh_CN",
	}
	rsp, err := cli.ListDepartment(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(rsp)
	//
	////cli.GetUserAccessToken(ctx, &api.GetUserAccessTokenRequest{
	////	ClientId:     "",
	////	ClientSecret: "",
	////	Code:         "",
	////	RefreshToken: "",
	////	GrantType:    "",
	////})
	////
	cu, _ := cli.GetContactUser(ctx, &api.GetContactUserRequest{
		Token:   "",
		UnionId: "",
	})
	//cu.
	//	cli.GetUser(ctx, &api.GetUserRequest{
	//	UserId:   "",
	//	Language: "",
	//})

	//
	//dpt, err := cli.GetDepartment(ctx, &api.GetDepartmentRequest{
	//	DeptId:   2974135,
	//	Language: "zh_CN",
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(dpt)
	//
	//users, hasMore, next, err := cli.ListUser(ctx, &api.ListUserRequest{
	//	DeptId:   rsp[0].DeptId,
	//	Size:     5,
	//	Language: "zh_CN",
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(users, hasMore, next)
}
