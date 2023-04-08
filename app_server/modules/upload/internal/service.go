package internal

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	upload2 "simple-video-server/app_server/modules/upload"
	"simple-video-server/config"
	"simple-video-server/constants/upload_class"
	"simple-video-server/constants/upload_type"
	"simple-video-server/core"
	"simple-video-server/pkg/business_code"
	"time"
)

type Service struct {
}

var service = &Service{}

func (s *Service) Upload(c *core.Context, data *upload2.UploadData) (string, error) {
	uid := *c.AuthUID
	file, header, err := c.Request.FormFile("file")
	// TODO: 校验文件大小、格式、是否存在
	if err != nil {
		panic(business_code.EmptyUploadFile)
	}
	fmt.Printf("filename: %s, size: %d \n", header.Filename, header.Size)
	uploadType := upload_type.NewByCode(data.Type)
	if uploadType == nil {
		//TODO:
		panic(errors.New("type不正确"))
	}

	uploadClass := upload_class.GetByCode(data.Class)
	if uploadClass == nil {
		// TODO:
		panic(errors.New("class 不正确"))
	}

	// TODO: 文件名抽离
	fileName := fmt.Sprintf("temp/admin_temp/simple-video/%s/%d/_%d_%s", uploadClass.ValueString, uid, time.Now().UnixNano(), header.Filename)

	err = upload(file, fileName)
	if err != nil {
		panic(err)
	}

	// TODO: 根据测试、正式环境返回对应域名的full url
	fullUrl := fmt.Sprintf("https://%s.oss-cn-shenzhen.aliyuncs.com/%s", config.Oss.BucketName, fileName)

	return fullUrl, err
}

func upload(file multipart.File, fileName string) error {
	client, err := oss.New(config.Oss.Endpoint, config.Oss.AccessKeyId, config.Oss.AccessKeySecret)
	if err != nil {
		panic(err)
	}

	//获取存储空间
	bucket, err := client.Bucket(config.Oss.BucketName)
	if err != nil {
		panic(err)
	}
	//fileName := fmt.Sprintf("temp/admin_temp/%d_%d_%s", uid, time.Now().UnixNano(), filename)

	fmt.Printf("fileName: %s\n", fileName)

	err = bucket.PutObject(fileName, file)

	return err
}
