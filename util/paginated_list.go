package util

import (
	"fmt"
	"strings"
)

// PaginatedList represents a paginated list of data items.
type PaginatedList struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	PageCount  int         `json:"page_count"`
	TotalCount int         `json:"total_count"`
	Items      interface{} `json:"items"`
}

// NewPaginatedList creates a new Paginated instance.
// The page parameter is 1-based and refers to the current page index/number.
// The perPage parameter refers to the number of items on each page.
// And the total parameter specifies the total number of data items.
// If total is less than 0, it means total is unknown.
func NewPaginatedList(page, perPage, total int) *PaginatedList {
	if perPage < 1 {
		perPage = 100
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	return &PaginatedList{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
		PageCount:  pageCount,
	}
}

// Offset returns the OFFSET value that can be used in a SQL statement.
func (p *PaginatedList) Offset() int {
	return (p.Page - 1) * p.PerPage
}

// Limit returns the LIMIT value that can be used in a SQL statement.
func (p *PaginatedList) Limit() int {
	return p.PerPage
}

// BuildLinkHeader returns an HTTP header containing the links about the pagination.
func (p *PaginatedList) BuildLinkHeader(baseUrl string, defaultPerPage int) string {
	links := p.BuildLinks(baseUrl, defaultPerPage)
	header := ""
	if links[0] != "" {
		header += fmt.Sprintf("<%v>; rel=\"first\", ", links[0])
		header += fmt.Sprintf("<%v>; rel=\"prev\"", links[1])
	}
	if links[2] != "" {
		if header != "" {
			header += ", "
		}
		header += fmt.Sprintf("<%v>; rel=\"next\"", links[2])
		if links[3] != "" {
			header += fmt.Sprintf(", <%v>; rel=\"last\"", links[3])
		}
	}
	return header
}

// BuildLinks returns the first, prev, next, and last links corresponding to the pagination.
// A link could be an empty string if it is not needed.
// For example, if the pagination is at the first page, then both first and prev links
// will be empty.
func (p *PaginatedList) BuildLinks(baseUrl string, defaultPerPage int) [4]string {
	var links [4]string
	pageCount := p.PageCount
	page := p.Page
	if pageCount >= 0 && page > pageCount {
		page = pageCount
	}
	if strings.Contains(baseUrl, "?") {
		baseUrl += "&"
	} else {
		baseUrl += "?"
	}
	if page > 1 {
		links[0] = fmt.Sprintf("%vpage=%v", baseUrl, 1)
		links[1] = fmt.Sprintf("%vpage=%v", baseUrl, page-1)
	}
	if pageCount >= 0 && page < pageCount {
		links[2] = fmt.Sprintf("%vpage=%v", baseUrl, page+1)
		links[3] = fmt.Sprintf("%vpage=%v", baseUrl, pageCount)
	} else if pageCount < 0 {
		links[2] = fmt.Sprintf("%vpage=%v", baseUrl, page+1)
	}
	if perPage := p.PerPage; perPage != defaultPerPage {
		for i := 0; i < 4; i++ {
			if links[i] != "" {
				links[i] += fmt.Sprintf("&per_page=%v", perPage)
			}
		}
	}

	return links
}
