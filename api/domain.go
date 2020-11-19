package API

import (
	"InterfaceMockView/models"
	"InterfaceMockView/utils/common"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
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
	s := c.PostForm("Domain")
	if err := json.Unmarshal([]byte(s), &domain); err != nil {
		common.GinFailWithMessage(fmt.Sprintf("数据解析失败%v", err), c)
	}
	domainVerify := common.Rules{
		"Domain":      {common.NotEmpty()},
	}
	domainVerifyErr := common.Verify(domain, domainVerify)
	if domainVerifyErr != nil {
		common.GinFailWithMessage(domainVerifyErr.Error(), c)
		return
	}

	if f, err := c.FormFile("CrtFile"); err == nil {
		name := domain.Domain + "_CrtFile.crt"
		if err := SaveFile(f,name,c); err != nil {
			common.GinFailWithMessage(fmt.Sprintf("证书文件保存失败%v", err), c)
			return
		}
	}

	if f, err := c.FormFile("KeyFile"); err == nil {
		name := domain.Domain + "_KeyFile.key"
		if err := SaveFile(f,name,c); err != nil {
			common.GinFailWithMessage(fmt.Sprintf("证书文件保存失败%v", err), c)
			return
		}
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
	s := c.PostForm("Domain")
	if err := json.Unmarshal([]byte(s), &domain); err != nil {
		common.GinFailWithMessage(fmt.Sprintf("数据解析失败%v", err), c)
	}
	domainVerify := common.Rules{
		"ID":          		{common.NotEmpty()},
		"Domain":      		{common.NotEmpty()},
	}
	domainVerifyErr := common.Verify(domain, domainVerify)
	if domainVerifyErr != nil {
		common.GinFailWithMessage(domainVerifyErr.Error(), c)
		return
	}

	if f, err := c.FormFile("CrtFile"); err == nil {
		name := domain.Domain + "_CrtFile.crt"
		if err := SaveFile(f,name,c); err != nil {
			common.GinFailWithMessage(fmt.Sprintf("证书文件保存失败%v", err), c)
			return
		}
	}

	if f, err := c.FormFile("KeyFile"); err == nil {
		name := domain.Domain + "_KeyFile.key"
		if err := SaveFile(f,name,c); err != nil {
			common.GinFailWithMessage(fmt.Sprintf("证书文件保存失败%v", err), c)
			return
		}
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

func SaveFile(f *multipart.FileHeader,SaveName string,c *gin.Context) (err error) {
	common.CreateDir(DomainPath)
	if !ExtValidator(f.Filename) {
		err = errors.New("后缀名不符合上传要求")
		return err
	}
	Filename := DomainPath + SaveName
	if er := c.SaveUploadedFile(f, Filename); er != nil {
		err = errors.New(fmt.Sprintf("文件保存失败%v", er))
		return err
	}
	return nil
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

func QueryApiProxy(ApiHost string) (bool, error) {
	var domain models.Domain
	if err := models.DB.Where("Domain = ? ", ApiHost).First(&domain).Error; err != nil {
		return false, err
	}else {
		return domain.IsHostAgent, err
	}
}
