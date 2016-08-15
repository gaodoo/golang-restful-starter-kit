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
	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("music")
	viper.BindEnv("dsn")
	viper.BindEnv("server_port")
	viper.BindEnv("error_file")

	viper.SetDefault("error_file", "config/errors.yaml")
	viper.SetDefault("dsn", "postgres://music:music@127.0.0.1:5432/music?sslmode=disable")
	viper.SetDefault("server_port", "8080")

	if err := errors.LoadMessages(viper.GetString("error_file")); err != nil {
		panic(fmt.Sprint("Failed to read the error message file: ", err))
	}

	logger := logrus.New()
	db, err := dbx.MustOpen("postgres", viper.GetString("dsn"))
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	r := routing.New()
	r.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		return c.Write("OK v" + app.Version)
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

	address := ":" + viper.GetString("server_port")
	logger.Infof("server v%v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

// serveResources wires up the DAOs, services, and APIs.
// This is the central place for injecting dependencies.
func serveResources(rg *routing.RouteGroup) {
	artistDAO := daos.NewArtistDAO()
	apis.ServeArtist(rg, services.NewArtistService(artistDAO))
}
