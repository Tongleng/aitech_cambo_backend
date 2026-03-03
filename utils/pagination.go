package utils

type PaginationResponse struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	LastPage int         `json:"last_page"`
}
