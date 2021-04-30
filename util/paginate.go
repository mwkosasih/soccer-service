package util

var TotalResultDefault = 10

type PaginationV4 struct {
	PageSize     int64 `json:"page_size"`
	CurrentPage  int64 `json:"current_page"`
	TotalPage    int64 `json:"total_page"`
	TotalResults int64 `json:"total_results"`
}

func NewPaginationV4(pageSize int64, CurrentPage int64) *PaginationV4 {
	return &PaginationV4{
		PageSize:     pageSize,
		CurrentPage:  CurrentPage,
		TotalPage:    0,
		TotalResults: 0,
	}
}
