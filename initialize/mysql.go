package initialize

import (
	"douyin/global"
	"douyin/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// Mysql 配置MySQL数据库
func Mysql() {
	m := global.Config.Mysql
	var dsn = fmt.Sprintf("%s:%s@%s", m.Username, m.Password, m.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(fmt.Errorf("gorm open error: %s\n", err))
	}

	// 自动迁移schema
	if err := db.AutoMigrate(&model.User{}, &model.Video{}, &model.Favorite{},
		&model.Comment{}, &model.Relation{}); err != nil {
		panic(fmt.Errorf("db automigrate err: %s\n", err))
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("db db error: %s\n", err))
	}
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
	global.Db = db
}
