package model

// User : User domain model
type User struct {
	id   string
	name string
}

// ID : Getter
func (u User) ID() string {
	return u.id
}

// Name : Getter
func (u User) Name() string {
	return u.name
}

// IsEmpty : Determine whether the user is empty
func (u User) IsEmpty() bool {
	return u.id == ""
}

// NewUser : Constructor
func NewUser(id, name string) User {
	return User{
		id:   id,
		name: name,
	}
}
