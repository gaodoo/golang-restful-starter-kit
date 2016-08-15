# Golang (Go) RESTful Application Starter Kit (Boilerplate)

This starter kit is designed to get you up and running with a project structure optimal for developing
RESTful services in Go. Using some of the best available Go packages and tools, the kit implements
a database-driven RESTful service that supports the typical CRUD operations of a database table.
The project structure, however, is also suitable for other kinds of RESTful services.

The kit uses the following packages:

* Routing framework: [ozzo-routing](https://github.com/go-ozzo/ozzo-routing)
* Database: [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)
* Data validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* Logging: [logrus](https://github.com/Sirupsen/logrus)
* Configuration: [viper](https://github.com/spf13/viper)
* Dependency management: [glide](https://github.com/Masterminds/glide)
* Testing: [testify](https://github.com/stretchr/testify)


## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires Go 1.5 or above.

After installing Go, run the following commands to download and install this starter kit:

```shell
# install the starter kit
go get github.com/qiangxue/golang-restful-starter-kit

# install glide (if you don't have it), a dependency management tool
go get -u github.com/Masterminds/glide

# fetch the dependent packages
cd $GOPATH/qiangxue/golang-restful-starter-kit
glide install -u -s
```

At this stage, you have the source code of the starter kit together with all its dependent packages under
the directory `$GOPATH/qiangxue/golang-restful-starter-kit`.

Next, create a PostgreSQL database using the SQL statements in the file `data/db.sql`. By default, the application
will attempt to connect to the database specified by the DSN `postgres://music:music@127.0.0.1:5432/music?sslmode=disable`,
which specifies that a PostgreSQL server is running on the local machine at port 5432 and there is a database named `music`
accessible with `music` being used as both the username and password. You may change the default DSN by setting
an environment variable named `MUSIC_DSN` before running the application.

Now you can build and run the application:

```shell
cd $GOPATH/qiangxue/golang-restful-starter-kit
go run server.go
```

This starts an HTTP server at port 8080, and you can try the following RESTful endpoints:

* `GET /ping`: a ping service mainly provided for health check purpose
* `GET /v1/artists`: returns a paginated list of the artists
* `GET /v1/artists/:id`: returns the detailed information of an artist
* `POST /v1/artists`: creates a new artist
* `PUT /v1/artists/:id`: updates an existing artist
* `DELETE /v1/artists/:id`: deletes an artist

For example, if you access the URL `http://localhost:8080/v1/artists` in a browser or via a `cURL` command:

```shell
curl -X GET "http://localhost:8080/v1/artists"
```

you should be able to receive a list of the artists in JSON format in the response.

## Project Dissection

TBD
