package article

import (
	"context"
	"kumparan/business/article"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoDBRepository The implementation of article.Repository object
type MongoDBRepository struct {
	col *mongo.Collection
}

type collection struct {
	ID        primitive.ObjectID `bson:"_id"`
	Author    string             `bson:"author"`
	Title     string             `bson:"title"`
	Body      string             `bson:"body"`
	CreatedAt time.Time          `bson:"created_at"`
}

func newCollection(article article.Article) (*collection, error) {
	objectID := primitive.NewObjectID()

	return &collection{
		objectID,
		article.Author,
		article.Title,
		article.Body,
		article.CreatedAt,
	}, nil
}

func (col *collection) ToArticle() article.Article {
	var article article.Article
	article.ID = col.ID.Hex()
	article.Title = col.Title
	article.Author = col.Author
	article.Body = col.Body
	article.CreatedAt = col.CreatedAt

	return article
}

//NewMongoDBRepository Generate mongo DB article repository
func NewMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		db.Collection("articles"),
	}
}

//FindArticleByID Find article based on given ID. Its return nil if not found
func (repo *MongoDBRepository) FindArticleByID(ID string) (*article.Article, error) {
	var col collection

	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		//if cannot be convert means that ID will be never found
		return nil, nil
	}

	filter := bson.M{
		"_id": objectID,
	}

	if err := repo.col.FindOne(context.TODO(), filter).Decode(&col); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	article := col.ToArticle()
	return &article, nil
}

//FindArticleByAuthor Find all articles based on given tag. Its return empty array if not found
func (repo *MongoDBRepository) FindArticleByAuthor(author string) ([]article.Article, error) {
	filter := bson.M{
		"author": author,
	}

	cursor, err := repo.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var articles []article.Article

	for cursor.Next(context.TODO()) {
		var col collection
		if err = cursor.Decode(&col); err != nil {
			return nil, err
		}

		articles = append(articles, col.ToArticle())
	}

	return articles, nil
}

//FindArticleByQuery Find all articles based on given tag. Its return empty array if not found
func (repo *MongoDBRepository) FindArticleByQuery(query string) ([]article.Article, error) {
	filter := bson.M{"$or": []bson.M{
		{"title": primitive.Regex{
			Pattern: ".*" + query + ".*", Options: "i"}},
		{"body": primitive.Regex{
			Pattern: ".*" + query + ".*", Options: "i"}},
	}}

	cursor, err := repo.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var articles []article.Article

	for cursor.Next(context.TODO()) {
		var col collection
		if err = cursor.Decode(&col); err != nil {
			return nil, err
		}

		articles = append(articles, col.ToArticle())
	}

	return articles, nil
}

//FindAllArticle Find all articles based on given tag. Its return empty array if not found
func (repo *MongoDBRepository) FindAllArticle() ([]article.Article, error) {
	filter := bson.M{}

	cursor, err := repo.col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var articles []article.Article

	for cursor.Next(context.TODO()) {
		var col collection
		if err = cursor.Decode(&col); err != nil {
			return nil, err
		}

		articles = append(articles, col.ToArticle())
	}

	return articles, nil
}

//InsertArticle Insert new article into database. Its return article id if success
func (repo *MongoDBRepository) InsertArticle(article article.Article) (string, error) {
	col, err := newCollection(article)
	if err != nil {
		return col.ID.String(), err
	}

	_, err = repo.col.InsertOne(context.TODO(), col)

	if err != nil {
		return col.ID.String(), err
	}

	return col.ID.String(), nil
}
