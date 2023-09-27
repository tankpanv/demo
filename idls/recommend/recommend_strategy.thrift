
include '../base/base.thrift'
include './recommend_base.thrift'
namespace go federation.recommend.strategy
struct Doc{
   1:optional string keyword 

}
struct StrategyResponse  {
	1:optional       list<Doc> data         
    255: optional  base.Extra Extra

}
struct StrategyRequest  {
    1:optional string channel
    2:optional i32 version
	3:optional recommend_base.PageParams page_params
}
service StrategyService {
    StrategyResponse Recommend(1: StrategyRequest req)
}