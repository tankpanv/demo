package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/tankpanv/demo/packer/kitex_gen/wealth/stock/article_base"
	packer "github.com/tankpanv/demo/packer/kitex_gen/wealth/stock/packer"
)

// StorePackerServiceImpl implements the last service interface defined in the IDL.
type StorePackerServiceImpl struct{}

// ArticleDatas implements the StorePackerServiceImpl interface.
func (s *StorePackerServiceImpl) ArticleDatas(ctx context.Context, request *packer.ArticleDatasReq) (resp *packer.ArticleDatasResp, err error) {
	// TODO: Your code here...
	klog.CtxInfof(ctx, "ArticleDatas")
	resp = packer.NewArticleDatasResp()
	r := article_base.NewArticleData()
	author := "tttt"
	r.Author = &author
	resp.ArticleDatas = append(resp.ArticleDatas,r )
	return
}
