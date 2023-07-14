package usecase

type ArticlesResponse struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdateAt  string `json:"updatedAt"`
}

type ArticlesRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

type Pagination struct {
	Limit  int
	Offset int
	Status string
}
