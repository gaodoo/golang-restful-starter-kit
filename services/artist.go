package services

import (
	"github.com/qiangxue/golang-restful-starter-kit/app"
	"github.com/qiangxue/golang-restful-starter-kit/models"
)

// artistDAO specifies the interface of the artist DAO needed by ArtistService.
type artistDAO interface {
	// Get returns the artist with the specified the artist ID.
	Get(rs app.RequestScope, id int) (*models.Artist, error)
	// Count returns the number of artists.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of the artists with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error)
	// Create saves a new artist in the storage.
	Create(rs app.RequestScope, artist *models.Artist) error
	// Update updates the artist with the given ID in the storage.
	Update(rs app.RequestScope, id int, artist *models.Artist) error
	// Delete removes the artist with the given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// ArtistService provides services related with artists.
type ArtistService struct {
	dao artistDAO
}

// NewArtistService creates a new ArtistService with the given artist DAO.
func NewArtistService(dao artistDAO) *ArtistService {
	return &ArtistService{dao}
}

// Get returns the artist with the specified the artist ID.
func (s *ArtistService) Get(rs app.RequestScope, id int) (*models.Artist, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new artist.
func (s *ArtistService) Create(rs app.RequestScope, model *models.Artist) (*models.Artist, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the artist with the specified ID.
func (s *ArtistService) Update(rs app.RequestScope, id int, model *models.Artist) (*models.Artist, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the artist with the specified ID.
func (s *ArtistService) Delete(rs app.RequestScope, id int) (*models.Artist, error) {
	artist, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return artist, err
}

// Count returns the number of artists.
func (s *ArtistService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the artists with the specified offset and limit.
func (s *ArtistService) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
	return s.dao.Query(rs, offset, limit)
}
