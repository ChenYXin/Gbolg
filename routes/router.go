package routes

import (
	"GBolg/handler/api_v1"
	"GBolg/utils/logrus_logger"
	"GBolg/utils/viper_config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {

	sAppModel := viper_config.VC.GetString("server.appModel")
	sHttpPort := viper_config.VC.GetString("server.httpPort")

	gin.SetMode(sAppModel)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		//User模块的路由接口
		routerV1.POST("user/add", api_v1.AddUser)
		routerV1.GET("userList", api_v1.GetUserList)
		routerV1.PUT("user/:id", api_v1.UpdateUser)
		routerV1.DELETE("user/:id", api_v1.DeleteUser)
		//Article模块的路由接口
		routerV1.POST("category/add", api_v1.AddCategory)
		routerV1.GET("categoryList", api_v1.GetCategoryList)
		routerV1.PUT("category/:id", api_v1.UpdateCategory)
		routerV1.DELETE("category/:id", api_v1.DeleteCategory)
		//Category模块的路由接口

	}

	logrus_logger.LogRus.Info("http run successful")

	r.Run(sHttpPort)

}
