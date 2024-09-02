package routes

import (
	"GBolg/handler/api_v1"
	"GBolg/handler/middleware"
	"GBolg/utils/logrus_logger"
	"GBolg/utils/viper_config"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter() {

	sAppModel := viper_config.VC.GetString("server.appModel")
	sHttpPort := viper_config.VC.GetString("server.httpPort")

	gin.SetMode(sAppModel)
	r := gin.Default()

	//全局中间件，记录每个接口的调用次数和每次的耗时
	r.Use(middleware.Metric())
	r.GET("/metrics", func(ctx *gin.Context) { //Promethus要来访问这个接口
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})

	//注意 Recover 要尽量放在第一个被加载
	//如不是的话，在recover前的中间件或路由，将不能被拦截到
	//程序的原理是：
	//1.请求进来，执行recover
	//2.程序异常，抛出panic
	//3.panic被 recover捕获，返回异常信息，并Abort,终止这次请求
	r.Use(middleware.Recover)
	r.Use(middleware.Cors())

	rv1 := r.Group("api/v1")
	{
		rv1.POST("user/login", api_v1.UserLogin)

		rv1.GET("userList", api_v1.GetUserList)
		rv1.GET("categoryList", api_v1.GetCategoryList)
		rv1.GET("articleList", api_v1.GetArticleList)
		rv1.GET("articleList/:cid", api_v1.GetArticleCategoryList)
		rv1.GET("article/:id", api_v1.GetArticleInfo)

	}
	rva1 := r.Group("api/v1")
	rva1.Use(middleware.JwtVerify)
	{
		//七牛云上传
		rva1.POST("file/upload", api_v1.Upload)
		//User模块的路由接口
		rva1.POST("user/add", api_v1.AddUser)

		rva1.PUT("user/:id", api_v1.UpdateUser)
		rva1.DELETE("user/:id", api_v1.DeleteUser)
		//Category模块的路由接口
		rva1.POST("category/add", api_v1.AddCategory)
		rva1.PUT("category/:id", api_v1.UpdateCategory)
		rva1.DELETE("category/:id", api_v1.DeleteCategory)
		//Article模块的路由接口
		rva1.POST("article/add", api_v1.AddArticle)
		rva1.PUT("article/:id", api_v1.UpdateArticle)
		rva1.DELETE("article/:id", api_v1.DeleteArticle)
	}

	logrus_logger.LogRus.Info("http run successful")

	r.Run(sHttpPort)

}
