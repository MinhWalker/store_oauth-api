package app

import (
	"github.com/MinhWalker/store_oauth-api/src/http"
	"github.com/MinhWalker/store_oauth-api/src/repository/db"
	"github.com/MinhWalker/store_oauth-api/src/repository/rest"
	"github.com/MinhWalker/store_oauth-api/src/services"
)

func mapUrls()  {
	atHandler := http.NewHandler(services.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
}