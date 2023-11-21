package mysql

import (
	"fmt"
	"webDevScaffold/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("sqlx connect error !! ", zap.Error(err))
		return
	}

	// 最大开放连接、空闲连接
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() (err error) {
	err = db.Close()
	if err != nil {
		zap.L().Error("mysql close error ! ", zap.Error(err))
		return
	}
	return
}
