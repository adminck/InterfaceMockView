package router

import (
	API "InterfaceMockView/api"
	"InterfaceMockView/utils/common"
	"github.com/gin-gonic/gin"
)

func InitApiInfoRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("ApiInfo").Use(common.JWTAuth())
	{
		BaseRouter.GET("GetData", API.GetApiInfoData)
		BaseRouter.POST("Insert", API.InsertApiInfoData)
	}
	return BaseRouter
}
