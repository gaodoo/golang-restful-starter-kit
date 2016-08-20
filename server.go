package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	_ "github.com/lib/pq"
	"github.com/qiangxue/golang-restful-starter-kit/apis"
	"github.com/qiangxue/golang-restful-starter-kit/app"
	"github.com/qiangxue/golang-restful-starter-kit/daos"
	"github.com/qiangxue/golang-restful-starter-kit/errors"
	"github.com/qiangxue/golang-restful-starter-kit/services"
	"github.com/go-ozzo/ozzo-routing/auth"
)

func main() {
	// load application configurations
	config, err := app.LoadConfig("./config")
	if err != nil {
		panic(err)
	}

	// load error messages
	if err := errors.LoadMessages(config.GetString("error_file")); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// create the logger
	logger := logrus.New()

	// connect to the database
	db, err := dbx.MustOpen("postgres", config.GetString("dsn"))
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	// wire up API routing
	http.Handle("/", buildRouter(config, logger, db))

	// start the server
	address := ":" + config.GetString("server_port")
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(config app.Config, logger *logrus.Logger, db *dbx.DB) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		return c.Write("OK " + app.Version)
	})

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		app.Transactional(db),
	)

	rg := router.Group("/v1")

	// Uncomment the following lines to enable JWT authentication
	// rg.Post("/auth", apis.Auth(config.GetString("jwt_signing_key")))
	// rg.Use(auth.JWT(config.GetString("jwt_verification_key")))

	artistDAO := daos.NewArtistDAO()
	apis.ServeArtistResource(rg, services.NewArtistService(artistDAO))

	// wire up more resource APIs here

	return router
}
