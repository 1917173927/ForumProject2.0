package main

import (
	adminController "api-main/internal/admin/controller"
	adminRepository "api-main/internal/admin/repository"
	adminService "api-main/internal/admin/service"
	studentController "api-main/internal/student/controller"
	studentRepository "api-main/internal/student/repository"
	studentService "api-main/internal/student/service"
	userController "api-main/internal/user/controller"
	userService "api-main/internal/user/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dsn := "root:coppklmja!BWZ@tcp(127.0.0.1:3306)/items?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Initialize services
	if err := userService.InitService(); err != nil {
		panic(fmt.Sprintf("Failed to initialize user service: %v", err))
	}
	if err := studentService.InitService(); err != nil {
		panic(fmt.Sprintf("Failed to initialize student service: %v", err))
	}

	// Initialize report repository
	if err := studentRepository.InitDB(); err != nil {
		panic(fmt.Sprintf("Failed to initialize report repository: %v", err))
	}

	// Initialize admin repository
	if err := adminRepository.InitDB(); err != nil {
		panic(fmt.Sprintf("Failed to initialize admin repository: %v", err))
	}

	r := gin.Default()

	// User routes
	userCtrl := userController.LoginController{}
	r.POST("/api/user/login", userCtrl.LoginHandler)

	registerCtrl := userController.RegisterController{}
	r.POST("/api/user/reg", registerCtrl.RegisterHandler)

	// Student routes
	postCtrl := studentController.PostController{}
	r.POST("/api/student/post", postCtrl.CreatePostHandler)
	r.DELETE("/api/student/post", postCtrl.DeletePostHandler)
	r.GET("/api/student/post", postCtrl.GetPostsHandler)
	r.PUT("/api/student/post", postCtrl.UpdatePostHandler)

	// Report routes
	reportCtrl := studentController.ReportController{}
	r.POST("/api/student/report-post", reportCtrl.ReportPostHandler)
	r.GET("/api/student/report-post", reportCtrl.GetReportStatusHandler)

	// Admin report
	reportCtrlAdmin := adminController.ReportController{}
	r.GET("/api/admin/report", reportCtrlAdmin.GetReportsHandler)

	// Admin report approval routes
	adminReportService := &adminService.AdminService{DB: db}
	adminReportController := adminController.AdminReportController{Service: adminReportService}
	r.GET("/api/admin/report", adminReportController.GetPendingReportsHandler)
	r.POST("/api/admin/report", adminReportController.ProcessReportApprovalHandler)

	fmt.Println("Server started at http://localhost:8080")
	r.Run(":8080")
}
