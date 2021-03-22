package article

import (
	"database/sql"
	"strconv"

	"kumparan/business/article"
)

//MySQLRepository The implementation of article.Repository object
type MySQLRepository struct {
	db *sql.DB
}

//NewMySQLRepository Generate mongo DB article repository
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db,
	}
}

//FindArticleByID Find article based on given ID. Its return nil if not found
func (repo *MySQLRepository) FindArticleByID(ID string) (*article.Article, error) {
	var article article.Article

	selectQuery := `SELECT id, author, title, body, created_at
		FROM articles i
		WHERE i.id = ?`

	err := repo.db.
		QueryRow(selectQuery, ID).
		Scan(
			&article.ID, &article.Author, &article.Title,
			&article.Body, &article.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &article, nil
}

//FindArticleByAuthor Find article based on given Author. Its return nil if not found
func (repo *MySQLRepository) FindArticleByAuthor(Author string) ([]article.Article, error) {
	selectQuery := `SELECT id, author, title, body, created_at
	FROM articles i
	WHERE i.author = ?`

	row, err := repo.db.Query(selectQuery, Author)
	if err != nil {
		return nil, err
	}

	defer row.Close()
	var articles []article.Article

	for row.Next() {
		var article article.Article

		err := row.Scan(
			&article.ID, &article.Author, &article.Title,
			&article.Body, &article.CreatedAt)

		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	if err != nil {
		return nil, err
	}

	return articles, nil
}

//FindArticleByQuery Find article based on given Author. Its return nil if not found
func (repo *MySQLRepository) FindArticleByQuery(Query string) ([]article.Article, error) {
	selectQuery := `SELECT id, author, title, body, created_at
		FROM articles i
		WHERE i.title like ?
		OR i.body like ?`

	row, err := repo.db.Query(selectQuery, Query, Query)
	if err != nil {
		return nil, err
	}

	defer row.Close()
	var articles []article.Article

	for row.Next() {
		var article article.Article

		err := row.Scan(
			&article.ID, &article.Author, &article.Title,
			&article.Body, &article.CreatedAt)

		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	if err != nil {
		return nil, err
	}

	return articles, nil
}

//FindAllArticle Find all articles based on given tag. Its return empty array if not found
func (repo *MySQLRepository) FindAllArticle() ([]article.Article, error) {
	//TODO: if feel have a performance issue in tag grouping, move the logic from db to here
	selectQuery := `SELECT id, author, title, body, created_at
		FROM articles i`

	row, err := repo.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var articles []article.Article

	for row.Next() {
		var article article.Article

		err := row.Scan(
			&article.ID, &article.Author, &article.Title,
			&article.Body, &article.CreatedAt)

		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	if err != nil {
		return nil, err
	}

	return articles, nil
}

//InsertArticle Insert new article into database. Its return article id if success
func (repo *MySQLRepository) InsertArticle(article article.Article) (string, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return "", err
	}

	articleQuery := `INSERT INTO articles (
			author, 
			title, 
			body, 
			created_at
		) VALUES (?, ?, ?, ?)`

	if err != nil {
		return "", err
	}

	res, err := tx.Exec(articleQuery,
		article.Author,
		article.Title,
		article.Body,
		article.CreatedAt,
	)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	lid, err := res.LastInsertId()

	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = tx.Commit()

	if err != nil {
		return "", err
	}

	lidString := strconv.Itoa(int(lid))
	return lidString, nil
}
