# 应用常用框架结构

### 常用推荐应用框架结构（简单目录结构）

```
conf            配置文件（conf|config）
controllers     控制器 入参处理 api的入口
datasource      数据源配置（数据库、数据文件等） 
models          数据模型
db              SQL数据文件 Postman接口文件等
repo            数据库的操作
middleware 中间件 jwt实现
route           路由注册（route｜router｜routing）
service         业务逻辑代码（service｜services）
util           工具类（util｜utils）
config.json 配置文件的映射
main.go 主程序入口
go.mod
go.sum
```

###### 推荐使用go modules管理package应用包

