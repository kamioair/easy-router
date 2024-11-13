package main

type refStruct struct{}

var refs refStruct

func (ref *refStruct) getDeviceCode() (string, error) {
	ctx, err := service.SendRequest("ClientManager", "GetDeviceCode", nil)
	if err != nil {
		return "", err
	}
	return ctx.Raw().(string), nil
}
