include './stock_base.thrift'
namespace go wealth.stock.article_base

struct Member {
	1:optional string avatar 
	2:optional    string name 
	3:optional     string id 
}
struct ArticleBase {
	1:optional      i64     id                        // id
	2:optional  string      content      // content 内容
	3:optional list<string>     category // 分类列表
	4:optional     list<string>    tags                 // 标签列表
	5:optional    i32  genre            // 题材
	6:optional   i32 editor_type         // 编辑器类型
	7:optional    string    title              // 文章标题
	8:optional string      abstract_info  // 摘要
	9:optional string     user_name   
	10:optional   string    reship    // 是否转载
	11:optional     string    logo     
    12:optional       i32       visibility
	13:optionalstring keyword
}
struct ArticleData {
    1:optionalArticleBase article_base
	2:optionalstring   author 
	3:optional string status   // 'normal' | 'exception' | 'active' | 'success';

	4:optionalbool is_filtered
	5:optionallist<string> filterd_reasons
	 
	  6:optional  string href  
	  7:optional  list<Member> members 
	  8:optional  string updated_at 
	  9:optional string created_at 
	  10:optional  string sub_description  
	  11:optional  string description  
	  12:optional    i32   active_user  
	  13:optional     i32   new_user   
	  14:optional     i32   star  
	  15:optional     i32  like   
	  16:optional     i32   message  
	// 提交的变量
	  17:optional string  avatar  
18:optional  string cover  
 19:optional i32   word_num  
	  
}
struct ArticleListResponse  {
	1:optional      list<ArticleData> data 
	2:optional     i64    count          
	
}
struct GetOwnArticleReq  {
	1:optionalstock_base.PageParams page_params
	2:optionalstock_base.UserInfos user_infos
}
