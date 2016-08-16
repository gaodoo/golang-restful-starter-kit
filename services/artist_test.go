package services

import (
	"errors"
	"testing"

	"github.com/qiangxue/golang-restful-starter-kit/app"
	"github.com/qiangxue/golang-restful-starter-kit/models"
	"github.com/stretchr/testify/assert"
)

func TestNewArtistService(t *testing.T) {
	dao := newMockArtistDAO()
	s := NewArtistService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestArtistService_Get(t *testing.T) {
	s := NewArtistService(newMockArtistDAO())
	artist, err := s.Get(nil, 1)
	if assert.Nil(t, err) && assert.NotNil(t, artist) {
		assert.Equal(t, "aaa", artist.Name)
	}

	artist, err = s.Get(nil, 100)
	assert.NotNil(t, err)
}

func TestArtistService_Create(t *testing.T) {
	s := NewArtistService(newMockArtistDAO())
	artist, err := s.Create(nil, &models.Artist{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, artist) {
		assert.Equal(t, 4, artist.Id)
		assert.Equal(t, "ddd", artist.Name)
	}

	// dao error
	_, err = s.Create(nil, &models.Artist{
		Id:   100,
		Name: "ddd",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Create(nil, &models.Artist{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestArtistService_Update(t *testing.T) {
	s := NewArtistService(newMockArtistDAO())
	artist, err := s.Update(nil, 2, &models.Artist{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, artist) {
		assert.Equal(t, 2, artist.Id)
		assert.Equal(t, "ddd", artist.Name)
	}

	// dao error
	_, err = s.Update(nil, 100, &models.Artist{
		Name: "ddd",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Update(nil, 2, &models.Artist{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestArtistService_Delete(t *testing.T) {
	s := NewArtistService(newMockArtistDAO())
	artist, err := s.Delete(nil, 2)
	if assert.Nil(t, err) && assert.NotNil(t, artist) {
		assert.Equal(t, 2, artist.Id)
		assert.Equal(t, "bbb", artist.Name)
	}

	_, err = s.Delete(nil, 2)
	assert.NotNil(t, err)
}

func TestArtistService_Query(t *testing.T) {
	s := NewArtistService(newMockArtistDAO())
	result, err := s.Query(nil, 1, 2)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockArtistDAO() artistDAO {
	return &mockArtistDAO{
		records: []models.Artist{
			{Id: 1, Name: "aaa"},
			{Id: 2, Name: "bbb"},
			{Id: 3, Name: "ccc"},
		},
	}
}

type mockArtistDAO struct {
	records []models.Artist
}

func (m *mockArtistDAO) Get(rs app.RequestScope, id int) (*models.Artist, error) {
	for _, record := range m.records {
		if record.Id == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockArtistDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockArtistDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.records), nil
}

func (m *mockArtistDAO) Create(rs app.RequestScope, artist *models.Artist) error {
	if artist.Id != 0 {
		return errors.New("Id cannot be set")
	}
	artist.Id = len(m.records) + 1
	m.records = append(m.records, *artist)
	return nil
}

func (m *mockArtistDAO) Update(rs app.RequestScope, id int, artist *models.Artist) error {
	artist.Id = id
	for i, record := range m.records {
		if record.Id == id {
			m.records[i] = *artist
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockArtistDAO) Delete(rs app.RequestScope, id int) error {
	for i, record := range m.records {
		if record.Id == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
