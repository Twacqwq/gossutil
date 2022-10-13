package aliyun

import (
	"gossutil/pkg/config"
	driver "gossutil/pkg/driver"
	"sync"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSSClient struct {
	OSSDriver driver.OSSDriver
	Client    *oss.Client
}

var once sync.Once
var ossclient *OSSClient

func NewOSSClient() *OSSClient {
	once.Do(func() {
		ossclient = &OSSClient{
			OSSDriver: &AliyunOSS{},
			Client:    newClient(),
		}
	})
	return ossclient
}

func (s *OSSClient) UploadFloder(bucket *oss.Bucket) {
	s.OSSDriver.UploadFloder(bucket)
}

func newClient() *oss.Client {
	cli, _ := oss.New(
		config.GetString("aliyun.endpoint"),
		config.GetString("aliyun.access_key_id"),
		config.GetString("aliyun.access_key_secret"),
	)
	return cli
}
