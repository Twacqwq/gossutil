package aliyun

import (
	"gossutil/pkg/config"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/fatih/color"
)

type AliyunOSS struct{}

func (aliyun *AliyunOSS) UploadFloder(bucket interface{}) {
	b := bucket.(*oss.Bucket)
	root := config.GetString("aliyun.upload_path")
	if _, err := os.Stat(root); err != nil {
		color.Red("directory [%s] exist? %v\n", root, err)
		return
	}
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && err == nil {
			return nil
		}
		if err != nil {
			color.Red("Bucket %s not found\n", b.BucketName)
		}
		err = b.PutObjectFromFile(path, path)
		if err != nil {
			color.Red("[UPLOAD] %s: %v\n", path, err)
			log.Fatalf("[UPLOAD]: %v", err)
			return err
		}
		color.Green("[UPLOAD] %s SUCCESS!\n", path)
		return nil
	})
}
