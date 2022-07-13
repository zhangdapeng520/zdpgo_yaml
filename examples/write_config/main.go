package main

import (
	"fmt"
	"io/ioutil"

	"github.com/zhangdapeng520/zdpgo_yaml"
)

// Config 类型里的属性，全是配置文件里的属性
type Config struct {
	Version       string   `yaml:"version" json:"version"`
	AesKey        string   `yaml:"aes_key" json:"aes_key"`
	CenterHost    string   `yaml:"center_host" json:"center_host"`
	CenterPort    int      `yaml:"center_port" json:"center_port"`
	NodeId        string   `yaml:"node_id" json:"node_id"`
	HeartInterval int      `yaml:"heart_internal" json:"heart_internal"`
	Region        string   `yaml:"region" json:"region"`
	Type          []string `yaml:"type" json:"type"`
}

func main() {
	var (
		config *Config
		err    error
	)

	configPath := "examples/write_config/config.yaml"
	y := zdpgo_yaml.New()

	//读取yaml文件到缓存中
	c, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("读取配置文件失败：", err)
		return
	}

	// yaml文件内容影射到结构体中
	err = y.Unmarshal(c, &config)
	if err != nil {
		fmt.Println("解析配置文件失败：", err)
		return
	}

	// 处理version
	if config.Version == "" {
		config.Version = "1.0.0.1"
	}

	// 自动生成AES Key
	if config.AesKey == "" {
		config.AesKey = y.RandomStr(16)
	}

	// 自动生成和NodeId
	if config.NodeId == "" {
		config.NodeId = y.RandomStr(16)
	}

	// 处理port
	if config.CenterPort == 0 {
		config.CenterPort = 9090
	}

	// 处理心跳间隔
	if config.HeartInterval == 0 {
		config.HeartInterval = 3
	}

	// 处理区域
	if config.Region == "" {
		config.Region = "成都"
	}

	// 处理类型
	if config.Type == nil || len(config.Type) == 0 {
		config.Type = []string{"sev"}
	}

	// 保存
	err = y.Save(configPath, config)
	if err != nil {
		fmt.Print("保存配置文件失败：", err)
		return
	}

	fmt.Println("保存配置成功")
}
