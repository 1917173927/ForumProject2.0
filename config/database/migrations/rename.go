package migrations

import (
	"gorm.io/gorm"
)

func RenameReporterToUser(db *gorm.DB) error {
	migrator := db.Migrator()
	if migrator.HasTable("reports") {
		if migrator.HasColumn("reports", "reporter_id") {
			return migrator.RenameColumn("reports", "reporter_id", "user_id")
		}
	}
	return nil
}
