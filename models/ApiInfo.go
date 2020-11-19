package models

import "time"

type ApiInfo struct {
	ID          uint       `json:"ID" gorm:"comment:'接口ID';AUTO_INCREMENT"`
	Name        string     `json:"Name" gorm:"comment:'接口名称'"`
	Path        string     `json:"Path" gorm:"comment:'接口路径'"`
	Domain      string     `json:"Domain" gorm:"comment:'接口域名'"`
	UpdatedUser string     `json:"UpdatedUser" gorm:"comment:'最后一次修改用户'"`
	UpdatedAt   time.Time  `json:"UpdatedAt" gorm:"comment:'最后一次修改时间'"`
}

func (s *ApiInfo) TableName() string {
	return "ApiInfo"
}
