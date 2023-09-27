
include './stock_base.thrift'
include './cell_base.thrift'
include './article_base.thrift'

// 注解说明地址：https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/usage/annotation/
// thrift --gen js Contact.thrift
namespace go wealth.stock.packer

struct ArticleDatasReq {
    1: optionali16 app_id 
    2: optionalstock_base.Source source 
    3: optionali16 version_code 
    4:optionalcell_base.PackerType  packer_type
    5:optionallist<string> keywords

}
// ArticleDatas 返回数组必须和 keyword 一一对应
struct ArticleDatasResp {
    1:optionallist<article_base.ArticleData> ArticleDatas;
}


service StorePackerService {
    ArticleDatasResp ArticleDatas(1: ArticleDatasReq request) 
}

