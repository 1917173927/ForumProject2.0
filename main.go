package main

import (
	"api-main/app/controllers"
	"api-main/app/services"
	"api-main/config/database"
	"fmt"
	"api-main/config/database/migrations"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db, err := database.GetDBConnection()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database: %v", err))
	}

	// Run migrations
	if err := migrations.CreateReportsTable(db); err != nil {
		panic(fmt.Sprintf("Failed to create reports table: %v", err))
	}
	if err := migrations.RenameReporterToUser(db); err != nil {
		panic(fmt.Sprintf("Failed to rename reporter_id to user_id: %v", err))
	}

	// Initialize services
	if err := services.InitServices(); err != nil {
		panic(fmt.Sprintf("Failed to initialize services: %v", err))
	}

	r := gin.Default()

	// User routes
	registerCtrl := controllers.RegisterController{}
	r.POST("/api/user/reg", registerCtrl.RegisterHandler)

	loginCtrl := controllers.LoginController{}
	r.POST("/api/user/login", loginCtrl.LoginHandler)

	// Student post routes
	postCtrl := controllers.PostController{}
	r.POST("/api/student/post", postCtrl.CreatePostHandler)
	r.GET("/api/student/post", postCtrl.GetPostsHandler)
	r.DELETE("/api/student/post", postCtrl.DeletePostHandler)
	r.PUT("/api/student/post", postCtrl.UpdatePostHandler)

	// Report routes
	reportCtrl := controllers.ReportController{}
	r.POST("/api/student/report-post", reportCtrl.CreateReportHandler)
	r.GET("/api/student/report-post", reportCtrl.GetReportsHandler)
	r.GET("/api/admin/report", reportCtrl.GetPendingReportsHandler)
	r.POST("/api/admin/report", reportCtrl.ReviewReportHandler)

	// TODO: Add admin routes after refactoring

	fmt.Println("Server started at http://localhost:8080")
	r.Run(":8080")
}
