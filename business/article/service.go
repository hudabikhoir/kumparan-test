package article

import (
	"encoding/json"
	"kumparan/business"
	"kumparan/business/article/spec"
	"time"

	validator "github.com/go-playground/validator/v10"
)

//Repository ingoing port for article
type Repository interface {
	//FindArticleByAuthor If no data match with the given tag, will return empty slice instead of nil
	FindArticleByAuthor(Author string) ([]Article, error)

	//FindArticleByQuery If no data match with the given tag, will return empty slice instead of nil
	FindArticleByQuery(Query string) ([]Article, error)

	//FindAllByTag If no data match with the given tag, will return empty slice instead of nil
	FindAllArticle() ([]Article, error)

	//InsertArticle Insert new article into storage
	InsertArticle(article Article) (string, error)

	//FindArticleByID If data not found will return nil without error
	FindArticleByID(ID string) (*Article, error)
}

//CacheRepository ingoing port for article
type CacheRepository interface {
	//SetKey function to set keys in redis
	SetKey(key string, value interface{}, expired int64) bool

	//Get function to get value by key in redis
	Get(key string) string

	//Get function to multiple get value by key in redis
	MGet(keys []string) []Article

	//GetAllKeys function to get all keys
	GetAllKeys() []string
}

//Service outgoing port for article
type Service interface {
	GetArticleByID(ID string) (*Article, error)

	GetArticles(author string, query string) ([]Article, error)

	CreateArticle(upsertarticleSpec spec.UpsertArticleSpec, createdBy string) (string, error)
}

//=============== The implementation of those interface put below =======================

type service struct {
	repository Repository
	cache      CacheRepository
	validate   *validator.Validate
}

//NewService Construct article service object
func NewService(repository Repository, cache CacheRepository) Service {
	return &service{
		repository,
		cache,
		validator.New(),
	}
}

//GetArticleByID Get item by given ID, return nil if not exist
func (s *service) GetArticleByID(ID string) (*Article, error) {
	return s.repository.FindArticleByID(ID)
}

//GetArticles Get all articles by given tag, return zero array if not match
func (s *service) GetArticles(author string, query string) ([]Article, error) {
	var articles []Article
	var err error

	if author != "" {
		articles, err = s.repository.FindArticleByAuthor(author)
	} else if query != "" {
		articles, err = s.repository.FindArticleByQuery(query)
	} else {
		keys := s.cache.GetAllKeys()
		articles = s.cache.MGet(keys)
		// check if there is data on cache, then get from db
		if keys == nil {
			articles, err = s.repository.FindAllArticle()
		}
	}
	if err != nil || articles == nil {
		return []Article{}, err
	}

	return articles, err
}

//CreateArticle Create new article and store into database
func (s *service) CreateArticle(upsertarticleSpec spec.UpsertArticleSpec, createdBy string) (string, error) {
	err := s.validate.Struct(upsertarticleSpec)
	if err != nil {
		return "", business.ErrInvalidSpec
	}

	article := NewArticle(
		upsertarticleSpec.Author,
		upsertarticleSpec.Title,
		upsertarticleSpec.Body,
		time.Now(),
	)

	lid, err := s.repository.InsertArticle(article)
	if err != nil {
		return "", err
	}

	article.ID = lid
	arti, err := json.Marshal(article)
	if err != nil {
		return "", err
	}

	isSuccess := s.cache.SetKey("article:"+lid+":"+upsertarticleSpec.Author, string(arti), 0)
	if !isSuccess {
		return "", err
	}

	return lid, nil
}
