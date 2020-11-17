package main

import (
	"InterfaceMockView/models"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

type Config struct {
	Db  *models.DbConfig `json:"db"`
	Log *LogConfig       `json:"log"`
}

type LogConfig struct {
	LogLevel    string `json:"logLevel"`    // 日志级别，支持：off/trace/debug/info/warn/error/panic/fatal
	ReserveDays int    `json:"reserveDays"` // 日志文件保留天数
	MaxSize     int    `json:"maxSize"`     // 日志文件最大大小，单位：MB
	PrintScreen bool   `json:"printScreen"` // 是否打印至标准输出
}

func (l *LogConfig) String() string {
	return fmt.Sprintf("%+v", *l)
}

var (
	defaultPrintScreen    = true
	defaultLogLevel       = "info"
	defaultLogReserveDays = 3
	defaultLogFileMaxSize = 100
)

func NewConfig() *Config {
	g := &Config{
		Log: &LogConfig{
			LogLevel:    defaultLogLevel,
			ReserveDays: defaultLogReserveDays,
			MaxSize:     defaultLogFileMaxSize,
			PrintScreen: defaultPrintScreen,
		},
	}

	return g
}

func (c *Config) Load() error {
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return errors.Wrap(err, "read config file error")
	}

	if err := json.Unmarshal(content, c); err != nil {
		return errors.Wrap(err, "unmarshal config error")
	}

	return nil
}
