package models

import "time"

type ApiInfo struct {
	ID uint					`json:"ID" gorm:"comment:'接口ID';AUTO_INCREMENT"`
	Name string	 			`json:"Name" gorm:"comment:'接口名称'"`
	Path string  			`json:"Path" gorm:"comment:'接口路径'"`
	Domain string			`json:"Domain" gorm:"comment:'接口域名'"`
	IsOpen bool				`json:"IsOpen" gorm:"comment:'是否启动'"`
	UpdatedUser string		`json:"UpdatedUser" gorm:"comment:'最后一次修改用户'"`
	CreatedAt time.Time		`json:"CreatedAt" gorm:"comment:'创建时间'"`
	UpdatedAt time.Time		`json:"UpdatedAt" gorm:"comment:'最后一次修改时间'"`
	DeletedAt *time.Time	`json:"DeletedAt" gorm:"comment:'删除时间'"`
}

func (s *ApiInfo) TableName() string {
	return "ApiInfo"
}
