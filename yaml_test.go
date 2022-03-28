package zdpgo_yaml

import (
	"testing"
)

func getYaml() *Yaml {
	y := Yaml{}
	return &y
}

//定义conf类型
//类型里的属性，全是配置文件里的属性
type conf struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

// 测试读取配置
func TestYaml_ReadConfig(t *testing.T) {
	y := getYaml()
	var c conf
	err := y.ReadConfig("test.yaml", &c)
	if err != nil {
		t.Error(err)
	}
	t.Log("配置：", c)
}

// 测试读取默认配置
func TestYaml_ReadDefaultConfig(t *testing.T) {
	y := getYaml()
	var c conf
	err := y.ReadDefaultConfig(&c)
	if err != nil {
		t.Error(err)
	}
	t.Log("配置：", c)
	t.Log(c.Host, c.User, c.Pwd, c.Dbname)
}
