package main

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	strategy "github.com/tankpanv/demo/recommend/kitex_gen/federation/recommend/strategy/strategyservice"
)

func main() {



	serviceName := "federation-demo-recommend"
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint("192.168.1.105:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	svr := strategy.NewServer(
		new(StrategyServiceImpl),
		//基于svc进行服务注册，9000为svc的port
		// server.WithServiceAddr(&net.TCPAddr{Port: 9000}),
		// server.WithMuxTransport(),
		server.WithSuite(tracing.NewServerSuite()),

		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),

		// server.WithMetaHandler(wtrace.ServerTTHeaderHandler), // registry

	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
