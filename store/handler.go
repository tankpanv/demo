package main

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/tankpanv/demo/store/client"
	"github.com/tankpanv/demo/store/kitex_gen/federation/recommend/base"
	"github.com/tankpanv/demo/store/kitex_gen/federation/recommend/strategy"
	mainstore "github.com/tankpanv/demo/store/kitex_gen/wealth/stock/mainstore"
	"github.com/tankpanv/demo/store/kitex_gen/wealth/stock/packer"
)

// MainstoreServiceImpl implements the last service interface defined in the IDL.
type MainstoreServiceImpl struct{}

// ItemDetail implements the MainstoreServiceImpl interface.
func (s *MainstoreServiceImpl) ItemDetail(ctx context.Context, req *mainstore.ItemDetailReq) (resp *mainstore.ItemDetailResp, err error) {
	// TODO: Your code here...
	return
}

// MainPage implements the MainstoreServiceImpl interface.
func (s *MainstoreServiceImpl) MainPage(ctx context.Context, req *mainstore.MainPageReq) (resp *mainstore.MainPageResp, err error) {
	// TODO: Your code here...
	recommendReq := strategy.NewStrategyRequest()
	channel := ""
	recommendReq.Channel = &channel
	start := time.Now()
	pageSize := int32(1)
	cu := int32(10)
	recommendReq.PageParams = &base.PageParams{PageSize: &pageSize, Current: &cu}
	recommendResp, err := client.Recommomend.Recommend(ctx, recommendReq)
	end := time.Since(start)
	klog.CtxInfof(ctx, "recommendResp %v err=%v time=%v s %v ms", recommendResp, err, end.Seconds(), end.Milliseconds())
	klog.CtxInfof(ctx, "recommend time=%v ms", end.Milliseconds())
	start = time.Now()
	packerReq := packer.NewArticleDatasReq()
	appId := int16(1)
	packerReq.AppId = &appId

	packerResp, err := client.Packer.ArticleDatas(ctx, packerReq)
	end = time.Since(start)
	klog.CtxInfof(ctx, "packer time=%v ms", end.Milliseconds())
	klog.CtxInfof(ctx, "packer resp %v error=%v time=%v s %v ms", packerResp, err, end.Seconds(), end.Milliseconds())

	return
}
