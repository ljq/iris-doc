
# 框架配置参数简述


## 配置分两种场景

* 框架基础服务配置
* 应用自定义配置

### 框架基础配置用法：

参数：
* ```iris.WithConfiguration(iris.Configuration{})```
* ```iris.WithoutServerError(iris.ErrServerClosed)```
* ```app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))```

配置项参数说明：

* 使用不定参数，引入常用配置型:
app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
    * ** iris.WithConfiguration(iris.Configuration{})** 
	* ** iris.WithoutServerError(iris.ErrServerClosed)** 

* iris.WithConfiguration(iris.Configuration{})

配置支持以内部加载和外部加载：
* 以内部以struct结构体传参直接配置；
* 外部文件加载模式支持：toml、yaml、json、xml、特殊场景自定义格式(环境变量、自定义文件)等。

* 个人推荐使用** 结构体内嵌传参**  或 ** *YAML、TOML** 中任一种，视使用场景而定。原因：使用json、xml需要编码时手动处理格式转换。

```

package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

```

### 框架服务配置方式

* 内嵌：
    * struct对象结构体传参数

* 外部加载：
    * YAML
    * TOML

补充说明：其他自定义方式:
针对特定场景，系统环境变量、加载解析外部自定义文件，赋值街头体，灵活多变。

### 1.struct结构体内嵌传参

```
//iris default configuration object
iris.WithConfiguration(iris.Configuration{
        DisableStartupLog:                 false,
        DisableInterruptHandler:           false,
        DisablePathCorrection:             false,
        EnablePathEscape:                  false,
        FireMethodNotAllowed:              false,
        DisableBodyConsumptionOnUnmarshal: false,
        DisableAutoFireStatusCode:         false,
        TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
        Charset:                           "UTF-8",
    }))
```

### 2.引入外部文件

* app.Configure( iris.WithConfiguration( iris.YAML() ) )
* app.Configure( iris.WithConfiguration( iris.TOML() ) )


## 应用配置

外部文件加载常用格式：
    * YAML
    * TOML
    * JSON
    * XML

注意在格式转换时，注意解析的处理，创建承载的街头体。

例JSON格式处理方法(** 处理方式多样，不拘泥于一种方式，视场景而定** )：
大致处理流程：
* 加载json文件
* 解析json数据
* 结构体赋值

例：
```
//声明struct结构体承载
type Coniguration struct {
	AppName string `json:"appname"`
	Port    int    `json:"port"`
    Extendfield string  `json:"extendfield"`
}
```

应用加载配置项
```

//loading json file
file, _ := os.Open("config.json")
defer file.Close()

decoder := json.NewDecoder(file)
conf := Coniguration{}
err := decoder.Decode(&conf)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(conf.AppName)

```






