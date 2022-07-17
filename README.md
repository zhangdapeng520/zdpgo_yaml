# zdpgo_yaml

Golang 处理 yaml 文件的专用库

项目地址：https://github.com/zhangdapeng520/zdpgo_yaml

## 版本历史

- v0.1.0 2022/03/28 基本功能，读取配置和默认配置
- v0.1.1 2022/04/21 新增：读写配置
- v0.1.2 2022/07/08 优化：代码优化
- v0.1.3 2022/07/13 优化：代码优化
- v0.1.4 2022/07/17 新增：读取配置文件

## 使用案例

## 读取配置文件
配置文件：test.yaml
```yaml
host: 127.0.0.1
user: zhangdapeng520
pwd: zhangdapeng520
dbname: zhangdapeng520
```

读取配置文件：main.go
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
	// 创建配置对象
	var c conf

	// 读取配置
	zdpgo_yaml.ReadConfig("examples/read_config/test.yaml", &c)

	// 查看读取结果
	fmt.Println("读取配置：", c.Host, c.User, c.Pwd, c.Dbname)
}
```
