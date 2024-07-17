package dto

type SearchCake struct {
	Name string `json:"name"`
	Page int `json:"page"`
	PageSize int `json:"page_size"`
}