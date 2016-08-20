package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/qiangxue/golang-restful-starter-kit/app"
	"github.com/qiangxue/golang-restful-starter-kit/models"
)

type (
	// artistService specifies the interface for the artist service needed by artistResource.
	artistService interface {
		Get(rs app.RequestScope, id int) (*models.Artist, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Artist) (*models.Artist, error)
		Update(rs app.RequestScope, id int, model *models.Artist) (*models.Artist, error)
		Delete(rs app.RequestScope, id int) (*models.Artist, error)
	}

	// artistResource defines the handlers for the CRUD APIs.
	artistResource struct {
		service artistService
	}
)

// ServeArtist sets up the routing of artist endpoints and the corresponding handlers.
func ServeArtistResource(rg *routing.RouteGroup, service artistService) {
	r := &artistResource{service}
	rg.Get("/artists/<id>", r.get)
	rg.Get("/artists", r.query)
	rg.Post("/artists", r.create)
	rg.Put("/artists/<id>", r.update)
	rg.Delete("/artists/<id>", r.delete)
}

func (r *artistResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *artistResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *artistResource) create(c *routing.Context) error {
	var model models.Artist
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *artistResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *artistResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
