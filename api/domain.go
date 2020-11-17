package API

import (
	"InterfaceMockView/models"
	"InterfaceMockView/utils/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

const DomainPath = "./data/Domain/"

func GetDomainData(c *gin.Context) {
	var domains []models.Domain
	if err := models.DB.Find(&domains).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("数据获取错误%v", err), c)
	} else {
		common.GinOkWithData(domains, c)
	}
}

func DeleteDomainData(c *gin.Context) {
	var domain models.Domain
	_ = c.ShouldBindJSON(&domain)
	domainVerify := common.Rules{
		"ID":          {common.NotEmpty()},
		"Domain":      {common.NotEmpty()},
		"CrtFilePath": {common.NotEmpty()},
		"KeyFilePath": {common.NotEmpty()},
		"IsOpen":      {common.NotEmpty()},
	}
	domainVerifyErr := common.Verify(domain, domainVerify)
	if domainVerifyErr != nil {
		common.GinFailWithMessage(domainVerifyErr.Error(), c)
		return
	}
	claims, _ := c.Get("claims")
	waitUse := claims.(*common.CustomClaims)
	domain.UpdatedUser = waitUse.NickName
	if err := models.DB.Delete(&domain).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("域名删除失败%v", err), c)
	} else {
		common.GinOkWithMessage("域名删除成功！", c)
	}
}

func InsertDomainData(c *gin.Context) {
	var domain models.Domain
	_ = c.ShouldBindJSON(&domain)
	domainVerify := common.Rules{
		"Domain":      {common.NotEmpty()},
		"CrtFilePath": {common.NotEmpty()},
		"KeyFilePath": {common.NotEmpty()},
		"IsOpen":      {common.NotEmpty()},
	}
	domainVerifyErr := common.Verify(domain, domainVerify)
	if domainVerifyErr != nil {
		common.GinFailWithMessage(domainVerifyErr.Error(), c)
		return
	}
	claims, _ := c.Get("claims")
	waitUse := claims.(*common.CustomClaims)
	domain.UpdatedUser = waitUse.NickName
	if err := models.DB.Create(&domain).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("域名添加失败%v", err), c)
	} else {
		common.GinOkWithMessage("域名添加成功！", c)
	}
}

func UpdateDomainData(c *gin.Context) {
	var domain models.Domain
	_ = c.ShouldBindJSON(&domain)
	domainVerify := common.Rules{
		"ID":          {common.NotEmpty()},
		"Domain":      {common.NotEmpty()},
		"CrtFilePath": {common.NotEmpty()},
		"KeyFilePath": {common.NotEmpty()},
		"IsOpen":      {common.NotEmpty()},
	}
	domainVerifyErr := common.Verify(domain, domainVerify)
	if domainVerifyErr != nil {
		common.GinFailWithMessage(domainVerifyErr.Error(), c)
		return
	}
	claims, _ := c.Get("claims")
	waitUse := claims.(*common.CustomClaims)
	domain.UpdatedUser = waitUse.NickName
	if err := models.DB.Save(&domain).Error; err != nil {
		common.GinFailWithMessage(fmt.Sprintf("域名修改失败%v", err), c)
	} else {
		common.GinOkWithMessage("域名修改成功！", c)
	}
}

func CreateKeyFile(c *gin.Context) {
	common.CreateDir(DomainPath)
	f, err := c.FormFile("Key")
	//错误处理
	if err != nil {
		common.GinFailWithMessage(fmt.Sprintf("文件上传失败%v", err), c)
		return
	} else {
		if !ExtValidator(f.Filename) {
			common.GinFailWithMessage(fmt.Sprintf("后缀名不符合上传要求"), c)
			return
		}
		Filename := DomainPath + f.Filename
		if err := c.SaveUploadedFile(f, Filename); err != nil {
			common.GinFailWithMessage(fmt.Sprintf("文件上传失败%v", err), c)
			return
		}
		//保存成功返回正确的Json数据
		common.GinOkWithMessage(f.Filename, c)
	}
}

func CreateCrtFile(c *gin.Context) {
	common.CreateDir(DomainPath)
	f, err := c.FormFile("Crt")
	//错误处理
	if err != nil {
		common.GinFailWithMessage(fmt.Sprintf("文件上传失败%v", err), c)
		return
	} else {
		if !ExtValidator(f.Filename) {
			common.GinFailWithMessage(fmt.Sprintf("后缀名不符合上传要求"), c)
			return
		}
		Filename := DomainPath + f.Filename
		if err := c.SaveUploadedFile(f, Filename); err != nil {
			common.GinFailWithMessage(fmt.Sprintf("文件上传失败%v", err), c)
			return
		}
		//保存成功返回正确的Json数据
		common.GinOkWithMessage(f.Filename, c)
	}
}

func ExtValidator(Filename string) bool {
	var ext string
	s := strings.Split(Filename, ".")
	if 2 == len(s) {
		ext = path.Ext(Filename)
	} else {
		for i := 0; i < len(s); i++ {
			ext = "." + s[i]
		}
	}

	AllowExtMap := map[string]bool{".crt": true, ".key": true, ".pem": true}
	if _, ok := AllowExtMap[ext]; !ok {
		return false
	}
	return true
}
