package article

import (
	"kumparan/business/article"
	"kumparan/util"
)

//RepositoryFactory Will return business.article.Repository based on active database connection
func RepositoryFactory(dbCon *util.DatabaseConnection) article.Repository {
	var articleRepo article.Repository

	// articleRepo = NewMySQLRepository(dbCon.MySQLDB)
	if dbCon.Driver == util.MySQL {
		articleRepo = NewMySQLRepository(dbCon.MySQLDB)
	} else if dbCon.Driver == util.MongoDB {
		articleRepo = NewMongoDBRepository(dbCon.MongoDB)
	}

	return articleRepo
}
