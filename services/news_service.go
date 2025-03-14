package services

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/ethan-t-hansen/codex-backend/models"
)

type NewsService struct {
	client    *resty.Client
	apiKey    string
	baseURL   string
}

func NewNewsService(apiKey string) *NewsService {
	return &NewsService{
		client:  resty.New(),
		apiKey:  apiKey,
		baseURL: "https://newsapi.org/v2",
	}
}

func (s *NewsService) GetTopHeadlines(params models.NewsRequestParams) (*models.NewsResponse, error) {

	params.APIKey = s.apiKey

	queryParams := url.Values{}
	if params.Query != "" {
		queryParams.Add("q", params.Query)
	}
	if params.Country != "" {
		queryParams.Add("country", params.Country)
	}
	if params.Category != "" {
		queryParams.Add("category", params.Category)
	}
	if params.PageSize > 0 {
		queryParams.Add("pageSize", fmt.Sprintf("%d", params.PageSize))
	}
	if params.Page > 0 {
		queryParams.Add("page", fmt.Sprintf("%d", params.Page))
	}
	if params.SortBy != "" {
		queryParams.Add("sortBy", params.SortBy)
	}
	queryParams.Add("apiKey", params.APIKey)

	resp, err := s.client.R().
		SetQueryParamsFromValues(queryParams).
		Get(s.baseURL + "/top-headlines")

	if err != nil {
		return nil, err
	}

	var newsResp models.NewsResponse
	if err := json.Unmarshal(resp.Body(), &newsResp); err != nil {
		return nil, err
	}

	return &newsResp, nil
}

func (s *NewsService) SearchNews(params models.NewsRequestParams) (*models.NewsResponse, error) {

	params.APIKey = s.apiKey

	queryParams := url.Values{}
	if params.Query != "" {
		queryParams.Add("q", params.Query)
	}
	if params.PageSize > 0 {
		queryParams.Add("pageSize", fmt.Sprintf("%d", params.PageSize))
	}
	if params.Page > 0 {
		queryParams.Add("page", fmt.Sprintf("%d", params.Page))
	}
	if params.SortBy != "" {
		queryParams.Add("sortBy", params.SortBy)
	}
	queryParams.Add("apiKey", params.APIKey)

	resp, err := s.client.R().
		SetQueryParamsFromValues(queryParams).
		Get(s.baseURL + "/everything")

	if err != nil {
		return nil, err
	}

	var newsResp models.NewsResponse
	if err := json.Unmarshal(resp.Body(), &newsResp); err != nil {
		return nil, err
	}

	return &newsResp, nil
}