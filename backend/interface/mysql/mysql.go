package mysql

import (
	"bSocial/domain"
	"bSocial/helpers"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySql *gorm.DB

func InitConnection() error {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/bSocialDB?charset=utf8mb4&parseTime=True&loc=Local",
			helpers.CONFIG.MySQL.DbUser,
			helpers.CONFIG.MySQL.DbPassword,
			helpers.CONFIG.MySQL.DbHost,
			helpers.CONFIG.MySQL.DbPort),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		return err
	}
	MySql = db
	return nil
}

func AutoMigrate() {
	MySql.AutoMigrate(domain.User{}, domain.Post{}, domain.Comment{}, domain.UserFollower{})
}
