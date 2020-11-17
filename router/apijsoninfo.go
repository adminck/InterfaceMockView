package router

import (
	API "InterfaceMockView/api"
	"InterfaceMockView/utils/common"
	"github.com/gin-gonic/gin"
)

func InitApiJsonInfoRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("ApiJsonInfo").Use(common.JWTAuth())
	{
		BaseRouter.GET("GetData", API.GetApiJsonInfoData)
		BaseRouter.POST("Insert", API.InsertApiJsonInfoData)
	}
	return BaseRouter
}