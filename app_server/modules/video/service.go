package video

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"simple-video-server/config"
	"simple-video-server/constants/video_status"
	"simple-video-server/models"
	"time"
)

type service struct {
}

var Service = &service{}

func (s *service) Add2(uid uint, add VideoAdd, url string, cover string) error {

	video := &models.Video{
		Uid:    uid,
		Title:  add.Title,
		Cover:  cover,
		Url:    url,
		Locked: false,
		Status: video_status.Auditing,
	}

	err := VideoDao.Add(video)

	return err
	//return db.MysqlDB.Model(&Video{}).Create(video).Error
}

func (s *service) Add(uid uint, add VideoAdd, file multipart.File, filename string) (string, error) {
	//创建oss client实例
	//client, err := oss.New(config.Oss.Endpoint, config.Oss.AccessKeyId, config.Oss.AccessKeySecret)
	//if err != nil {
	//	panic(err)
	//}
	//
	////获取存储空间
	//bucket, err := client.Bucket(config.Oss.BucketName)
	//if err != nil {
	//	panic(err)
	//}
	fileName := fmt.Sprintf("temp/admin_temp/%d_%d_%s", uid, time.Now().UnixNano(), filename)

	fmt.Printf("fileName: %s\n", fileName)

	//err = bucket.PutObject(fileName, file)

	//if err != nil {
	//	panic(err)
	//}
	err := upload(file, fileName)
	if err != nil {
		panic(err)
	}

	fullUrl := fmt.Sprintf("https://%s.oss-cn-shenzhen.aliyuncs.com/%s", config.Oss.BucketName, fileName)
	cover := fmt.Sprintf("%s?x-oss-process=video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast", fullUrl)

	err = Service.Add2(uid, add, fullUrl, cover)

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
