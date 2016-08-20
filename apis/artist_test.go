package apis

import (
	"net/http"
	"testing"

	"github.com/qiangxue/golang-restful-starter-kit/daos"
	"github.com/qiangxue/golang-restful-starter-kit/services"
	"github.com/qiangxue/golang-restful-starter-kit/testdata"
)

func TestArtist(t *testing.T) {
	testdata.ResetDB()
	ServeArtistResource(&router.RouteGroup, services.NewArtistService(daos.NewArtistDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	nameRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"Name","error":"cannot be blank"}]}`

	runAPITests(t, []apiTestCase{
		{"t1 - get an artist", "GET", "/artists/2", "", http.StatusOK, `{"id":2,"name":"Accept"}`},
		{"t2 - get a nonexisting artist", "GET", "/artists/99999", "", http.StatusNotFound, notFoundError},
		{"t3 - create an artist", "POST", "/artists", `{"name":"Qiang"}`, http.StatusOK, `{"id": 276, "name":"Qiang"}`},
		{"t4 - create an artist with validation error", "POST", "/artists", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t5 - update an artist", "PUT", "/artists/2", `{"name":"Qiang"}`, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t6 - update an artist with validation error", "PUT", "/artists/2", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t7 - update a nonexisting artist", "PUT", "/artists/99999", "{}", http.StatusNotFound, notFoundError},
		{"t8 - delete an artist", "DELETE", "/artists/2", ``, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t9 - delete a nonexisting artist", "DELETE", "/artists/99999", "", http.StatusNotFound, notFoundError},
		{"t10 - get a list of artists", "GET", "/artists?page=3&per_page=2", "", http.StatusOK, `{"page":3,"per_page":2,"page_count":138,"total_count":275,"items":[{"id":6,"name":"Antônio Carlos Jobim"},{"id":7,"name":"Apocalyptica"}]}`},
	})
}
