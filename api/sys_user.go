package API

import (
	"InterfaceMockView/models"
	"InterfaceMockView/utils/common"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type LoginResponse struct {
	User      models.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type SysUserResult struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	UpDataPassword string `json:"rePassword"`
}

func Login(c *gin.Context) {
	var L SysUserResult
	_ = c.ShouldBindJSON(&L)
	UserVerify := common.Rules{
		"Username": {common.NotEmpty()},
		"Password": {common.NotEmpty()},
	}
	UserVerifyErr := common.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		common.GinFailWithMessage(UserVerifyErr.Error(), c)
		return
	}

	U := &models.SysUser{Username: L.Username, Password: common.MD5V([]byte(L.Password))}
	if err := U.Login(); err != nil {
		common.GinFailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), c)
	} else {
		tokenNext(c, U)
	}

}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user *models.SysUser) {
	j := common.NewJWT()
	clams := common.CustomClaims{
		ID:       user.ID,
		NickName: user.NickName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,  // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60, // 过期时间 一小时
			Issuer:    "qmPlus",                  // 签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		common.GinFailWithMessage("获取token失败", c)
		return
	} else {
		common.GinOkWithData(LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	}
}

func UserUpdate(c *gin.Context) {
	var L SysUserResult
	_ = c.ShouldBindJSON(&L)
	UserVerify := common.Rules{
		"ID":             {common.NotEmpty()},
		"Username":       {common.NotEmpty()},
		"Password":       {common.NotEmpty()},
		"UpDataPassword": {common.NotEmpty()},
	}
	UserVerifyErr := common.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		common.GinFailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	U := &models.SysUser{ID: L.ID, Username: L.Username, Password: common.MD5V([]byte(L.Password))}
	if err := U.Update(common.MD5V([]byte(L.UpDataPassword))); err != nil {
		common.GinFailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), c)
	} else {
		common.GinOkWithMessage("密码修改成功！", c)
	}
}

func UserCreate(c *gin.Context) {
	var L SysUserResult
	_ = c.ShouldBindJSON(&L)
	UserVerify := common.Rules{
		"Username": {common.NotEmpty()},
		"Password": {common.NotEmpty()},
	}
	UserVerifyErr := common.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		common.GinFailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	U := &models.SysUser{Username: L.Username, Password: common.MD5V([]byte(L.Password))}
	if err := U.Create(); err != nil {
		common.GinFailWithMessage(fmt.Sprintf("创建失败！%v", err), c)
	} else {
		common.GinOkWithMessage("创建成功！", c)
	}
}
