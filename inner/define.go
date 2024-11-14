package main

import (
	"easy-router/inner/blls"
	"easy-router/inner/config"
	"easy-router/inner/models"
	"errors"
	"github.com/kamioair/quick-utils/qconvert"
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
	deviceCode string
	routeBll   *blls.Route
)

// 初始化
func onInit() {
	// 配置初始化
	config.Init(service.Module)

	// 如果没生成客户端唯一码，重新生成并重置客户端
	// 然后保存到文件
	if deviceCode == "" {
		code, err := refs.getDeviceCode()
		if err != nil {
			panic(err)
		}
		err = qservice.DeviceCode.SaveToFile(code)
		if err != nil {
			panic(err)
		}
		deviceCode = code
		service.ResetClient(code)
	}

	// 业务初始化
	routeBll = blls.NewRouteBll(service.Module, deviceCode, service.SendRequest)
}

// 处理外部请求
func onReqHandler(route string, ctx qdefine.Context) (any, error) {
	switch route {
	case "Request":
		info := qconvert.ToAny[models.RouteInfo](ctx.Raw())
		return routeBll.Req(info)
	}
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
