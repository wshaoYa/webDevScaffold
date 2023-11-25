# webDevScaffold
web development scaffold  （golang web开发通用脚手架-模板）

**用途**：新开发一个项目时，这些一般是必要的前置准备工作，总结成脚手架模板，旨在提高后续开发效率。

## 项目结构

- `dao` (数据访问层)
- `logger`（日志层）
- `routers`（路由层）
- `settings`（配置层）
- `config.yaml`（配置文件）
- `main.go`（程序启动入口）

## main.go 程序入口

1. `viper` 加载配置
2. `zap` 初始化日志
3. `sqlx` 初始化mysql连接
4. `go-redis` 初始化redis连接
5. `gin` 注册路由
6. `endless` 启动服务(优雅关机、重启)

## **项目来源**

七米老师博客/视频

- [【置顶】Go语言学习之路/Go语言教程 | 李文周的博客 (liwenzhou.com)](https://www.liwenzhou.com/posts/Go/golang-menu/)
- [Go Web开发进阶实战（gin框架） - 网易云课堂 (163.com)](https://study.163.com/course/courseMain.htm?courseId=1210171207)

## 自使用注意事项

- 配置文件-更改
  - 项目名称、日志名称、数据库密码、端口等
- 引用包名称-更改
  - `import` 项目内的包时，第一目录名称由 `webDevScaffold` 改成自己项目名称
- 自行增加 controller、logic/service、models、pkg等目录（待后续补充完善此仓库qwq）