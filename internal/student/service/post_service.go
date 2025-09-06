package service

import (
	"errors"
	"log"
	"strconv"
	"time"

	"api-main/internal/student/repository"

	"gorm.io/gorm"
)

type PostRequest struct {
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

type PostResponse struct {
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	Content   string    `gorm:"column:content"`
	UserID    int       `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

type UpdatePostRequest struct {
	UserID  int    `json:"user_id"`
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
}

var db *gorm.DB

func InitService() error {
	var err error
	db, err = repository.GetDBConnection()
	if err != nil {
		return errors.New("failed to initialize database connection")
	}
	return nil
}

func CreatePost(req PostRequest) (PostResponse, error) {
	if db == nil {
		return PostResponse{}, errors.New("database connection is not initialized")
	}

	if req.Content == "" {
		return PostResponse{}, errors.New("content is required")
	}
	if req.UserID <= 0 {
		return PostResponse{}, errors.New("userID must be a positive number")
	}

	post := Post{
		Content:   req.Content,
		UserID:    req.UserID,
		CreatedAt: time.Now(),
	}

	if err := db.Create(&post).Error; err != nil {
		log.Printf("Failed to save post: %v", err)
		return PostResponse{}, errors.New("failed to save post")
	}

	return PostResponse{
		Content:   post.Content,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt,
	}, nil
}

func DeletePost(userID string, postID string) error {
	db, err := repository.GetDBConnection()
	if err != nil {
		return errors.New("database connection error")
	}

	// 查询帖子是否存在
	var post repository.Post
	if err := db.Where("post_id = ?", postID).First(&post).Error; err != nil {
		return errors.New("post not found")
	}

	// 验证用户权限
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return errors.New("invalid user_id format")
	}
	if post.UserID != userIDInt {
		return errors.New("unauthorized to delete this post")
	}

	// 删除帖子
	if err := db.Where("post_id = ?", postID).Delete(&repository.Post{}).Error; err != nil {
		return errors.New("failed to delete post")
	}

	return nil
}

func UpdatePost(req UpdatePostRequest) error {
	db, err := repository.GetDBConnection()
	if err != nil {
		return errors.New("database connection error")
	}

	// 查询帖子是否存在且用户权限验证
	var post repository.Post
	if err := db.Where("post_id = ?", req.PostID).First(&post).Error; err != nil {
		return errors.New("post not found")
	}

	if post.UserID != req.UserID {
		return errors.New("unauthorized")
	}

	// 更新帖子内容
	post.Content = req.Content
	if err := db.Save(&post).Error; err != nil {
		return errors.New("failed to update post")
	}

	return nil
}

func GetPosts() ([]repository.Post, error) {
	db, err := repository.GetDBConnection()
	if err != nil {
		return nil, errors.New("database connection error")
	}

	var posts []repository.Post
	if err := db.Find(&posts).Error; err != nil {
		return nil, errors.New("failed to retrieve posts")
	}

	return posts, nil
}
