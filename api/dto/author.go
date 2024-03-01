package dto

import "time"

type CreateAuthorRequest struct {
	FirstName   string    `json:"firstName" binding:"required,alpha,min=2,max=15"`
	LastName    string    `json:"lastName" binding:"required,alpha,min=2,max=20"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type UpdateAuthorRequest struct {
	FirstName   string    `json:"firstName,omitempty"`
	LastName    string    `json:"lastName,omitempty"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type AuthorResponse struct {
	Id          int          `json:"id"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	DateOfBirth time.Time    `json:"dateOfBirth"` //it is better not to use omitempty with time.Time because it would be interpreted as the zero value of time.Time, which is January 1, year 1, UTC
	Books       BookResponse `json:"books,omitempty"`
}
