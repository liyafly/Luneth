package model

import (
	"fmt"
	"github.com/liyafly/Luneth/global"
	"github.com/liyafly/Luneth/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

type TableNamed interface {
	TableName() string
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	dsn := fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParesTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}
	db.Config.NamingStrategy = schema.NamingStrategy{SingularTable: true}
	mysqlDB, e := db.DB()
	if e != nil {
		return nil, e
	}
	mysqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	mysqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}
