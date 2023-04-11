package model

import (
	"WowjoyProject/WADO_URL/pkg/setting"
	"database/sql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type FileData struct {
	file, ip, virpath sql.NullString
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.DBType, databaseSetting.DBConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	db.SetMaxIdleConns(databaseSetting.MaxIdleConns)

	return db, nil
}
