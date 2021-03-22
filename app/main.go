package main

import (
	"context"
	"fmt"
	api "kumparan/api"
	articleControllerV1 "kumparan/api/v1/article"
	businessArticle "kumparan/business/article"
	"kumparan/config"
	articleCacheRepo "kumparan/modules/cache/article"
	articleRepo "kumparan/modules/repository/article"
	"kumparan/util"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	dbCon := util.NewDatabaseConnection(config)

	//initiate article repository
	articleRepo := articleRepo.RepositoryFactory(dbCon)

	//initialize database connection based on given config
	dbCacheCon := util.NewCacheConnection(config)

	//initiate article cache repository
	articleCacheRepo := articleCacheRepo.CacheRepositoryFactory(dbCacheCon)

	//initiate article service
	articleService := businessArticle.NewService(articleRepo, articleCacheRepo)

	//initiate article controller
	articleControllerV1 := articleControllerV1.NewController(articleService)

	//create echo http
	e := echo.New()

	//register API path and handler
	api.RegisterPath(e, articleControllerV1)

	// run server
	go func() {
		address := fmt.Sprintf("localhost:%d", config.Port)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	//close db
	defer dbCon.CloseConnection()

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
