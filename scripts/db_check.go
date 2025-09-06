package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"api-main/config/database"
)

func main() {
	// 获取数据库连接
	db, err := database.GetDBConnection()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	fmt.Println("数据库连接成功")

	// 检查reports表是否存在
	if !db.Migrator().HasTable("reports") {
		log.Fatal("reports表不存在")
	}
	fmt.Println("reports表存在")

	// 查询reports表中的记录数
	var count int64
	if err := db.Model(&Report{}).Count(&count).Error; err != nil {
		log.Fatalf("查询reports表失败: %v", err)
	}
	fmt.Printf("reports表中有 %d 条记录\n", count)

	// 如果有记录，打印前5条
	if count > 0 {
		var reports []Report
		if err := db.Limit(5).Find(&reports).Error; err != nil {
			log.Fatalf("查询reports记录失败: %v", err)
		}
		fmt.Println("前5条举报记录:")
		for i, r := range reports {
			fmt.Printf("%d. ID:%d PostID:%d ReporterID:%d Reason:%s\n", 
				i+1, r.ID, r.PostID, r.ReporterID, r.Reason)
		}
	}
}

type Report struct {
	gorm.Model
	PostID     uint   `gorm:"not null"`
	ReporterID uint   `gorm:"not null"`
	Reason     string `gorm:"type:text;not null"`
	Status     string `gorm:"type:varchar(20);default:'pending'"`
	Type       string `gorm:"type:varchar(20)"`
}
