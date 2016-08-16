package daos

import (
	"time"

	"github.com/go-ozzo/ozzo-dbx"
	"github.com/qiangxue/golang-restful-starter-kit/app"
)

func testDBCall(db *dbx.DB, f func(rs app.RequestScope)) {
	rs := mockRequestScope(db)

	defer func() {
		rs.Tx().Rollback()
	}()

	f(rs)
}

type requestScope struct {
	app.Logger
	tx *dbx.Tx
}

func mockRequestScope(db *dbx.DB) app.RequestScope {
	tx, _ := db.Begin()
	return &requestScope{
		tx: tx,
	}
}

func (rs *requestScope) UserID() string {
	return "tester"
}

func (rs *requestScope) Tx() *dbx.Tx {
	return rs.tx
}

func (rs *requestScope) SetTx(tx *dbx.Tx) {
	rs.tx = tx
}

func (rs *requestScope) Rollback() bool {
	return false
}

func (rs *requestScope) SetRollback(v bool) {
}

func (rs *requestScope) Now() time.Time {
	return time.Now()
}

