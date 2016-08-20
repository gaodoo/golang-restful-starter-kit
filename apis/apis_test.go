package apis

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/qiangxue/golang-restful-starter-kit/app"
	"github.com/qiangxue/golang-restful-starter-kit/testdata"
	"github.com/stretchr/testify/assert"
)

type apiTestCase struct {
	tag      string
	method   string
	url      string
	body     string
	status   int
	response string
}

var router *routing.Router

func init() {
	logger := logrus.New()
	logger.Level = logrus.PanicLevel

	router = routing.New()

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		app.Transactional(testdata.DB),
	)
}

func testAPI(method, URL, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, URL, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res
}

func runAPITests(t *testing.T, tests []apiTestCase) {
	for _, test := range tests {
		res := testAPI(test.method, test.url, test.body)
		assert.Equal(t, test.status, res.Code, test.tag)
		if test.response != "" {
			assert.JSONEq(t, test.response, res.Body.String(), test.tag)
		}
	}
}
