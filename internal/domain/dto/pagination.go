package dto

type Paginate struct {
	Limit int `query:"limit" mod:"default=10"`
	Page  int `query:"page" mod:"default=1"`
}

type Pagination struct {
	Limit      int         `json:"limit,omitempty"`
	Page       int         `json:"page,omitempty"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}
