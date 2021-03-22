package article

import "time"

//Article product article that available to rent or sell
type Article struct {
	ID        string
	Author    string
	Title     string
	Body      string
	CreatedAt time.Time
}

//NewArticle create new article
func NewArticle(
	author string,
	title string,
	body string,
	createdAt time.Time) Article {

	return Article{
		Author:    author,
		Title:     title,
		Body:      body,
		CreatedAt: createdAt,
	}
}

//ModifyArticle update existing article data
func (oldArticle *Article) ModifyArticle(newAuthor string, newTitle string, newBody string) Article {
	return Article{
		ID:        oldArticle.ID,
		Author:    newAuthor,
		Title:     newTitle,
		Body:      newBody,
		CreatedAt: oldArticle.CreatedAt,
	}
}
