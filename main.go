package main

import (
	"flag"
	c "gossutil/config"
	"gossutil/pkg/aliyun"
	"gossutil/pkg/config"
)

func init() {
	c.Setup()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载env")
	flag.Parse()
	config.InitConfig(env)
	client := aliyun.NewOSSClient()
	bucket, _ := client.Client.Bucket(config.GetString("aliyun.bucket"))
	client.UploadFloder(bucket)
}
