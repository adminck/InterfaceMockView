package router

import (
	API "InterfaceMockView/api"
	"InterfaceMockView/utils/common"
	"github.com/gin-gonic/gin"
)

func InitApiInfoRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("ApiInfo").Use(common.JWTAuth())
	{
		BaseRouter.POST("GetData", API.GetApiInfoData)
		BaseRouter.POST("Insert", API.InsertApiInfoData)
		BaseRouter.POST("Update", API.UpdateApiInfoData)
		BaseRouter.POST("Delete", API.DeleteApiInfoData)
	}
	return BaseRouter
}
