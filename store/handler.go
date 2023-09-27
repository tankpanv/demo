package main

import (
	"context"

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

	pageSize := int32(1)
	cu := int32(10)
	recommendReq.PageParams =&base.PageParams{PageSize: &pageSize,Current: &cu}
	recommendResp ,err := client.Recommomend.Recommend(ctx, recommendReq)
	klog.CtxInfof(ctx, "recommendResp %v err=%v",recommendResp,err)

	packerReq := packer.NewArticleDatasReq()
	appId := int16(1)
	packerReq.AppId = &appId

	packerResp,err := client.Packer.ArticleDatas(ctx, packerReq)
	klog.CtxInfof(ctx, "packer resp %v error=%v",packerResp,err)

	return
}
