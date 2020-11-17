package API

import (
	"InterfaceMockView/models"
	"InterfaceMockView/utils/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetApiInfoData(c *gin.Context) {
	var ApiInfos []models.ApiInfo
	if err := models.DB.Find(&ApiInfos).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("数据获取错误%v", err), c)
	} else {
		common.GinOkWithData(ApiInfos, c)
	}
}

func InsertApiInfoData(c *gin.Context) {
	var ApiInfo models.ApiInfo
	_ = c.ShouldBindJSON(&ApiInfo)
	ApiInfoVerify := common.Rules{
		"Name":   {common.NotEmpty()},
		"Path":   {common.NotEmpty()},
		"Domain": {common.NotEmpty()},
		"IsOpen": {common.NotEmpty()},
	}
	ApiInfoVerifyErr := common.Verify(ApiInfo, ApiInfoVerify)
	if ApiInfoVerifyErr != nil {
		common.GinFailWithMessage(ApiInfoVerifyErr.Error(), c)
		return
	}

	if err := models.DB.Create(&ApiInfo).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("添加失败%v", err), c)
	} else {
		common.GinOkWithMessage("添加成功！", c)
	}
}

func QueryApi(ApiPath, ApiType, ApiHost string) (models.ApiInfo, error) {
	var ApiInfo models.ApiInfo
	err := models.DB.Where("Path =? AND Type = ? AND Domain = ?", ApiPath, ApiType, ApiHost).Find(&ApiInfo).Error
	return ApiInfo, err
}
