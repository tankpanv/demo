// Code generated by Kitex v0.6.1. DO NOT EDIT.

package storepackerservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	packer "github.com/tankpanv/demo/store/kitex_gen/wealth/stock/packer"
)

func serviceInfo() *kitex.ServiceInfo {
	return storePackerServiceServiceInfo
}

var storePackerServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "StorePackerService"
	handlerType := (*packer.StorePackerService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ArticleDatas": kitex.NewMethodInfo(articleDatasHandler, newStorePackerServiceArticleDatasArgs, newStorePackerServiceArticleDatasResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "packer",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func articleDatasHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*packer.StorePackerServiceArticleDatasArgs)
	realResult := result.(*packer.StorePackerServiceArticleDatasResult)
	success, err := handler.(packer.StorePackerService).ArticleDatas(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStorePackerServiceArticleDatasArgs() interface{} {
	return packer.NewStorePackerServiceArticleDatasArgs()
}

func newStorePackerServiceArticleDatasResult() interface{} {
	return packer.NewStorePackerServiceArticleDatasResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ArticleDatas(ctx context.Context, request *packer.ArticleDatasReq) (r *packer.ArticleDatasResp, err error) {
	var _args packer.StorePackerServiceArticleDatasArgs
	_args.Request = request
	var _result packer.StorePackerServiceArticleDatasResult
	if err = p.c.Call(ctx, "ArticleDatas", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
