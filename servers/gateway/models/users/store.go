package users

import (
	"errors"
)

//ErrUserNotFound is returned when the user can't be found
var ErrUserNotFound = errors.New("user not found")

//ErrUserLogNotFound is returned when the user can't be found
var ErrUserLogNotFound = errors.New("user log not found")

//ErrInsertingUser is the error given when the user did not get correctly inserted into the DB
var ErrInsertingUser = errors.New("Error inserting user into DB")

//ErrInsertingUserLog is the error given when the user did not get correctly inserted into the DB
var ErrInsertingUserLog = errors.New("Error inserting user log into DB")

//ErrUpdateUser is the error given when the user did not get correctly inserted into the DB
var ErrUpdateUser = errors.New("Error updating user")

//ErrDeleteUser is the error given when the user did not get correctly deleted
var ErrDeleteUser = errors.New("Error deleting user")

//ErrDeleteUserLog is the error given when the user did not get correctly deleted
var ErrDeleteUserLog = errors.New("Error deleting user log")

//Store represents a store for Users
type Store interface {
	//GetByID returns the User with the given ID
	GetByID(id int64) (*User, error)

	//GetByEmail returns the User with the given email
	GetByEmail(email string) (*User, error)

	//GetByUserName returns the User with the given Username
	GetByUserName(username string) (*User, error)

	//Insert inserts the user into the database, and returns
	//the newly-inserted User, complete with the DBMS-assigned ID
	Insert(user *User) (*User, error)

	//InsertUserLog inserts a log of a successful user login attempt
	InsertUserLog(userLog *UserLog) (*UserLog, error)

	//Update applies UserUpdates to the given user ID
	//and returns the newly-updated user
	Update(id int64, updates *Updates) (*User, error)

	//Delete deletes the user with the given ID
	Delete(id int64) error
}