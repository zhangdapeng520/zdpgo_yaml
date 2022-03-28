package zdpgo_yaml

import (
	"github.com/zhangdapeng520/zdpgo_yaml/yaml"
	"io/ioutil"
)

type Yaml struct {
}

func New() *Yaml {
	y := Yaml{}
	return &y
}

// ReadConfig 读取yaml文件配置
func (y *Yaml) ReadConfig(configPath string, configObj interface{}) error {
	// 读取配置文件
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	// 解析配置文件
	err = yaml.Unmarshal(yamlFile, configObj)
	return err
}

// ReadDefaultConfig 读取默认配置。默认使用config/config.yaml和config/secret/.config.yaml
func (y *Yaml) ReadDefaultConfig(configObj interface{}) error {
	// 默认的配置我呢就
	publicConfig := "config/config.yaml"
	privateConfig := "config/secret/.config.yaml"

	// 读取公共配置
	err := y.ReadConfig(publicConfig, configObj)
	if err != nil {
		return err
	}

	// 读取私有的配置
	err = y.ReadConfig(privateConfig, configObj)
	return err
}
