package request

import "kumparan/business/article/spec"

//CreateArticleRequest create article request payload
type CreateArticleRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//ToUpsertArticleSpec convert into article.UpsertArticleSpec object
func (req *CreateArticleRequest) ToUpsertArticleSpec() *spec.UpsertArticleSpec {
	var upsertArticleSpec spec.UpsertArticleSpec
	upsertArticleSpec.Author = req.Author
	upsertArticleSpec.Title = req.Title
	upsertArticleSpec.Body = req.Body

	return &upsertArticleSpec
}
