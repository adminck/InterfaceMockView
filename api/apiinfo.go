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

func DeleteApiInfoData(c *gin.Context) {
	var ApiInfo models.ApiInfo
	_ = c.ShouldBindJSON(&ApiInfo)
	ApiInfoVerify := common.Rules{
		"Name":   {common.NotEmpty()},
		"Path":   {common.NotEmpty()},
		"Domain": {common.NotEmpty()},
	}
	ApiInfoVerifyErr := common.Verify(ApiInfo, ApiInfoVerify)
	if ApiInfoVerifyErr != nil {
		common.GinFailWithMessage(ApiInfoVerifyErr.Error(), c)
		return
	}


	claims, _ := c.Get("claims")
	waitUse := claims.(*common.CustomClaims)
	ApiInfo.UpdatedUser = waitUse.NickName

	if err := models.DB.Delete(&ApiInfo).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("接口删除失败%v", err), c)
	} else {
		common.GinOkWithMessage("接口删除成功！", c)
	}
}

func UpdateApiInfoData(c *gin.Context) {
	var ApiInfo models.ApiInfo
	_ = c.ShouldBindJSON(&ApiInfo)
	ApiInfoVerify := common.Rules{
		"Name":   {common.NotEmpty()},
		"Path":   {common.NotEmpty()},
		"Domain": {common.NotEmpty()},
	}
	ApiInfoVerifyErr := common.Verify(ApiInfo, ApiInfoVerify)
	if ApiInfoVerifyErr != nil {
		common.GinFailWithMessage(ApiInfoVerifyErr.Error(), c)
		return
	}

	if err := models.DB.Save(&ApiInfo).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("接口数据更新失败%v", err), c)
	} else {
		common.GinOkWithMessage("接口数据更新成功！", c)
	}
}

func QueryApi(ApiPath, ApiHost string) (models.ApiInfo, error) {
	var ApiInfo models.ApiInfo
	err := models.DB.Where("Path =? AND Domain = ?", ApiPath, ApiHost).First(&ApiInfo).Error
	return ApiInfo, err
}
