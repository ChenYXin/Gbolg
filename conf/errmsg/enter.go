package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//errmsg =1000... 用户模块的错误
	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WRONG     = 1002
	ERROR_USERNAME_NOT_EXIST = 1003
	ERROR_TOKEN_NOT_EXIST    = 1004
	ERROR_TOKEN_RUNTIME      = 1005
	ERROR_TOKEN_WRONG        = 1006
	ERROR_TOKEN_TYPE_WRONG   = 1007

	//errmsg =2000... 分类模块的错误
	ERROR_CATEGORY_USED = 2001

	//errmsg =3000... 文章模块的错误
	ErrorArticleUsed                 = 3001
	ErrorArticleInfoNotFound         = 3002 //没找到文章信息
	ErrorArticleListNotFound         = 3003 //没找到文章列表
	ErrorArticleCategoryListNotFound = 3004 //没找到文章分类列表
)

var codeMsg = map[int]string{
	SUCCESS:                          "OK",
	ERROR:                            "FAIL",
	ERROR_USERNAME_USED:              "用户名已存在",
	ERROR_PASSWORD_WRONG:             "密码错误",
	ERROR_USERNAME_NOT_EXIST:         "用户不存在",
	ERROR_TOKEN_NOT_EXIST:            "token不存在",
	ERROR_TOKEN_RUNTIME:              "token已过期",
	ERROR_TOKEN_WRONG:                "token不正确",
	ERROR_TOKEN_TYPE_WRONG:           "token格式错误",
	ERROR_CATEGORY_USED:              "该分类已被使用",
	ErrorArticleUsed:                 "该文章已被使用",
	ErrorArticleInfoNotFound:         "文章不存在",
	ErrorArticleListNotFound:         "没找到文章列表",
	ErrorArticleCategoryListNotFound: "没找到文章分类列表",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
