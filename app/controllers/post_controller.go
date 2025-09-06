package controllers

import (
	"api-main/app/services"
	"api-main/app/utils"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type PostController struct{}

func (pc *PostController) CreatePostHandler(c *gin.Context) {
	var req services.PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSON(c.Writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	resp, err := services.CreatePost(req)
	if err != nil {
		utils.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSON(c.Writer, http.StatusOK, "Post created successfully", resp)
}

func (pc *PostController) GetPostsHandler(c *gin.Context) {
	posts, err := services.GetPosts()
	if err != nil {
		utils.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSON(c.Writer, http.StatusOK, "Success", gin.H{"posts": posts})
}

func (pc *PostController) DeletePostHandler(c *gin.Context) {
	postIDStr := c.Query("post_id")
	if postIDStr == "" {
		utils.JSON(c.Writer, http.StatusBadRequest, "post_id parameter is required", nil)
		return
	}
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		utils.JSON(c.Writer, http.StatusBadRequest, "post_id must be a positive integer", nil)
		return
	}

	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		utils.JSON(c.Writer, http.StatusBadRequest, "user_id parameter is required", nil)
		return
	}
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		utils.JSON(c.Writer, http.StatusBadRequest, "user_id must be a positive integer", nil)
		return
	}

	if err := services.DeletePost(uint(postID), uint(userID)); err != nil {
		utils.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSON(c.Writer, http.StatusOK, "Post deleted successfully", nil)
}

func (pc *PostController) UpdatePostHandler(c *gin.Context) {
	var req struct {
		PostID  uint   `json:"post_id"`
		UserID  uint   `json:"user_id"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSON(c.Writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	if req.PostID == 0 {
		utils.JSON(c.Writer, http.StatusBadRequest, "post_id is required", nil)
		return
	}

	if req.UserID == 0 {
		utils.JSON(c.Writer, http.StatusBadRequest, "user_id is required", nil)
		return
	}

	if req.Content == "" {
		utils.JSON(c.Writer, http.StatusBadRequest, "content cannot be empty", nil)
		return
	}

	resp, err := services.UpdatePost(req.PostID, req.UserID, req.Content)
	if err != nil {
		utils.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSON(c.Writer, http.StatusOK, "Post updated successfully", resp)
}
