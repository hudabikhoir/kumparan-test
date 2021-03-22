package http

import (
	"kumparan/api/v1/article"

	"github.com/labstack/echo"
)

//RegisterPath Registera V1 API path
func RegisterPath(e *echo.Echo, articleController *article.Controller) {
	if articleController == nil {
		panic("article controller cannot be nil")
	}

	//article
	articleV1 := e.Group("v1/articles")
	articleV1.GET("", articleController.GetArticles)
	articleV1.POST("", articleController.CreateNewArticle)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
