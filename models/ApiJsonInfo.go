package models

import "time"

type ApiJsonInfo struct {
	ID          uint      `json:"ID" gorm:"comment:'ID';AUTO_INCREMENT"`
	ApiID       uint      `json:"ApiID" gorm:"comment:'关联接口ID'"`
	ParamType   uint      `json:"ParamType" gorm:"comment:'参数类型'"`
	Parameter   string    `json:"Parameter" gorm:"comment:'校验接口参数'"`
	JsonContent string    `json:"JsonContent" gorm:"comment:'返回接口内容';Size:2048"`
	UpdatedUser string    `json:"UpdatedUser" gorm:"comment:'最后一次修改用户'"`
	IsOpen      bool      `json:"IsOpen" gorm:"comment:'是否启动校验'"`
	UpdatedAt   time.Time `json:"UpdatedAt" gorm:"comment:'最后一次修改时间'"`
}

func (s *ApiJsonInfo) TableName() string {
	return "ApiJsonInfo"
}
