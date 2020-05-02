
# 初始化应用 Run Web Application

### 应用创建

##### 框架启动参数常用主配置项：

* ```iris.WithConfiguration(iris.Configuration{ })```
* ```iris.WithoutServerError(iris.ErrServerClosed)```
* ```app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))```

```

package main

import (
	"github.com/kataras/iris"
)

func main() {

	//创建初始化对象
	app := iris.New()

	//设置日志级别
	app.Logger().SetLevel("debug")
	
	//配置监听端口
	//application.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	app.Run(iris.Addr(":8080"))

}

```

备注：配置参数见配置说明。
