include '../base/base.thrift'
include './stock_base.thrift'
include './cell_base.thrift'
namespace go wealth.stock.mainstore

struct ItemDetailReq {
    1: optionali16 app_id  (api.query="app_id" ,api.form="app_id")
    2: optionalstock_base.Source source  (api.query="source" ,api.form="source")
    // 3: optionalbase.TabType  TabType (api.query="tab_type" ,api.form="tab_type")
    3: optionali16 version_code (api.query="version_code" ,api.form="version_code")
}


struct ItemDetailResp {
    1:optionalcell_base.CellView cell_view;
}

struct MainPageReq {
    1: optionali16 app_id  (api.query="app_id" ,api.form="app_id")
    2: optionalstock_base.Source source  (api.query="source" ,api.form="source")
    3: optionalcell_base.MainTabType  main_tab_type (api.query="main_tab_type" ,api.form="main_tab_type")
    4: optionali16 version_code (api.query="version_code" ,api.form="version_code")
    5: optionali32 next_offset
    6: optionali32 page_size
}

struct TabInfo{
    1:optionalcell_base.MainTabType main_tab_type // 校验MainPageResp 的MainTabType 和TabInfo 的TabInfo 一致才展示
    2:optionallist<cell_base.CellView> cell_views;
    3:optionali32 count
    4:optionali32 next_offset
    5:optionali32 has_more
}
struct MainPageResp {
    1:optionallist<TabInfo> tab_infos // 展示tab的list，下标和MainTabType 不是一一对应关系。
    2:optionalcell_base.MainTabType main_tab_type // 表示当前请求的tab，MainTabType
  
    255: optional  base.Extra extra
}

service MainstoreService {
    ItemDetailResp ItemDetail(1: ItemDetailReq req)
    MainPageResp MainPage(1: MainPageReq req)
}

