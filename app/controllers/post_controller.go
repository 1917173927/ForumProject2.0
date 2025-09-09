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
		utils.JsonErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	resp, err := services.CreatePost(req)
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonSuccessResponse(c, resp)
}

func (pc *PostController) GetPostsHandler(c *gin.Context) {
	posts, err := services.GetPosts()
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonSuccessResponse(c, gin.H{"posts": posts})
}

func (pc *PostController) DeletePostHandler(c *gin.Context) {
	postIDStr := c.Query("post_id")
	if postIDStr == "" {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "post_id parameter is required")
		return
	}
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "post_id must be a positive integer")
		return
	}

	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "user_id parameter is required")
		return
	}
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "user_id must be a positive integer")
		return
	}

	if err := services.DeletePost(uint(postID), uint(userID)); err != nil {
		utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

func (pc *PostController) UpdatePostHandler(c *gin.Context) {
	var req struct {
		PostID  uint   `json:"PostID"`
		UserID  uint   `json:"UserID"`
		Content string `json:"Content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.PostID == 0 {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "post_id is required")
		return
	}

	if req.UserID == 0 {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "user_id is required")
		return
	}

	if req.Content == "" {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "content cannot be empty")
		return
	}

	resp, err := services.UpdatePost(req.PostID, req.UserID, req.Content)
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonSuccessResponse(c, resp)
}
