package response

//CreateNewArticleResponse Create article response payload
type CreateNewArticleResponse struct {
	ID string `json:"id"`
}

//NewCreateNewArticleResponse construct CreateNewArticleResponse
func NewCreateNewArticleResponse(id string) *CreateNewArticleResponse {
	return &CreateNewArticleResponse{
		id,
	}
}
