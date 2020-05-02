# 路由策略

* 简单路由
* 分组路由策略

# 基础用法

```
//func handler(ctx iris.Context)
app.Get("/", handler)
```

```
app.Handle("GET", "/users", func(ctx iris.Context) {
    ctx.HTML("request path: /users")
})
```

### 分组路由

* 分组路由 由路径前缀分组的一组路由可以（可选）共享相同的中间件处理程序和模板布局。一个组也可以有一个嵌套组。
* .Party 正在用于分组路由，开发人员可以声明无限数量的（嵌套）组。

```
// /users
userParty := app.Party("/users", func(context context.Context) { 
    app.Logger().Info("request path: users")
    context.Next()
})

// /users/login
userParty.Get("/login", func(context context.Context) {
	app.Logger().Info("user login")
})
```

### 路由处理时注意context.Context的处理

##### ```context.Context```核心结构

```
type Context interface {
  Deadline() (deadline time.Time, ok bool)
  Done() <-chan struct{}
  Err() error
  Value(key interface{}) interface{}
}
```

##### 路由中间件

```

func userMiddleware(context iris.Context) {

    // 路由器中Done()方法执行的前提是在Next()方法之后
    // 应用场景：此方法可以跟踪路由器的整个周期的状态信息。
	context.Next()

}
```

```
usersRouter := app.Party("/users", userMiddleware)

//Done：
usersRouter.Done(func(context context.Context) {
    context.Application().Logger().Infof("response sent to " + context.Path())
})
```

### 复杂路由策略

##### 动态路由参数

特性：macros：Iris有自己的路径（就像编程语言一样），用于路由的路径语法及其路径参数解析和评估. 使用了"macros"(宏)的概念.

|       参数模式            |       框架处理流程        |
|:------------------------:|:--------------------:|
| 不含需要任何特殊正则表达式 | 只是用低级路径语法注册路由 |
| 含特殊正则表达式          | 预先编译正则表达式并添加必要的中间件 |

相对其他框架而言，** 使用macros的优势** ：这意味着相对于其他路由器或Web框架 您的性能成本为零。

具体参数参考(route-params.md)官方文档:
(route-params.md)[route-params.md]


### 小结

路由分组的处理灵活多变，可以借助contenxt进行路由到控制器的各个生命周期状态进行跟踪，除在web领域，也可以在一些特殊业务场景进行业务分发。


