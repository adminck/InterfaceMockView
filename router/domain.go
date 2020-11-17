package router

import (
	API "InterfaceMockView/api"
	"InterfaceMockView/utils/common"
	"github.com/gin-gonic/gin"
)

func InitDomainRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("Domain").Use(common.JWTAuth())
	{
		BaseRouter.POST("GetData", API.GetDomainData)
		BaseRouter.POST("Insert", API.InsertDomainData)
		BaseRouter.POST("Update", API.UpdateDomainData)
		BaseRouter.POST("UpCrtFile", API.CreateCrtFile)
		BaseRouter.POST("UpKeyFile", API.CreateKeyFile)
		BaseRouter.POST("Delete", API.DeleteDomainData)
	}
	return BaseRouter
}
