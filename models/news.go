package models

import "time"

type Article struct {
	Source      Source    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type NewsRequestParams struct {
	Query      string `form:"q"`
	Country    string `form:"country"`
	Category   string `form:"category"`
	PageSize   int    `form:"pageSize"`
	Page       int    `form:"page"`
	SortBy     string `form:"sortBy"`
	APIKey     string
}