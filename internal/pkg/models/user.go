package models

import (
	"errors"
	"fmt"
)

/*
User object
*/
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

//GetUsers ...
func GetUsers() []*User {
	return users
}

//AddUser ...
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("ID must not be preset")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

//GetUserByID ...
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

//UpdateUser ...
func UpdateUser(tu User) (User, error) {
	for i, u := range users {
		if u.ID == tu.ID {
			users[i] = &tu
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", tu.ID)
}

//RemoveUserByID ...
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
