package main

import (
	"InterfaceMockView/models"
	"InterfaceMockView/utils/log"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
)

type Config struct {
	Db  *models.DbConfig `json:"db"`
	Log *log.Config      `json:"log"`
}

var (
	defaultPrintScreen    = true
	defaultLogLevel       = "info"
	defaultLogReserveDays = 3
	defaultLogFileMaxSize = 100
)

func NewConfig() *Config {
	g := &Config{
		Log: &log.Config{
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
