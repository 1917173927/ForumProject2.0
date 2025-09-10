package controllers

import (
	"api-main/app/services"
	"api-main/app/utils"
	"net/http"

	"strconv"
	"github.com/gin-gonic/gin"
)

type PostController struct{}

// @Summary Create a post
// @Description Create a new post
// @Tags post
// @Accept json
// @Produce json
// @Param request body services.PostRequest true "Post request"
// @Success 200 {object} services.PostResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /post [post]
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

// @Summary Get all posts
// @Description Get a list of all posts
// @Tags post
// @Accept json
// @Produce json
// @Success 200 {object} services.PostListResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /posts [get]
func (pc *PostController) GetPostsHandler(c *gin.Context) {
	posts, err := services.GetPosts()
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonSuccessResponse(c, gin.H{"posts": posts})
}

// @Summary Delete a post
// @Description Delete a post by ID
// @Tags post
// @Accept json
// @Produce json
// @Param id query string true "Post ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /post [delete]
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

// @Summary Update a post
// @Description Update a post by ID
// @Tags post
// @Accept json
// @Produce json
// @Param request body services.PostRequest true "Post request"
// @Param id query string true "Post ID"
// @Success 200 {object} services.PostResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /post [put]
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
