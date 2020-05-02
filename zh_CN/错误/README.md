#  错误处理

* 应用场景：可以在发生特定的http错误代码时定义自己的处理程序。
* 错误代码是大于或等于400的http状态代码，如404未找到和500内部服务器。

```
package main

import (
    "github.com/kataras/iris"
)

func main(){
    app := iris.New()
    app.OnErrorCode(iris.StatusNotFound, notFound)
    app.OnErrorCode(iris.StatusInternalServerError, internalServerError)
        // 注册一个处理函数处理HTTP错误codes >=400:
    // app.OnAnyErrorCode(handler)
    app.Get("/", index)
    app.Run(iris.Addr(":8080"))
}

func notFound(ctx iris.Context) {
    // 当http.status=400 时向客户端渲染模板$views_dir/errors/404.html
    ctx.View("errors/404.html")
}

//当出现错误的时候，再试一次
func internalServerError(ctx iris.Context) {
    ctx.WriteString("Oups something went wrong, try again")
}
func index(ctx context.Context) {
    ctx.View("index.html")
}
```