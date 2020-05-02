#  注意事项


### 版本注意事项

* 针对iris，Go 1.9支持类型别名，iris为Go 1.8.3版本已作预留（```kataras/iris/context.go```）对Go 1.9的限制构建访问权限声明了iris的所有类型别名。
    * ** 低于 Go 1.9 版本** ，** 必须** 导入```import github.com/kataras/iris/context```来创建一个Handler
    * ** 高于 Go 1.9 版本** ，** 无需** 导入```import github.com/kataras/iris/context```

