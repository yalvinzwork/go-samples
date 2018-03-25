package config

import (
	"github.com/yalvinz/ipreit/util/config"
)

type MainConfig struct {
	Server ServerCfg
	Hello  HelloCfg
	Redis  map[string]*RedisConfig
}

func InitConfig(cfg interface{}, module string, path ...string) error {
	if len(path) == 0 {
		path = GetDefaultConfigPaths()
	}

	for _, p := range path {
		err := config.ReadModuleConfig(cfg, p, module)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetDefaultConfigPaths() []string {
	return []string{"files/etc/ipreit", "/etc/ipreit"}
}
