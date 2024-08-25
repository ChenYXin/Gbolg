package routes

import (
	"GBolg/handler/api_v1"
	"GBolg/handler/middleware"
	"GBolg/utils/logrus_logger"
	"GBolg/utils/viper_config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {

	sAppModel := viper_config.VC.GetString("server.appModel")
	sHttpPort := viper_config.VC.GetString("server.httpPort")

	gin.SetMode(sAppModel)
	r := gin.Default()

	//注意 Recover 要尽量放在第一个被加载
	//如不是的话，在recover前的中间件或路由，将不能被拦截到
	//程序的原理是：
	//1.请求进来，执行recover
	//2.程序异常，抛出panic
	//3.panic被 recover捕获，返回异常信息，并Abort,终止这次请求
	r.Use(middleware.Recover)
	r.Use(middleware.Cors())

	routerV1 := r.Group("api/v1")
	{
		routerV1.GET("userList", api_v1.GetUserList)
		routerV1.GET("categoryList", api_v1.GetCategoryList)
		routerV1.GET("articleList", api_v1.GetArticleList)
		routerV1.GET("articleList/:cid", api_v1.GetArticleCategoryList)
		routerV1.GET("article/:id", api_v1.GetArticleInfo)

		routerV1.POST("user/add", api_v1.AddUser)
		routerV1.Use(middleware.JwtVerify)
		//User模块的路由接口
		routerV1.PUT("user/:id", api_v1.UpdateUser)
		routerV1.DELETE("user/:id", api_v1.DeleteUser)
		//Category模块的路由接口
		routerV1.POST("category/add", api_v1.AddCategory)
		routerV1.PUT("category/:id", api_v1.UpdateCategory)
		routerV1.DELETE("category/:id", api_v1.DeleteCategory)
		//Article模块的路由接口
		routerV1.POST("article/add", api_v1.AddArticle)
		routerV1.PUT("article/:id", api_v1.UpdateArticle)
		routerV1.DELETE("article/:id", api_v1.DeleteArticle)

	}

	logrus_logger.LogRus.Info("http run successful")

	r.Run(sHttpPort)

}
