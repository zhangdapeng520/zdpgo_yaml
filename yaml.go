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
	return y.Read(configPath, configObj)
}

// Read 读取yaml文件并映射到结构体
func (y *Yaml) Read(configPath string, configObj interface{}) error {
	// 读取配置文件
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	// 解析配置文件
	err = yaml.Unmarshal(yamlFile, configObj)
	return err
}

// Save 保存数据到yaml
func (y *Yaml) Save(filePath string, configObj interface{}) error {
	data, err := yaml.Marshal(configObj)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

// Marshal 将结构体对象序列化为yaml数据
func (y *Yaml) Marshal(in interface{}) (out []byte, err error) {
	data, err := yaml.Marshal(in)
	return data, err
}

// DumpsString 将结构体对象序列化为yaml数据，得到的是字符串
func (y *Yaml) DumpsString(in interface{}) (result string, err error) {
	data, err := y.Marshal(in)
	return string(data), err
}

func (y *Yaml) Dumps(in interface{}) (result []byte, err error) {
	data, err := y.Marshal(in)
	return data, err
}

// Unmarshal 将yaml数据反序列化为结构体
func (y *Yaml) Unmarshal(in []byte, out interface{}) (err error) {
	err = yaml.Unmarshal(in, out)
	if err != nil {
		return err
	}
	return nil
}

// LoadsString 加载yaml字符串为结构体
func (y *Yaml) LoadsString(in string, out interface{}) (err error) {
	err = y.Unmarshal([]byte(in), out)
	if err != nil {
		return err
	}
	return nil
}

func (y *Yaml) Loads(in []byte, out interface{}) (err error) {
	err = y.Unmarshal(in, out)
	if err != nil {
		return err
	}
	return nil
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
