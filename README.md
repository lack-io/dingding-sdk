# 简介

钉钉 go 版本的 sdk，支持旧版接口

# 使用

```go
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
	rsp, err := cli.ListDepartment(context.TODO(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(rsp)
}
```