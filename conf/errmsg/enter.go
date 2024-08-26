package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//用户模块的错误
	ErrorUserNameIsExist        = 1001
	ErrorUserIsExist            = 1002
	ErrorUserNotExist           = 1003
	ErrorPassword               = 1004
	ErrorUserIllegalPermissions = 1005

	//token相关的错误
	ErrorTokenSigningFail = 4001
	ErrorTokenNotExist    = 4002
	ErrorTokenParseFail   = 4003
	ErrorTokenValidFail   = 4004

	//errmsg =2000... 分类模块的错误
	ERROR_CATEGORY_USED = 2001

	//errmsg =3000... 文章模块的错误
	ErrorArticleUsed                 = 3001
	ErrorArticleInfoNotFound         = 3002 //没找到文章信息
	ErrorArticleListNotFound         = 3003 //没找到文章列表
	ErrorArticleCategoryListNotFound = 3004 //没找到文章分类列表
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ErrorUserNameIsExist:        "用户名已存在",
	ErrorUserIsExist:            "用户已存在",
	ErrorPassword:               "密码错误",
	ErrorUserNotExist:           "用户不存在",
	ErrorUserIllegalPermissions: "用户权限不足",

	ErrorTokenSigningFail: "token生成失败",
	ErrorTokenNotExist:    "token不存在",
	ErrorTokenParseFail:   "token解析失败",

	ERROR_CATEGORY_USED: "该分类已被使用",

	ErrorArticleUsed:                 "该文章已被使用",
	ErrorArticleInfoNotFound:         "文章不存在",
	ErrorArticleListNotFound:         "没找到文章列表",
	ErrorArticleCategoryListNotFound: "没找到文章分类列表",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
