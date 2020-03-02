package users

import "fmt"

// UserDTO is the data transfer object structure used for json marshalling
type UserDTO struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Timezone  string `json:"timezone"`
}

// User This interface represents a Harvest user
type User interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetTimezone() string
}

type user struct {
	id       string
	name     string
	email    string
	timeZone string
}

// GetID returns the user ID from the database
func (u *user) GetID() string {
	return u.id
}

// GetName returns the user's name
func (u *user) GetName() string {
	return u.name
}

// GetEmail returns the user's email
func (u *user) GetEmail() string {
	return u.email
}

// GetPhoneNumber returns the user's phone number
func (u *user) GetTimezone() string {
	return u.timeZone
}

// NewFromDTO serializes a DTO into a User model
func NewFromDTO(d *UserDTO) (User, error) {
	return NewUser(
		d.ID,
		fmt.Sprintf("%s %s", d.FirstName, d.LastName),
		d.Email,
		d.Timezone,
	), nil
}

// NewUser is the constructor for the user type
func NewUser(id string, name string, email string, tz string) User {
	return &user{id, name, email, tz}
}
