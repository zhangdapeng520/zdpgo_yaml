# zdpgo_yaml
Golang处理yaml文件的专用库

项目地址：https://github.com/zhangdapeng520/zdpgo_yaml

## 版本历史
- 2022年3月28日 版本0.1.0 基本功能，读取配置和默认配置

## 使用案例
### 案例1：读取yaml配置
```go
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
	// 创建yaml对象
	y := zdpgo_yaml.New()

	// 创建配置对象
	var c conf

	// 读取配置
	y.ReadConfig("test.yaml", &c)

	// 查看读取结果
	fmt.Println("读取配置：", c.Host, c.User, c.Pwd, c.Dbname)
}
```

### 案例2：读取默认配置
```go
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
	// 创建yaml对象
	y := zdpgo_yaml.New()

	// 创建配置对象
	var c conf

	// 读取默认配置
	// 默认读取config/config.yaml和config/secret/.config.yaml两个文件
	// 后面的文件是私密文件，可以存储密码之类的。私密文件会覆盖公开文件中的配置。
	y.ReadDefaultConfig(&c)

	// 查看读取结果
	fmt.Println("读取配置：", c.Host, c.User, c.Pwd, c.Dbname)
}
```