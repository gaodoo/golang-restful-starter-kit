package apis

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/stretchr/testify/assert"
)

func Test_getPaginatedListFromRequest(t *testing.T) {
	tests := []struct {
		Tag                 string
		Page, PerPage       int
		ExpPage, ExpPerPage int
	}{
		{"t1", 1, 10, 1, 10},
		{"t2", -1, -1, 1, DEFAULT_PAGE_SIZE},
		{"t2", 0, 0, 1, DEFAULT_PAGE_SIZE},
		{"t3", 2, MAX_PAGE_SIZE + 1, 2, MAX_PAGE_SIZE},
	}
	for _, test := range tests {
		url := "http://www.example.com/search?foo=1"
		if test.Page >= 0 {
			url = fmt.Sprintf("%s&page=%v", url, test.Page)
		}
		if test.PerPage >= 0 {
			url = fmt.Sprintf("%s&per_page=%v", url, test.PerPage)
		}
		req, _ := http.NewRequest("GET", url, nil)
		c := routing.NewContext(nil, req)
		pl := getPaginatedListFromRequest(c, 100000)
		assert.Equal(t, test.ExpPage, pl.Page)
		assert.Equal(t, test.ExpPerPage, pl.PerPage)
	}
}

func Test_parseInt(t *testing.T) {
	assert.Equal(t, 123, parseInt("123", 1))
	assert.Equal(t, 1, parseInt("a123", 1))
	assert.Equal(t, 1, parseInt("", 1))
}
