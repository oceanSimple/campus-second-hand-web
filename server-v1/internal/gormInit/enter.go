package gormInit

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func GetSqlConnect(config SqlConfig) *gorm.DB {
	var _db *gorm.DB
	timeout := "10s" // Database connection timeout
	//Stitch the dsn parameters below
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		config.Username, config.Password, config.Host, config.Port, config.DbName, timeout)

	// Connect to sql to obtain a database instance for subsequent database read and write operations
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect to sql database, error=" + err.Error())
	}
	// 设置会话时区
	_db.Exec("SET time_zone = '+08:00'")
	sqlDB, _ := _db.DB()
	sqlDB.SetMaxIdleConns(20)  // Sets the maximum number of connections in the free connection pool
	sqlDB.SetMaxOpenConns(100) // Sets the maximum number of open database connections

	return _db
}
