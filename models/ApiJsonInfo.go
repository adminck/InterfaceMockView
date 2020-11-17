package models

import "time"

type ApiJsonInfo struct {
	ID           uint       `json:"ID" gorm:"comment:'ID';AUTO_INCREMENT"`
	ApiID        uint       `json:"ApiID" gorm:"comment:'关联接口ID'"`
	ParamType    uint       `json:"ParamType" gorm:"comment:'参数类型'"`
	Parameter    string     `json:"Parameter" gorm:"comment:'校验接口参数'"`
	JsonFilePath string     `json:"JsonFilePath" gorm:"comment:'返回接口文件路径'"`
	UpdatedUser  string     `json:"UpdatedUser" gorm:"comment:'最后一次修改用户'"`
	IsOpen       bool       `json:"IsOpen" gorm:"comment:'是否启动校验'"`
	CreatedAt    time.Time  `json:"CreatedAt" gorm:"comment:'创建时间'"`
	UpdatedAt    time.Time  `json:"UpdatedAt" gorm:"comment:'最后一次修改时间'"`
	DeletedAt    *time.Time `json:"DeletedAt" gorm:"comment:'删除时间'"`
}

func (s *ApiJsonInfo) TableName() string {
	return "ApiJsonInfo"
}
