
include '../base/base.thrift'
include './search_base.thrift'
namespace go federation.search.strategy
struct Doc{
   1:optionalstring keyword 
   2:optionalstring title
   3:optionaldouble score
   4:optional string doc_info
}
struct StrategyResponse  {
	1:optional      list<Doc> data         
    255: optional  base.Extra Extra

}
struct StrategyRequest  {
    1:optionalstring channel
    2:optionali32 version
    3:optionalstring query
	4:optional search_base.PageParams page_params
}
service SearchStrategyService {
    StrategyResponse search(1: StrategyRequest req)
}