package spec

//UpsertArticleSpec create and update item spec
type UpsertArticleSpec struct {
	Author string `validate:"required"`
	Title  string `validate:"required"`
	Body   string `validate:"required"`
}
