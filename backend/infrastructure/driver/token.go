package driver

import "home/config"

type TokenDriver struct {
	*config.Config
}

func NewTokenDriver(conf *config.Config) *TokenDriver {
	return &TokenDriver{
		conf,
	}
}
