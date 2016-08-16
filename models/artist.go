package models

import "github.com/go-ozzo/ozzo-validation"

// Artist represents an artist record.
type Artist struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Validate validates the Artist fields.
func (m Artist) Validate(attrs ...string) error {
	return validation.StructRules{}.
		Add("Name", validation.Required, validation.Length(0, 120)).
		Validate(m, attrs...)
}
