package daos

import (
	"testing"

	"github.com/qiangxue/golang-restful-starter-kit/app"
	"github.com/qiangxue/golang-restful-starter-kit/models"
	"github.com/qiangxue/golang-restful-starter-kit/testdata"
	"github.com/stretchr/testify/assert"
)

func TestArtistDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewArtistDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			artist, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, artist) {
				assert.Equal(t, 2, artist.Id)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			artist := &models.Artist{
				Id:   1000,
				Name: "tester",
			}
			err := dao.Create(rs, artist)
			assert.Nil(t, err)
			assert.NotEqual(t, 1000, artist.Id)
			assert.NotZero(t, artist.Id)
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			artist := &models.Artist{
				Id:   2,
				Name: "tester",
			}
			err := dao.Update(rs, artist.Id, artist)
			assert.Nil(t, rs.Tx().Commit())
			assert.Nil(t, err)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			artist := &models.Artist{
				Id:   2,
				Name: "tester",
			}
			err := dao.Update(rs, 99999, artist)
			assert.Nil(t, rs.Tx().Commit())
			assert.NotNil(t, err)
		})
	}

	{
		// Delete
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 2)
			assert.Nil(t, rs.Tx().Commit())
			assert.Nil(t, err)
		})
	}

	{
		// Delete with error
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 99999)
			assert.Nil(t, rs.Tx().Commit())
			assert.NotNil(t, err)
		})
	}

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			artists, err := dao.Query(rs, 1, 3)
			assert.Nil(t, rs.Tx().Commit())
			assert.Nil(t, err)
			assert.Equal(t, 3, len(artists))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs)
			assert.Nil(t, rs.Tx().Commit())
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
