package app

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/go-ozzo/ozzo-dbx"
)

// RequestScope contains the application-specific information that are carried around in a request.
type RequestScope interface {
	Logger
	// UserID returns the ID of the user for the current request
	UserID() string
	// Tx returns the currently active database transaction that can be used for DB query purpose
	Tx() *dbx.Tx
	// SetTx sets the database transaction
	SetTx(tx *dbx.Tx)
	// Rollback returns a value indicating whether the current database transaction should be rolled back
	Rollback() bool
	// SetRollback sets a value indicating whether the current database transaction should be rolled back
	SetRollback(bool)
	// Now returns the timestamp representing the time when the request is being processed
	Now() time.Time
}

type requestScope struct {
	Logger             // the logger tagged with the current request information
	now      time.Time // the time when the request is being processed
	userID   string    // an ID identifying the current user
	rollback bool      // whether to roll back the current transaction
	tx       *dbx.Tx   // the currently active transaction
}

func (rs *requestScope) UserID() string {
	return rs.userID
}

func (rs *requestScope) Tx() *dbx.Tx {
	return rs.tx
}

func (rs *requestScope) SetTx(tx *dbx.Tx) {
	rs.tx = tx
}

func (rs *requestScope) Rollback() bool {
	return rs.rollback
}

func (rs *requestScope) SetRollback(v bool) {
	rs.rollback = v
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

// newRequestScope creates a new RequestScope with the current request information.
func newRequestScope(now time.Time, logger *logrus.Logger, request *http.Request) RequestScope {
	userID := authenticate(request)
	return &requestScope{
		Logger: NewLogger(logger, logrus.Fields{
			"UserID": userID,
		}),
		now:    now,
		userID: userID,
	}
}

// authenticate authenticates the current user.
// The default implementation simply returns the value of the X-User-Id HTTP header, which assumes
// the authentication has been done by an API gateway.
// If this is not the case, you may customize this method with the actual authentication logic.
func authenticate(request *http.Request) string {
	return request.Header.Get("X-User-Id")
}
