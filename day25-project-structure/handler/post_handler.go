package handler

import (
	"net/http"
	"strconv"

	"go-basics/day25-project-structure/domain"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service domain.PostService
}

func NewPostHandler(service domain.PostService) *PostHandler {
	return &PostHandler{service: service}
}

type createPostInput struct {
	Title   string `json:"title"   binding:"required,min=3,max=200"`
	Content string `json:"content" binding:"required,min=10"`
}

type updatePostInput struct {
	Title     string `json:"title"     binding:"required,min=3,max=200"`
	Content   string `json:"content"   binding:"required,min=10"`
	Published bool   `json:"published"`
}

func optionalUserID(c *gin.Context) *uint {
	raw, exists := c.Get("userID")
	if !exists {
		return nil
	}
	id, ok := raw.(uint)
	if !ok {
		return nil
	}
	return &id
}

func (h *PostHandler) ListPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	result, err := h.service.ListPublished(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *PostHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")
	post, err := h.service.GetBySlug(slug, c.ClientIP(), c.GetHeader("User-Agent"), optionalUserID(c))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) LikePost(c *gin.Context) {
	slug := c.Param("slug")
	if err := h.service.LikePost(slug, c.ClientIP(), c.GetHeader("User-Agent"), optionalUserID(c)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "post liked"})
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	userID, _ := c.Get("userID")
	var input createPostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := h.service.CreatePost(userID.(uint), input.Title, input.Content)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}
	var input updatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := h.service.UpdatePost(uint(id64), input.Title, input.Content, input.Published)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}
	if err := h.service.DeletePost(uint(id64)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}

func (h *PostHandler) GetAnalytics(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}
	summary, err := h.service.GetAnalytics(uint(id64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summary)
}
