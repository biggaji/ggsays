package dto

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	UserName  string    `json:"userName"`
}

type UserJwtPayload struct {
	ID uuid.UUID `json:"id"`
}

type UserAuthentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
