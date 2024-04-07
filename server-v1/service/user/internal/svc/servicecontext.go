package svc

import (
	"gorm.io/gorm"
	gormInit "server-v1/internal/gormInit"
	"server-v1/service/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := gormInit.GetSqlConnect(gormInit.SqlConfig{
		Host:     c.Mysql.Host,
		Port:     c.Mysql.Port,
		Username: c.Mysql.Username,
		Password: c.Mysql.Password,
		DbName:   c.Mysql.Dbname,
	})
	return &ServiceContext{
		Config: c,
		Db:     db,
	}
}
