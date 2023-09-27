
include './stock_base.thrift'
include './cell_base.thrift'

// 注解说明地址：https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/usage/annotation/
// thrift --gen js Contact.thrift
namespace go api.storeapi

struct ItemDetailReq {
    1: optionali16 AppID  (api.query="app_id" ,api.form="app_id")
    2: optionalstock_base.Source Source  (api.query="source" ,api.form="source")
    // 3: optionalbase.TabType  TabType (api.query="tab_type" ,api.form="tab_type")
    3: optionali16 VersionCode (api.query="version_code" ,api.form="version_code")
}
struct MainStoreTabReq {
    1: optionali16 AppID  (api.query="app_id" ,api.form="app_id")
    2: optionalstock_base.Source Source  (api.query="source" ,api.form="source")
    // 3: optionalbase.TabType  TabType (api.query="tab_type" ,api.form="tab_type")
    3: optionali16 VersionCode (api.query="version_code" ,api.form="version_code")
}

struct ItemDetailResp {
    1:optionalcell_base.CellView CellView;
}
struct TabInfo{
    1:optionalcell_base.MainTabType main_tab_type // 校验MainPageResp 的MainTabType 和TabInfo 的TabInfo 一致才展示
    2:optionallist<cell_base.CellView> cell_views;
}
struct MainStoreTabResp {
    1:optionallist<TabInfo> tab_infos;
    2:optionalcell_base.MainTabType main_tab_type // 表示当前请求的tab，MainTabType
}


service StoreApiService {
    ItemDetailResp StoreDetail(1: ItemDetailReq request) (api.get="/main/store/detail");
    MainStoreTabResp MainStoreTab(1: MainStoreTabReq request) (api.get="/main/store/list");
}

