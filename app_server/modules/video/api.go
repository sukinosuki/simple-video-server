package video

import (
	"fmt"
	"go.uber.org/zap"
	"simple-video-server/core"
	"simple-video-server/pkg/app_ctx"
	"simple-video-server/pkg/business_code"
	"simple-video-server/pkg/log"
)

type controller struct {
}

var Controller = &controller{}

func (ctl *controller) Add(c *core.Context) (*string, error) {
	uid, _ := app_ctx.GetUid(c.Context)
	fmt.Println("uid ", uid)

	var form VideoAdd

	err := c.ShouldBind(&form)
	if err != nil {
		panic(business_code.RequestErr)
	}

	file, header, err := c.Request.FormFile("file")

	// TODO: 校验文件大小、格式、是否存在
	if err != nil {
		panic(err)
	}

	fmt.Printf("filename: %s, size: %d \n", header.Filename, header.Size)

	if err != nil {
		panic(err)
	}

	url, err := Service.Add(*uid, form, file, header.Filename)

	if err != nil {
		panic(err)
	}

	return &url, nil
}

func (ctl *controller) GetById(c *core.Context) (string, error) {
	log := log.GetCtx(c.Request.Context())
	id := c.Param("id")

	log.Info("", zap.String("video id ", id))

	return "hanami", nil
}