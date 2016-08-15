package daos

import (
	"github.com/qiangxue/golang-restful-starter-kit/app"
	"github.com/qiangxue/golang-restful-starter-kit/models"
)

type ArtistDAO struct{}

func NewArtistDAO() *ArtistDAO {
	return &ArtistDAO{}
}

func (dao *ArtistDAO) Get(rs app.RequestScope, id int) (*models.Artist, error) {
	var artist models.Artist
	err := rs.Tx().Select().Model(id, &artist)
	return &artist, err
}

func (dao *ArtistDAO) Create(rs app.RequestScope, artist *models.Artist) error {
	return rs.Tx().Model(artist).Exclude("Id").Insert()
}

func (dao *ArtistDAO) Update(rs app.RequestScope, id int, artist *models.Artist) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	artist.Id = id
	return rs.Tx().Model(artist).Exclude("Id").Update()
}

func (dao *ArtistDAO) Delete(rs app.RequestScope, id int) error {
	artist, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(&artist).Delete()
}

func (dao *ArtistDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
	return count, err
}

func (dao *ArtistDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
	artists := []models.Artist{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&artists)
	return artists, err
}
