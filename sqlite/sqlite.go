package sqlite

import (
	"fmt"
	"linux-host/config"
	"sync"
	"time"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)
var (
	db   *gorm.DB
	once sync.Once
)
func  DbInit()  {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(config.ParamsConfig.GetString("sqlite.path")), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
			Logger:         logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			fmt.Println("数据库打开成功")
		} else {
			fmt.Println("数据库打开失败")
			return
		}
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println("通用数据库对象获取失败")
			return
		}
		if err != nil {
			sqlDB.Close()
			fmt.Println("数据库同步失败")
			return
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(100)
		sqlDB.SetConnMaxIdleTime(time.Hour)
		//err = db.AutoMigrate(&model.SsOrgan{})
		//err = db.AutoMigrate(&model.SsHost{})
		//err = db.AutoMigrate(&model.SsDefence{})
		//err = db.AutoMigrate(&model.SsZone{})
		//err = db.AutoMigrate(&model.SsDevice{})
	})
}
