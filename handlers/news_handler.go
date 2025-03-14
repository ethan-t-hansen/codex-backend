package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ethan-t-hansen/codex-backend/models"
	"github.com/ethan-t-hansen/codex-backend/services"
)

type NewsHandler struct {
	newsService *services.NewsService
}

func NewNewsHandler(newsService *services.NewsService) *NewsHandler {
	return &NewsHandler{
		newsService: newsService,
	}
}

func (h *NewsHandler) GetTopHeadlines(c *gin.Context) {
	params := models.NewsRequestParams{
		Query:    c.Query("q"),
		Country:  c.Query("country"),
		Category: c.Query("category"),
		SortBy:   c.Query("sortBy"),
	}

	// Parse page and pageSize if present
	if pageSize := c.Query("pageSize"); pageSize != "" {
		size, err := strconv.Atoi(pageSize)
		if err == nil {
			params.PageSize = size
		}
	}

	if page := c.Query("page"); page != "" {
		pageNum, err := strconv.Atoi(page)
		if err == nil {
			params.Page = pageNum
		}
	}

	// Get news from service
	news, err := h.newsService.GetTopHeadlines(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}

func (h *NewsHandler) SearchNews(c *gin.Context) {
	params := models.NewsRequestParams{
		Query:  c.Query("q"),
		SortBy: c.Query("sortBy"),
	}

	// Parse page and pageSize if present
	if pageSize := c.Query("pageSize"); pageSize != "" {
		size, err := strconv.Atoi(pageSize)
		if err == nil {
			params.PageSize = size
		}
	}

	if page := c.Query("page"); page != "" {
		pageNum, err := strconv.Atoi(page)
		if err == nil {
			params.Page = pageNum
		}
	}

	// Get news from service
	news, err := h.newsService.SearchNews(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}