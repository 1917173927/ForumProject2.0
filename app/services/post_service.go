package services

import (
	"api-main/app/models"
	"api-main/config/database"
	"errors"
)

type PostRequest struct {
	Content string `json:"Content" gorm:"not null"`
	UserID  uint   `json:"UserID" gorm:"not null"`
}

type PostResponse struct {
	ID      uint   `json:"ID" gorm:"not null"`
	Content string `json:"Content" gorm:"not null"`
	UserID  uint   `json:"UserID" gorm:"not null"`
}

type PostListResponse struct {
	Posts []PostResponse `json:"posts"`
}

func CreatePost(req PostRequest) (PostResponse, error) {
	if req.Content == "" || req.UserID == 0 {
		return PostResponse{}, errors.New("content and user_id are required")
	}

	db := database.GetDB()

	post := models.Post{
		Content: req.Content,
		UserID:  req.UserID,
	}

	if err := db.Create(&post).Error; err != nil {
		return PostResponse{}, errors.New("failed to create post")
	}

	return PostResponse{
		ID:      post.PostID,
		Content: post.Content,
		UserID:  post.UserID,
	}, nil
}

func GetPosts() ([]models.Post, error) {
	db := database.GetDB()

	var posts []models.Post
	if err := db.Find(&posts).Error; err != nil {
		return nil, errors.New("failed to get posts")
	}

	return posts, nil
}

func DeletePost(postID uint, userID uint) error {
	db := database.GetDB()

	var post models.Post
	if err := db.Where("`post_id` = ? AND `user_id` = ?", postID, userID).First(&post).Error; err != nil {
		return errors.New("post not found or unauthorized")
	}

	if err := db.Delete(&post).Error; err != nil {
		return errors.New("failed to delete post")
	}

	return nil
}

func UpdatePost(postID uint, userID uint, content string) (PostResponse, error) {
	if content == "" {
		return PostResponse{}, errors.New("content cannot be empty")
	}

	db := database.GetDB()

	var post models.Post
	if err := db.Where("`post_id` = ? AND `user_id` = ?", postID, userID).First(&post).Error; err != nil {
		return PostResponse{}, errors.New("post not found or unauthorized")
	}

	post.Content = content
	if err := db.Save(&post).Error; err != nil {
		return PostResponse{}, errors.New("failed to update post")
	}

	return PostResponse{
		ID:      post.PostID,
		Content: post.Content,
		UserID:  post.UserID,
	}, nil
}
