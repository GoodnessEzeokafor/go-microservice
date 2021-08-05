package data

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Email     string
	Address   string
	CreatedOn string
	UpdateOn  string
	DeletedOn string
}

var userList = []*User{
	&User{
		ID:        1,
		FirstName: "Goodness",
		LastName:  "Ezeokafor",
		Username:  "goody",
		Email:     "test1@test.com",
		Address:   "Test Address",
		CreatedOn: time.Now().UTC().String(),
		UpdateOn:  time.Now().UTC().String(),
	},
	&User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Username:  "john",
		Email:     "test2@test.com",
		Address:   "Test Address 2",
		CreatedOn: time.Now().UTC().String(),
		UpdateOn:  time.Now().UTC().String(),
	},
	&User{
		ID:        3,
		FirstName: "Jane",
		LastName:  "Doe",
		Username:  "jane",
		Email:     "test3@test.com",
		Address:   "Test Address 2",
		CreatedOn: time.Now().UTC().String(),
		UpdateOn:  time.Now().UTC().String(),
	},
}
