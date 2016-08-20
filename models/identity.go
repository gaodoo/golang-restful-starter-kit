package models

type Identity interface {
	GetID() string
	GetName() string
}

type User struct {
	ID   string
	Name string
}

func (u User) GetID() string {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}
