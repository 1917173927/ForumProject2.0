package controller

import (
	"api-main/internal/common/response"
	"api-main/internal/student/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func (pc *PostController) CreatePostHandler(c *gin.Context) {
	var req service.PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSON(c.Writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	resp, err := service.CreatePost(req)
	if err != nil {
		response.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.JSON(c.Writer, http.StatusOK, "Success", resp)
}

func (pc *PostController) DeletePostHandler(c *gin.Context) {
	userID := c.Query("user_id")
	postID := c.Query("post_id")

	if userID == "" || postID == "" {
		response.JSON(c.Writer, http.StatusBadRequest, "Missing user_id or post_id", nil)
		return
	}

	err := service.DeletePost(userID, postID)
	if err != nil {
		response.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.JSON(c.Writer, http.StatusOK, "Post deleted successfully", nil)
}

func (pc *PostController) GetPostsHandler(c *gin.Context) {
	posts, err := service.GetPosts()
	if err != nil {
		response.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.JSON(c.Writer, http.StatusOK, "Success", gin.H{"post_list": posts})
}

func (pc *PostController) UpdatePostHandler(c *gin.Context) {
	var req service.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSON(c.Writer, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	err := service.UpdatePost(req)
	if err != nil {
		response.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.JSON(c.Writer, http.StatusOK, "Post updated successfully", nil)
}
