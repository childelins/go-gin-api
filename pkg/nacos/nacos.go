package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type ConfigClient struct {
	iClient config_client.IConfigClient
}

func NewConfigClient(host string, port uint64, namespace string) (*ConfigClient, error) {
	sc := []constant.ServerConfig{
		{
			IpAddr: host,
			Port:   port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         namespace, // 如果需要支持多namespace，我们可以产生多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "storage/nacos/log",
		CacheDir:            "storage/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	if err != nil {
		return nil, err
	}

	return &ConfigClient{configClient}, nil
}

func (cfg *ConfigClient) GetConfig(dataId, group string) (string, error) {
	return cfg.iClient.GetConfig(vo.ConfigParam{DataId: dataId, Group: group})
}
