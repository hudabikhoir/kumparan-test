package response

import (
	"kumparan/business/article"
	"time"
)

//GetArticleByIDResponse Get article by ID response payload
type GetArticleByIDResponse struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"ttile"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

//NewGetArticleByIDResponse construct GetArticleByIDResponse
func NewGetArticleByIDResponse(article article.Article) *GetArticleByIDResponse {
	var articleResponse GetArticleByIDResponse
	articleResponse.ID = article.ID
	articleResponse.Author = article.Author
	articleResponse.Title = article.Title
	articleResponse.CreatedAt = article.CreatedAt
	articleResponse.Body = article.Body

	return &articleResponse
}
