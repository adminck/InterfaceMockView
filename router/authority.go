package router

import (
	API "InterfaceMockView/api"
	"github.com/gin-gonic/gin"
	"net/http"
)


func InitAuthorityLoginRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("user")
	{
		BaseRouter.POST("login", API.Login)
		BaseRouter.POST("changePassword", API.UserUpdate)
		BaseRouter.POST("register", API.UserCreate)
	}
	return BaseRouter
}

func InitHomeRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("")
	{
		ApiRouter.GET("", Home) // 创建Api
		ApiRouter.GET("login", Home) // 创建Api
		ApiRouter.GET("register", Home) // 创建Api
	}
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
