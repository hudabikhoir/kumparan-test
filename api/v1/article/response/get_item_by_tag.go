package response

import "kumparan/business/article"

//GetArticleByTagResponse Get article by tag response payload
type GetArticleByTagResponse struct {
	Articles []*GetArticleByIDResponse `json:"articles"`
}

//NewGetArticleByTagResponse construct GetArticleByTagResponse
func NewGetArticleByTagResponse(articles []article.Article) *GetArticleByTagResponse {
	var articleResponses []*GetArticleByIDResponse
	articleResponses = make([]*GetArticleByIDResponse, 0)

	for _, article := range articles {
		articleResponses = append(articleResponses, NewGetArticleByIDResponse(article))
	}

	return &GetArticleByTagResponse{
		articleResponses,
	}
}
