package article

import (
	"kumparan/business/article"
	"kumparan/util"
)

//RepositoryFactory Will return business.article.Repository based on active database connection
func CacheRepositoryFactory(dbCon *util.CacheConnection) article.CacheRepository {
	var articleRepo article.CacheRepository

	articleRepo = NewCacheRepository(dbCon.Redis)

	return articleRepo
}
