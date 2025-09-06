package migrations

import (
	"gorm.io/gorm"
)

func AddIDToAccounts(db *gorm.DB) error {
	// 检查表是否存在
	if !db.Migrator().HasTable("accounts") {
		return nil
	}

	// 检查是否已有ID列
	if db.Migrator().HasColumn("accounts", "id") {
		return nil
	}

	// 添加ID列
	err := db.Exec("ALTER TABLE accounts ADD COLUMN id INT AUTO_INCREMENT PRIMARY KEY FIRST").Error
	if err != nil {
		return err
	}

	return nil
}
