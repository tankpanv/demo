namespace go wealth.stock.base
namespace js wealth.stock.base
struct PageParams {
    1:optional i32 current
	2:optionali32  pageSize
    
}
 enum Source {
    MainStore = 1 
 }
 enum TabType {
    Article = 1
 }
struct UserSecret {
	1:optionalstring username // 密码
    2:optionalstring password // 密码
}

struct UserInfos {
	1:optionali64 id
	2:optionalstring username 
	3:optionalstring show_name 
	4:optionalstring  avatar  
	5:optionallist<string>  tags
	6:optional   string email          // 邮箱
	7:optional   string phone        // 手机
	8:optional   string level       // 用户等级
}