package models

import "time"

type Domain struct {
	ID uint					`json:"ID" gorm:"comment:'ID';AUTO_INCREMENT"`
	Domain string			`json:"Domain" gorm:"comment:'域名';not null;unique"`
	CrtFilePath string		`json:"CrtFilePath" gorm:"comment:'Crt文件路径'"`
	KeyFilePath string		`json:"KeyFilePath" gorm:"comment:'Key文件路径'"`
	HostAgent string		`json:"HostAgent" gorm:"comment:'接口反向代理HOST'"`
	IsHostAgent	bool		`json:"IsHostAgent" gorm:"comment:'是否启动反向代理'"`
	UpdatedUser string		`json:"UpdatedUser" gorm:"comment:'最后一次修改用户'"`
	CreatedAt time.Time		`json:"CreatedAt" gorm:"comment:'创建时间'"`
	UpdatedAt time.Time		`json:"UpdatedAt" gorm:"comment:'最后一次修改时间'"`
	DeletedAt *time.Time	`json:"DeletedAt" gorm:"comment:'删除时间'"`
}

func (s *Domain) TableName() string {
	return "Domain"
}

