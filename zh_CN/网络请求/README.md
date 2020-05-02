
# 网络请求


### 支持的方法

```
package main
    import (
        "github.com/kataras/iris"
    )
    func main(){
        app := iris.New()
        //GET 方法
        app.Get("/", handler)
        // POST 方法
        app.Post("/", handler)
        // PUT 方法
        app.Put("/", handler)
        // DELETE 方法
        app.Delete("/", handler)
        //OPTIONS 方法
        app.Options("/", handler)
        //TRACE 方法
        app.Trace("/", handler)
        //CONNECT 方法
        app.Connect("/", handler)
        //HEAD 方法
        app.Head("/", handler)
        // PATCH 方法
        app.Patch("/", handler)
        //任意的http请求方法如option等
        app.Any("/", handler)

    }
    func handler(ctx iris.Context){
        ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
    }
```