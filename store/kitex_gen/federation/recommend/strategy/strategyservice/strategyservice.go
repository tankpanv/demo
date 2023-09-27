// Code generated by Kitex v0.6.1. DO NOT EDIT.

package strategyservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	strategy "github.com/tankpanv/demo/store/kitex_gen/federation/recommend/strategy"
)

func serviceInfo() *kitex.ServiceInfo {
	return strategyServiceServiceInfo
}

var strategyServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "StrategyService"
	handlerType := (*strategy.StrategyService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Recommend": kitex.NewMethodInfo(recommendHandler, newStrategyServiceRecommendArgs, newStrategyServiceRecommendResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "strategy",
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

func recommendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*strategy.StrategyServiceRecommendArgs)
	realResult := result.(*strategy.StrategyServiceRecommendResult)
	success, err := handler.(strategy.StrategyService).Recommend(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStrategyServiceRecommendArgs() interface{} {
	return strategy.NewStrategyServiceRecommendArgs()
}

func newStrategyServiceRecommendResult() interface{} {
	return strategy.NewStrategyServiceRecommendResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Recommend(ctx context.Context, req *strategy.StrategyRequest) (r *strategy.StrategyResponse, err error) {
	var _args strategy.StrategyServiceRecommendArgs
	_args.Req = req
	var _result strategy.StrategyServiceRecommendResult
	if err = p.c.Call(ctx, "Recommend", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}