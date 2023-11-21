package main

import (
	"fmt"
	"webDevScaffold/dao/mysql"
	"webDevScaffold/dao/redis"
	"webDevScaffold/logger"
	"webDevScaffold/routers"
	"webDevScaffold/settings"

	"github.com/fvbock/endless"

	"go.uber.org/zap"
)

/*
golang web项目 开发脚手架 模板

1、viper 加载配置
2、zap 初始化日志
3、sqlx 初始化mysql
4、go-redis 初始化redis
5、gin 注册路由
6、endless 启动服务

*/

func main() {
	// 1、加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("settings.Init() err:", err)
		return
	}

	// 2、初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		zap.L().Error("logger.Init err ! ", zap.Error(err))
		return
	}

	// 刷新所有缓冲的日志
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			l.Error("zap sync err ! ", zap.Error(err))
			return
		}
	}(zap.L())
	zap.L().Debug("logger init success")

	// 3、初始化MySQL连接
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		zap.L().Error("mysql init error !! ", zap.Error(err))
		return
	}

	// 关闭mysql连接
	defer func() {
		err := mysql.Close()
		if err != nil {
			zap.L().Error("mysql close err !! ", zap.Error(err))
			return
		}
	}()
	zap.L().Debug("mysql init success")

	// 4、初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		zap.L().Error("redis init error ! ", zap.Error(err))
		return
	}

	// 关闭redis连接
	defer func() {
		err := redis.Close()
		if err != nil {
			zap.L().Error("redis close error ! ", zap.Error(err))
			return
		}
	}()
	zap.L().Debug("redis init success")

	// 5、注册路由
	r := routers.Setup(settings.Conf.Mode)

	// 6、启动服务（endless 优雅关机、重启）
	err := endless.ListenAndServe(
		fmt.Sprintf(":%d", settings.Conf.Port), r)
	if err != nil {
		zap.L().Error("endless.ListenAndServe error ! ", zap.Error(err))
		return
	}
}
