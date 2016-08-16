package daos

import (
	"github.com/qiangxue/golang-restful-starter-kit/app"
	"github.com/qiangxue/golang-restful-starter-kit/models"
)

// ArtistDAO persists artist data in database
type ArtistDAO struct{}

// NewArtistDAO creates a new ArtistDAO
func NewArtistDAO() *ArtistDAO {
	return &ArtistDAO{}
}

// Get reads the artist with the specified ID from the database.
func (dao *ArtistDAO) Get(rs app.RequestScope, id int) (*models.Artist, error) {
	var artist models.Artist
	err := rs.Tx().Select().Model(id, &artist)
	return &artist, err
}

// Create saves a new artist record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *ArtistDAO) Create(rs app.RequestScope, artist *models.Artist) error {
	artist.Id = 0
	return rs.Tx().Model(artist).Insert()
}

// Update saves the changes to an artist in the database.
func (dao *ArtistDAO) Update(rs app.RequestScope, id int, artist *models.Artist) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	artist.Id = id
	return rs.Tx().Model(artist).Exclude("Id").Update()
}

// Delete deletes an artist with the specified ID from the database.
func (dao *ArtistDAO) Delete(rs app.RequestScope, id int) error {
	artist, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(artist).Delete()
}

// Count returns the number of the artist records in the database.
func (dao *ArtistDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
	return count, err
}

// Query retrieves the artist records with the specified offset and limit from the database.
func (dao *ArtistDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
	artists := []models.Artist{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&artists)
	return artists, err
}
