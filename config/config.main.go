package config

import (
	"time"
)

type ServerCfg struct {
	Port string
}

type HelloCfg struct {
	HelloCount int
}

type RedisConfig struct {
	Host     string
	PoolSize int8
	MaxIdle  int8
	Timeout  time.Duration
}
