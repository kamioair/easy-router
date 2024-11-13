package blls

import (
	"fmt"
	"github.com/kamioair/quick-utils/qio"
	"runtime"
)

type DeviceCode struct {
}

func NewDeviceCode() *DeviceCode {
	return &DeviceCode{}
}

func (client *DeviceCode) NewCode(refGetCodeFunc func() (string, error)) (string, error) {
	// 向上级客户端请求一个新的设备码
	code, err := refGetCodeFunc()
	if err != nil {
		return "", err
	}
	// 写入文件
	file := getFile()
	if file == "" {
		return "", err
	}
	err = qio.WriteString(file, code, false)
	if err != nil {
		return "", err
	}
	return code, nil
}

func GetDeviceCode() string {
	file := getFile()
	if qio.PathExists(file) {
		code, _ := qio.ReadAllString(file)
		return code
	}
	return ""
}

func getFile() string {
	root := qio.GetCurrentRoot()
	switch runtime.GOOS {
	case "windows":
		return fmt.Sprintf("%s\\Program Files\\qf\\device.yaml", root)
	case "linux":
		return "/dev/qf/device.yaml"
	}
	return ""
}
