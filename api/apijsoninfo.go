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
	"reflect"
	"strconv"
	"strings"
)

func GetApiJsonInfoData(c *gin.Context) {
	var ApiJsonInfos []models.ApiJsonInfo
	ApiID,_ := strconv.Atoi(c.PostForm("ApiID"))
	if err := models.DB.Where("api_id = ?",ApiID).Find(&ApiJsonInfos).Error; err != nil {
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
		"ParamType": 	{common.NotEmpty()},
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
		"ApiID":        {common.NotEmpty()},
		"Parameter":    {common.NotEmpty()},
		"JsonFilePath": {common.NotEmpty()},
		"ParamType": 	{common.NotEmpty()},
		"IsOpen": 		{common.NotEmpty()},
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
		"ApiID":        {common.NotEmpty()},
		"Parameter":    {common.NotEmpty()},
		"JsonFilePath": {common.NotEmpty()},
		"ParamType": 	{common.NotEmpty()},
		"IsOpen": 		{common.NotEmpty()},
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

func GetJsonDataC(c *gin.Context) {
	JsonFilePath := c.Query("JsonFilePath")
	JsonFilePath =  "./data/json/" + JsonFilePath
	var mapResult map[string]interface{}
	if data, err := ioutil.ReadFile(JsonFilePath); err != nil {
		common.Result(common.ERROR, gin.H{}, fmt.Sprintf("接口json文件读取失败%v", err), c)
	} else {
		json.Unmarshal(data, &mapResult)
		common.GinOkWithData(mapResult,c)
	}
}

func GetJsonData(c *gin.Context,JsonFilePath string) {
	JsonFilePath =  "./data/json/" + JsonFilePath
	if data, err := ioutil.ReadFile(JsonFilePath); err != nil {
		common.Result(common.ERROR, gin.H{}, fmt.Sprintf("接口json文件读取失败%v", err), c)
		c.Abort()
		return
	}else {
		if IsDoZlib(c) {
			if callback := c.Query("callback"); callback != "" {
				buf, _ := common.DoZlibCompress([]byte(callback + "(" + string(data) + ")"))
				c.Writer.Write(buf)
			} else {
				buf, _ := common.DoZlibCompress(data)
				c.Writer.Write(buf)
			}
		} else {
			if callback := c.Query("callback"); callback != "" {
				c.String(http.StatusOK, callback+"("+string(data)+")")
			} else {
				c.String(http.StatusOK, string(data))
			}
		}
	}
}

func QueryApiJsonInfo(c *gin.Context, ID uint) (info models.ApiJsonInfo, err error) {
	var ApiJsonInfos []models.ApiJsonInfo
	if err := models.DB.Where("api_id = ? AND is_open = ?", ID, true).Find(&ApiJsonInfos).Error; err != nil {
		return info, err
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

	map2 := make(map[string]interface{})
	if err := json.Unmarshal(data, &map2); err != nil {
		return false
	}

	isField := false
	FieldStr := strings.Split(ApiJsonInfo.Parameter, ";")
	for _, Field := range FieldStr {
		s := strings.Split(Field, "=")
		if len(s) == 2 && s[0] != "" && s[1] != "" {
			for k, v := range map2 {
				if k != s[0] {
					continue
				}
				switch s[1][0:1] {
				case "{":
					if mv, ok := v.(map[string]interface{}); ok {
						map3 := make(map[string]interface{})
						if err := json.Unmarshal([]byte(s[1]), &map3); err != nil {
							isField = false
							break
						}
						if reflect.DeepEqual(mv, map3) {
							isField = true
							break
						}
					}
					break
				case "[":
					if mv, ok := v.([]interface{}); ok {
						switch mv[0].(type) {
						case bool:
							arr := strings.Split(s[1][1:len(s[1])-1], ",")
							for index, value := range arr {
								if index <= len(mv) {
									if value == "true" && mv[index].(bool) {
										isField = true
										continue
									} else if !mv[index].(bool) {
										isField = true
										continue
									}
								}
								isField = false
								break
							}
						case string:
							arr := strings.Split(s[1][1:len(s[1])-1], ",")
							for index, value := range arr {
								if index <= len(mv) {
									if value == mv[index].(string) {
										isField = true
										continue
									}
								}
								isField = false
								break
							}
						case float64:
							arr := strings.Split(s[1][1:len(s[1])-1], ",")
							for index, value := range arr {
								if index <= len(mv) {
									if v2, err := strconv.ParseFloat(value, 64); err == nil && v2 == mv[index].(float64) {
										isField = true
										continue
									}
								}
								isField = false
								break
							}
						}
					}
					break
				default:
					if mv, ok := v.(string); ok {
						if s[1] == mv {
							isField = true
							break
						}
					}
					if mv, ok := v.(float64); ok {
						if v2, err := strconv.ParseFloat(s[1], 64); err == nil && v2 == mv {
							isField = true
							break
						}
					}
					if mv, ok := v.(bool); ok {
						if s[1] == "true" && mv {
							isField = true
							break
						} else if !mv {
							isField = true
							break
						}
					}
					break
				}
				if isField {
					break
				}
			}
		}
	}
	return isField
}


func IsDoZlib(c *gin.Context) bool {
	if c.Request.Header["Accept-Encoding"] != nil && c.Request.Header["Accept-Encoding"][0] == "zip" {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("Content-Encoding", "zip")
		return true
	}
	return false
}