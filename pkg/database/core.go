package database

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type DBInfo struct {
	FilePath string
	Server   string
	Password string
}

func InitDB(dbInfo *DBInfo) {
	var err error
	// db, err = gorm.Open(sqlite.Open(dbInfo.FilePath), &gorm.Config{})
	dsn := fmt.Sprintf("root:%v@tcp(%v)/my_dtcg?charset=utf8mb4&parseTime=True&loc=Local", dbInfo.Password, dbInfo.Server)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("连接数据库失败: %v", err)
	}

	// AutoMigrate 用来刷新数据表，不存在则创建，表名默认为结构体名称的复数，e.g.这里会创建一个名为 products 的表，假如 Product 为 ProductTest，则会创建出一个名为 product_test 的表
	// 结构体中的每个字段都是该表的列，字段名称即是表中列的名称，如果字段名中有多个大写字母，则列名使用下划线分隔，e.g.CreatedAt 字段的列名为 cretaed_at
	// 当结构体中增加字段时，会自动在表中增加列；但是删除结构体中的属性时，并不会删除列
	db.AutoMigrate(&CardDescFromOfficial{}, &CardGroupFromOfficial{}, &CardDescFromDtcgDB{}, &CardGroupFromDtcgDB{})
}
