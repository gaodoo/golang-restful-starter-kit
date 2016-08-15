package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPaginatedList(t *testing.T) {
	tests := []struct {
		tag                                                                    string
		page, perPage, total                                                   int
		expectedPage, expectedPerPage, expectedTotal, pageCount, offset, limit int
	}{
		// varying page
		{"t1", 1, 20, 50, 1, 20, 50, 3, 0, 20},
		{"t2", 2, 20, 50, 2, 20, 50, 3, 20, 20},
		{"t3", 3, 20, 50, 3, 20, 50, 3, 40, 20},
		{"t4", 4, 20, 50, 3, 20, 50, 3, 40, 20},
		{"t5", 0, 20, 50, 1, 20, 50, 3, 0, 20},

		// varying perPage
		{"t6", 1, 0, 50, 1, 100, 50, 1, 0, 100},
		{"t7", 1, -1, 50, 1, 100, 50, 1, 0, 100},
		{"t8", 1, 100, 50, 1, 100, 50, 1, 0, 100},

		// varying total
		{"t9", 1, 20, 0, 1, 20, 0, 0, 0, 20},
		{"t10", 1, 20, -1, 1, 20, -1, -1, 0, 20},
	}

	for _, test := range tests {
		p := NewPaginatedList(test.page, test.perPage, test.total)
		assert.Equal(t, test.expectedPage, p.Page, test.tag)
		assert.Equal(t, test.expectedPerPage, p.PerPage, test.tag)
		assert.Equal(t, test.expectedTotal, p.TotalCount, test.tag)
		assert.Equal(t, test.pageCount, p.PageCount, test.tag)
		assert.Equal(t, test.offset, p.Offset(), test.tag)
		assert.Equal(t, test.limit, p.Limit(), test.tag)
	}
}

func TestPaginatedList_LinkHeader(t *testing.T) {
	baseUrl := "/tokens"
	defaultPerPage := 10
	tests := []struct {
		tag                  string
		page, perPage, total int
		header               string
	}{
		{"t1", 1, 20, 50, "</tokens?page=2&per_page=20>; rel=\"next\", </tokens?page=3&per_page=20>; rel=\"last\""},
		{"t2", 2, 20, 50, "</tokens?page=1&per_page=20>; rel=\"first\", </tokens?page=1&per_page=20>; rel=\"prev\", </tokens?page=3&per_page=20>; rel=\"next\", </tokens?page=3&per_page=20>; rel=\"last\""},
		{"t3", 3, 20, 50, "</tokens?page=1&per_page=20>; rel=\"first\", </tokens?page=2&per_page=20>; rel=\"prev\""},
		{"t4", 0, 20, 50, "</tokens?page=2&per_page=20>; rel=\"next\", </tokens?page=3&per_page=20>; rel=\"last\""},
		{"t5", 4, 20, 50, "</tokens?page=1&per_page=20>; rel=\"first\", </tokens?page=2&per_page=20>; rel=\"prev\""},
		{"t6", 1, 20, 0, ""},
	}
	for _, test := range tests {
		p := NewPaginatedList(test.page, test.perPage, test.total)
		assert.Equal(t, test.header, p.BuildLinkHeader(baseUrl, defaultPerPage), test.tag)
	}

	baseUrl = "/tokens?from=10"
	p := NewPaginatedList(1, 20, 50)
	assert.Equal(t, "</tokens?from=10&page=2&per_page=20>; rel=\"next\", </tokens?from=10&page=3&per_page=20>; rel=\"last\"", p.BuildLinkHeader(baseUrl, defaultPerPage))
}
