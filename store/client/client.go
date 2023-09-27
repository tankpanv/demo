package client

import (
	"context"
	"fmt"
	"time"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	xds2 "github.com/cloudwego/kitex/pkg/xds"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/xds"
	"github.com/kitex-contrib/xds/xdssuite"
	"github.com/tankpanv/demo/store/kitex_gen/federation/recommend/strategy/strategyservice"
	"github.com/tankpanv/demo/store/kitex_gen/wealth/stock/packer/storepackerservice"
)

var PassHeaderKeys = []string{
	"x-env",
	"x-env-type",
}

const (
	Namespace = "default"

	Suffix            = "svc.cluster.local"
	ClientServicePort = "8888"
	ServerServicePort = "80"

	POD_NAMESPACE_KEY = "POD_NAMESPACE"
)

func ServiceName(ServerSvc string) string {
	return fmt.Sprintf("%s.%s.%s:%s", ServerSvc, Namespace, Suffix, ClientServicePort)
}
func PassHeaderFunc(ctx context.Context) map[string]string {
	r := make(map[string]string)
	for idx := range PassHeaderKeys {
		key := PassHeaderKeys[idx]
		if v, ok := metainfo.GetPersistentValue(ctx, key); ok {
			r[key] = v
			// logs.CtxInfof(ctx, "PassHeaderKeys key:%v v:%v", key, v)
		}
	}

	return r
}

var Recommomend strategyservice.Client
var Packer storepackerservice.Client

func init() {
	var err error
	err = xds.Init()
	if err != nil {
		panic(err)

	}

	Recommomend, err = strategyservice.NewClient(
		ServiceName("federation-demo-recommend"),
		client.WithXDSSuite(xds2.ClientSuite{
			RouterMiddleware: xdssuite.NewXDSRouterMiddleware(
				xdssuite.WithRouterMetaExtractor(PassHeaderFunc),
			),
			Resolver: xdssuite.NewXDSResolver(),
		}),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithSuite(tracing.NewClientSuite()),
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "store"}),
		// client.WithRPCTimeout(1*time.Second),
		client.WithRPCTimeout(200*time.Millisecond),
	)
	if err != nil {
		panic(fmt.Sprintf("init CaifustoreClient fail %v", err))
	}

	Packer, err = storepackerservice.NewClient(
		ServiceName("federation-demo-packer"),
		client.WithXDSSuite(xds2.ClientSuite{
			RouterMiddleware: xdssuite.NewXDSRouterMiddleware(
				xdssuite.WithRouterMetaExtractor(PassHeaderFunc),
			),
			Resolver: xdssuite.NewXDSResolver(),
		}),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithSuite(tracing.NewClientSuite()),
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "store"}),
		
	)

	if err != nil {
		panic(fmt.Sprintf("init CaifustoreClient fail %v", err))
	}

	if err != nil {
		panic(fmt.Sprintf("init CaifustoreClient fail %v", err))
	}

}
