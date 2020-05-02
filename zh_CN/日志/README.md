#  日志


### 日志基本

设置日志级别：

* disable
* fatal
* error
* warn
* info
* debug

##### 框架封装日志模式：
```
Example Code:

	golog.SetLevel("disable")
	golog.SetLevel("fatal")
	golog.SetLevel("error")
	golog.SetLevel("warn")
	golog.SetLevel("info")
	golog.SetLevel("debug")

```

##### 框架调用：

常用场景：配合```context.Context```对象使用，跟踪记录应用运行生命周期的状态日志信息
```
//log level
app.Logger().SetLevel("debug")
```


