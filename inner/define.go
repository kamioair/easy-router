package main

import (
	"easy-router/inner/blls"
	"easy-router/inner/config"
	"errors"
	"github.com/kamioair/quick-utils/qdefine"
	"github.com/kamioair/quick-utils/qservice"
)

const (
	Version   = "V1.0.0"   // 版本
	DefModule = "Route"    // 模块名称
	DefDesc   = "路由调度管理模块" // 模块描述
)

var (
	service *qservice.MicroService

	// 其他业务
	devCodeBll *blls.DeviceCode
)

// 初始化
func onInit() {
	// 配置初始化
	config.Init(service.Module)

	// 业务初始化
	devCodeBll = blls.NewDeviceCode()

	// 如果没生成客户端唯一码，重新生成并重置客户端
	if blls.GetDeviceCode() == "" {
		code, err := devCodeBll.NewCode(func() (string, error) {
			return refs.getDeviceCode()
		})
		if err != nil {
			panic(err)

		}
		service.ResetClient(code)
	}
}

// 处理外部请求
func onReqHandler(route string, ctx qdefine.Context) (any, error) {
	return nil, errors.New("route Not Matched")
}

// 处理外部通知
func onNoticeHandler(route string, ctx qdefine.Context) {

}

// 发送通知
func onNotice(route string, content any) {
	service.SendNotice(route, content)
}

// 发送日志
func onLog(logType qdefine.ELog, content string, err error) {
	service.SendLog(logType, content, err)
}
