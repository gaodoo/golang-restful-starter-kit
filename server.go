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
)

func main() {
	config, err := app.LoadConfig("./config")
	if err != nil {
		panic(err)
	}

	if err := errors.LoadMessages(config.GetString("error_file")); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	logger := logrus.New()

	db, err := dbx.MustOpen("postgres", config.GetString("dsn"))
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	r := routing.New()
	r.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		return c.Write("OK " + app.Version)
	})
	r.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		app.Transactional(db),
	)

	serveResources(r.Group("/v1"))

	http.Handle("/", r)

	address := ":" + config.GetString("server_port")
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

// serveResources wires up the DAOs, services, and APIs.
// This is the central place for injecting dependencies.
func serveResources(rg *routing.RouteGroup) {
	artistDAO := daos.NewArtistDAO()
	apis.ServeArtist(rg, services.NewArtistService(artistDAO))
}
