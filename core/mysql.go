package core

import (
	"app_ws/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

// Sqlx Init 初始化MySQL连接
func Sqlx() (db *sqlx.DB) {
	// "user:password@tcp(host:port)/dbname"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", global.GVA_CONFIG.Mysql.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err := sqlx.Connect("mysql", global.GVA_CONFIG.Mysql.Dsn())
	if err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		os.Exit(0)
		return nil
	}
	db.SetMaxOpenConns(global.GVA_CONFIG.Mysql.MaxOpenConns)
	db.SetMaxIdleConns(global.GVA_CONFIG.Mysql.MaxIdleConns)
	return db
}
