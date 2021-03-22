package article

import (
	"kumparan/api/common"
	"kumparan/api/v1/article/request"
	"kumparan/api/v1/article/response"
	"kumparan/business"
	articleBusiness "kumparan/business/article"
	"net/http"

	v10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

//Controller Get article API controller
type Controller struct {
	service   articleBusiness.Service
	validator *v10.Validate
}

//NewController Construct article API controller
func NewController(service articleBusiness.Service) *Controller {
	return &Controller{
		service,
		v10.New(),
	}
}

//GetArticles Get article by ID echo handler
func (controller *Controller) GetArticles(c echo.Context) error {
	author := c.QueryParam("author")
	query := c.QueryParam("query")

	articles, err := controller.service.GetArticles(author, query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewGetArticleByTagResponse(articles)
	return c.JSON(http.StatusOK, response)
}

//CreateNewArticle Create new article echo handler
func (controller *Controller) CreateNewArticle(c echo.Context) error {
	createArticleRequest := new(request.CreateArticleRequest)

	if err := c.Bind(createArticleRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	ID, err := controller.service.CreateArticle(*createArticleRequest.ToUpsertArticleSpec(), "creator")
	if err != nil {
		if err == business.ErrInvalidSpec {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewCreateNewArticleResponse(ID)
	return c.JSON(http.StatusCreated, response)
}
