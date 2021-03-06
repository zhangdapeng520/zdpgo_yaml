package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_yaml"
)

//定义conf类型
//类型里的属性，全是配置文件里的属性
type conf struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

func main() {
	// 创建配置对象
	var c conf

	// 读取配置
	zdpgo_yaml.ReadConfig("examples/read_config/test.yaml", &c)

	// 查看读取结果
	fmt.Println("读取配置：", c.Host, c.User, c.Pwd, c.Dbname)
}
