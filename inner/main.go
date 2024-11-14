package main

import (
	"github.com/google/uuid"
	"github.com/kamioair/quick-utils/qservice"
)

func main() {
	// 创建微服务
	code, _ := qservice.DeviceCode.LoadFromFile()
	if code == "" {
		uid, _ := uuid.NewUUID()
		code = uid.String()
	} else {
		deviceCode = code
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
