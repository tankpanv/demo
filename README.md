```环境：
k8s 1.26.6
store、recommend、packer 3个部署在k8s 的kitex 服务。deployment、destinationRule、virtualService 的部署文件在deployments 文件下
使用 xDS 协议进而以 Proxyless 模式运行 
istio-1.18.2

设置路由的key为 x-env=ppe

入口是服务store,store 会分别调用recommend和packer服务
问题：
1、在第一次对store服务发起请求，服务store请求完recommend 接着请求packer服务。此次请求均是正常的。
2、接着第二次再对store服务发起请求，请求recommend服务的时候会报错
err="no matched route for service federation-demo-recommend.default.svc.cluster.local:8888, err=get route failed: [XDS] manager, fetch RouteConfig resource[federation-demo-recommend.default.svc.cluster.local:8888] timeout"
此时接着请求packer服务并不会报错。
3、接着第三次再往后的请求会复现第二步的结果
4、等待几分钟后，再重新发起一次请求，会恢复到第一步。间隔时间大概几分钟
复现路径：
1、在store/testApi 执行目录下
step 1 执行 go main.go发起一个store请求。在log/app/run.log 可以看到 服务store请求完recommend 接着请求packer服务也是正常。
step 2 执行go main.go第二次请求，会看到log下报错：
no matched route for service federation-demo-recommend.default.svc.cluster.local:8888, err=get route failed: [XDS] manager, fetch RouteConfig resource[federation-demo-recommend.default.svc.cluster.local:8888] timeout
复现demo链接：https://github.com/tankpanv/demo

```