package models

import "github.com/go-ozzo/ozzo-validation"

type Artist struct {
	Id      int      `json:"id" db:"id"`
	Name    string   `json:"name" db:"name"`
}

func (m Artist) Validate(attrs ...string) error {
	return validation.StructRules{}.
		Add("Name", validation.Required, validation.Length(0, 120)).
		Validate(m, attrs...)
}
