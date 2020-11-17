package models

import (
	"InterfaceMockView/utils/log"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Path         string `json:"path"`
	Dbname       string `json:"dbname"`
	Config       string `json:"config"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
	LogMode      bool   `json:"log_mode"`
}

func DBinit(config *DbConfig) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.Username+":"+config.Password+"@("+config.Path+")/"+config.Dbname+"?"+config.Config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("MySQL启动异常,%s", err))
	}
	DB = db
	DB.DB().SetMaxIdleConns(config.MaxIdleConns)
	DB.DB().SetMaxOpenConns(config.MaxOpenConns)
	DB.LogMode(config.LogMode)
	if err := DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&SysUser{}, &ApiInfo{}, &ApiJsonInfo{}, &Domain{}).Error; err != nil {
		log.Error(err)
	}
	return DB, nil
}
