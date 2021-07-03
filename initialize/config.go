package initialize

import (
	"encoding/json"
	"time"

	"github.com/childelins/go-gin-api/global"
	"github.com/childelins/go-gin-api/pkg/config"
	"github.com/childelins/go-gin-api/pkg/nacos"
)

func InitConfig() error {
	conf, err := config.NewConfig("config.yaml")
	if err != nil {
		return err
	}

	err = conf.Unmarshal(&global.NacosConfig)
	if err != nil {
		return err
	}

	// 从nacos中读取配置信息
	client, err := nacos.NewConfigClient(global.NacosConfig.Host, global.NacosConfig.Port, global.NacosConfig.Namespace)
	if err != nil {
		return err
	}

	content, err := client.GetConfig(global.NacosConfig.DataId, global.NacosConfig.Group)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		return err
	}

	global.ServerConfig.WriteTimeout *= time.Second
	global.ServerConfig.ReadTimeout *= time.Second
	global.ServerConfig.WriteTimeout *= time.Second
	return nil
}

/*
func InitConfig() error {
	conf, err := config.NewConfig("config.yaml.bak")
	if err != nil {
		return err
	}
	err = conf.Unmarshal(&global.ServerConfig)
	if err != nil {
		return err
	}

	global.ServerConfig.WriteTimeout *= time.Second
	global.ServerConfig.ReadTimeout *= time.Second
	global.ServerConfig.WriteTimeout *= time.Second
	return nil
}
*/
