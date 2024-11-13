package main

import (
	"easy-router/inner/blls"
	"github.com/google/uuid"
	"github.com/kamioair/quick-utils/qservice"
)

func main() {
	// 创建微服务
	code := blls.GetDeviceCode()
	if code == "" {
		uid, _ := uuid.NewUUID()
		code = uid.String()
	}
	setting := qservice.NewSetting(DefModule, DefDesc, Version).
		BindInitFunc(onInit).
		BindReqFunc(onReqHandler).
		BindNoticeFunc(onNoticeHandler).
		SetDeviceCode(code)
	service = qservice.NewService(setting)

	// 启动微服务
	service.Run()
}
