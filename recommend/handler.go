package main

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	strategy "github.com/tankpanv/demo/recommend/kitex_gen/federation/recommend/strategy"
)

// StrategyServiceImpl implements the last service interface defined in the IDL.
type StrategyServiceImpl struct{}

// Recommend implements the StrategyServiceImpl interface.
func (s *StrategyServiceImpl) Recommend(ctx context.Context, req *strategy.StrategyRequest) (resp *strategy.StrategyResponse, err error) {
	// TODO: Your code here...
	klog.CtxInfof(ctx, "StrategyServiceImpl")
	resp = strategy.NewStrategyResponse()
	d := strategy.NewDoc()
	keyword := "212"
	d.Keyword = &keyword
	resp.Data = append(resp.Data, d)
	time.Sleep(2* time.Second)
	return
}
