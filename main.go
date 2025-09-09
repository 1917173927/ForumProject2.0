package main

import (
	"api-main/app/controllers"
	"api-main/app/services"
	"api-main/config/database"
	"fmt"
	"log"

	"api-main/config/database/migrations"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db := database.GetDB()

	// Migrations
	if err := migrations.CreateTables(db); err != nil {
		log.Printf("Failed to create tables: %v", err)
		return
	}
	if err := migrations.RenameReporterToUser(db); err != nil {
		log.Printf("Failed to rename reporter_id to user_id: %v", err)
		return
	}

	// Initialize services
	if err := services.InitServices(); err != nil {
		log.Printf("Failed to initialize services: %v", err)
		return
	}

	r := gin.Default()

	// User
	registerCtrl := controllers.RegisterController{}
	r.POST("/api/user/reg", registerCtrl.RegisterHandler)

	loginCtrl := controllers.LoginController{}
	r.POST("/api/user/login", loginCtrl.LoginHandler)

	// Student post
	postCtrl := controllers.PostController{}
	r.POST("/api/student/post", postCtrl.CreatePostHandler)
	r.GET("/api/student/post", postCtrl.GetPostsHandler)
	r.DELETE("/api/student/post", postCtrl.DeletePostHandler)
	r.PUT("/api/student/post", postCtrl.UpdatePostHandler)

	// Report
	reportCtrl := controllers.ReportController{}
	r.POST("/api/student/report-post", reportCtrl.CreateReportHandler)
	r.GET("/api/student/report-post", reportCtrl.GetReportsHandler)
	r.GET("/api/admin/report", reportCtrl.GetPendingReportsHandler)
	r.POST("/api/admin/report", reportCtrl.ReviewReportHandler)

	fmt.Println("Server started at http://localhost:8080")
	r.Run(":8080")
}
