package API

import (
	"InterfaceMockView/models"
	"InterfaceMockView/utils/common"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetApiJsonInfoData(c *gin.Context) {
	var ApiJsonInfos []models.ApiJsonInfo
	ApiID, _ := strconv.Atoi(c.PostForm("ApiID"))
	if err := models.DB.Where("api_id = ?", ApiID).Find(&ApiJsonInfos).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("数据获取错误%v", err), c)
	} else {
		common.GinOkWithData(ApiJsonInfos, c)
	}
}

func InsertApiJsonInfoData(c *gin.Context) {
	var ApiJsonInfo models.ApiJsonInfo
	_ = c.ShouldBindJSON(&ApiJsonInfo)
	ApiJsonInfoVerify := common.Rules{
		"ID":    {common.NotEmpty()},
		"ApiID": {common.NotEmpty()},
	}
	ApiJsonInfoVerifyErr := common.Verify(ApiJsonInfo, ApiJsonInfoVerify)
	if ApiJsonInfoVerifyErr != nil {
		common.GinFailWithMessage(ApiJsonInfoVerifyErr.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	waitUse := claims.(*common.CustomClaims)
	ApiJsonInfo.UpdatedUser = waitUse.NickName

	if err := models.DB.Create(&ApiJsonInfo).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("添加失败%v", err), c)
	} else {
		common.GinOkWithMessage("添加成功！", c)
	}
}

func UpdateApiJsonInfoData(c *gin.Context) {
	var ApiJsonInfo models.ApiJsonInfo
	_ = c.ShouldBindJSON(&ApiJsonInfo)
	ApiJsonInfoVerify := common.Rules{
		"ID":    {common.NotEmpty()},
		"ApiID": {common.NotEmpty()},
	}
	ApiJsonInfoVerifyErr := common.Verify(ApiJsonInfo, ApiJsonInfoVerify)
	if ApiJsonInfoVerifyErr != nil {
		common.GinFailWithMessage(ApiJsonInfoVerifyErr.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	waitUse := claims.(*common.CustomClaims)
	ApiJsonInfo.UpdatedUser = waitUse.NickName

	if err := models.DB.Save(&ApiJsonInfo).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("添加失败%v", err), c)
	} else {
		common.GinOkWithMessage("添加成功！", c)
	}
}

func DeleteApiJsonInfoData(c *gin.Context) {
	var ApiJsonInfo models.ApiJsonInfo
	_ = c.ShouldBindJSON(&ApiJsonInfo)
	ApiJsonInfoVerify := common.Rules{
		"ApiID":     {common.NotEmpty()},
		"Parameter": {common.NotEmpty()},
		"ParamType": {common.NotEmpty()},
		"IsOpen":    {common.NotEmpty()},
	}
	ApiJsonInfoVerifyErr := common.Verify(ApiJsonInfo, ApiJsonInfoVerify)
	if ApiJsonInfoVerifyErr != nil {
		common.GinFailWithMessage(ApiJsonInfoVerifyErr.Error(), c)
		return
	}

	if err := models.DB.Delete(&ApiJsonInfo).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("接口配置删除失败%v", err), c)
	} else {
		common.GinOkWithMessage("接口配置删除成功！", c)
	}
}

func GetJsonData(c *gin.Context, JsonData string) {
	if IsDoZlib(c) {
		if callback := c.Query("callback"); callback != "" {
			buf, _ := common.DoZlibCompress([]byte(callback + "(" + JsonData + ")"))
			c.Writer.Write(buf)
		} else {
			buf, _ := common.DoZlibCompress([]byte(JsonData))
			c.Writer.Write(buf)
		}
	} else {
		if callback := c.Query("callback"); callback != "" {
			c.String(http.StatusOK, callback+"("+JsonData+")")
		} else {
			c.String(http.StatusOK, JsonData)
		}
	}
}

func QueryApiJsonInfo(c *gin.Context, ID uint) (info models.ApiJsonInfo, err error) {
	var ApiJsonInfos []models.ApiJsonInfo
	if err := models.DB.Where("api_id = ? AND is_open = ?", ID, true).Find(&ApiJsonInfos).Error; err != nil {
		return info, err
	}

	if len(ApiJsonInfos) <= 0 {
		return info, errors.New("CheckApiParam close")
	}

	for _, ApiJsonInfo := range ApiJsonInfos {
		if CheckApiParam(c, ApiJsonInfo) {
			return ApiJsonInfo, nil
		}
	}

	return info, errors.New("CheckApiParam fail")
}

func CheckApiParam(c *gin.Context, ApiJsonInfo models.ApiJsonInfo) bool {
	switch ApiJsonInfo.ParamType {
	case 0:
		return Param(c, ApiJsonInfo)
	case 1:
		return Raw(c, ApiJsonInfo)
	case 2:
		return FormData(c, ApiJsonInfo)
	default:
		return true
	}
}

func Param(c *gin.Context, ApiJsonInfo models.ApiJsonInfo) bool {
	isField := false
	FieldStr := strings.Split(ApiJsonInfo.Parameter, ";")
	for _, Field := range FieldStr {
		s := strings.Split(Field, "=")
		if len(s) == 2 && s[0] != "" && s[1] != "" {
			if c.Query(s[0]) == s[1] {
				isField = true
			} else {
				isField = false
				break
			}
		}
	}
	return isField
}

func FormData(c *gin.Context, ApiJsonInfo models.ApiJsonInfo) bool {
	isField := false
	FieldStr := strings.Split(ApiJsonInfo.Parameter, ";")
	for _, Field := range FieldStr {
		s := strings.Split(Field, "=")
		if len(s) == 2 && s[0] != "" && s[1] != "" {
			if c.PostForm(s[0]) == s[1] {
				isField = true
			} else {
				isField = false
				break
			}
		}
	}
	return isField
}

func Raw(c *gin.Context, ApiJsonInfo models.ApiJsonInfo) bool {
	data, _ := ioutil.ReadAll(c.Request.Body)
	if len(data) <= 0 {
		return false
	}

	map1 := make(map[string]interface{})
	map2 := make(map[string]interface{})
	if err := json.Unmarshal([]byte(ApiJsonInfo.Parameter), &map1); err != nil {
		return false
	}
	if err := json.Unmarshal(data, &map2); err != nil {
		return false
	}

	return common.CompareTwoMapInterface(map1, map2)
}

func IsDoZlib(c *gin.Context) bool {
	if c.Request.Header["Accept-Encoding"] != nil && c.Request.Header["Accept-Encoding"][0] == "zip" {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("Content-Encoding", "zip")
		return true
	}
	return false
}
