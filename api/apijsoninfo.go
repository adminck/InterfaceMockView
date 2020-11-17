package API

import (
	"InterfaceMockView/models"
	"InterfaceMockView/utils/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetApiJsonInfoData(c *gin.Context) {
	var ApiJsonInfos []models.ApiJsonInfo
	if err := models.DB.Find(&ApiJsonInfos).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("数据获取错误%v", err), c)
	} else {
		common.GinOkWithData(ApiJsonInfos, c)
	}
}

func InsertApiJsonInfoData(c *gin.Context) {
	var ApiJsonInfo models.ApiJsonInfo
	_ = c.ShouldBindJSON(&ApiJsonInfo)
	ApiJsonInfoVerify := common.Rules{
		"ApiID":        {common.NotEmpty()},
		"Parameter":    {common.NotEmpty()},
		"JsonFilePath": {common.NotEmpty()},
	}
	ApiJsonInfoVerifyErr := common.Verify(ApiJsonInfo, ApiJsonInfoVerify)
	if ApiJsonInfoVerifyErr != nil {
		common.GinFailWithMessage(ApiJsonInfoVerifyErr.Error(), c)
		return
	}

	if err := models.DB.Create(&ApiJsonInfo).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("添加失败%v", err), c)
	} else {
		common.GinOkWithMessage("添加成功！", c)
	}
}

func QueryApiJsonInfo(c *gin.Context, ID uint) (string, error) {
	JsonString := ""
	var ApiJsonInfos []models.ApiJsonInfo
	if err := models.DB.Where("app_id = ? AND IsOpen = ?", ID, true).Find(&ApiJsonInfos).Error; err != nil {
		return JsonString, err
	}

	for _, ApiJsonInfo := range ApiJsonInfos {
		JsonString = CheckApiParam(c, ApiJsonInfo)
		if JsonString != "" {
			return JsonString, nil
		}
	}

	return "", nil
}

func CheckApiParam(c *gin.Context, ApiJsonInfo models.ApiJsonInfo) string {
	switch ApiJsonInfo.ParamType {
	case 0:

	case 1:
	case 2:

	}

	return ""
}
