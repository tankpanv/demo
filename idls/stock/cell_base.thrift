include './article_base.thrift'
namespace go wealth.stock.cell_base
enum ShowType{
    ArticleDetail = 1
}
enum HanderType{
    ArticleDetail = 1
}
enum PackerType {
    Recommend = 0
}
struct CellView {
    1:optionalShowType show_type
    2:optionalHanderType handler_type 
    3:optional  article_base.ArticleData article_data
    4:optional  list<CellView> sub_cell_view
}
enum MainTabType {
    RecommendTab = 0
}

