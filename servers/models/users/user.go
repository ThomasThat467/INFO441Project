package users

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

//gravatarBasePhotoURL is the base URL for Gravatar image requests.
//See https://id.gravatar.com/site/implement/images/ for details
const gravatarBasePhotoURL = "https://www.gravatar.com/avatar/"

//bcryptCost is the default bcrypt cost to use when hashing passwords
var bcryptCost = 13

//User represents a user account in the database
type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"-"` //never JSON encoded/decoded
	PassHash  []byte `json:"-"` //never JSON encoded/decoded
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PhotoURL  string `json:"photoURL"`
}

//Credentials represents user sign-in credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

//Updates represents allowed updates to a user profile
type Updates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//Validate validates the new user and returns an error if
//any of the validation rules fail, or nil if its valid
func (nu *NewUser) Validate() error {
	if _, emailErr := mail.ParseAddress(nu.Email); emailErr != nil {
		return emailErr
	}

	if len(nu.Password) < 6 {
		return fmt.Errorf("Password must be at least 6 characters")
	}

	if nu.Password != nu.PasswordConf {
		return fmt.Errorf("Password and confirmation must match")
	}

	if len(nu.UserName) == 0 || strings.Contains(nu.UserName, " ") {
		return fmt.Errorf("UserName must be non-zero length and may not contain spaces")
	}

	return nil
}

//ToUser converts the NewUser to a User, setting the
//PhotoURL and PassHash fields appropriately
func (nu *NewUser) ToUser() (*User, error) {
	validationErr := nu.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	newUser := &User{
		Email:     nu.Email,
		UserName:  nu.UserName,
		FirstName: nu.FirstName,
		LastName:  nu.LastName,
	}

	passwordHashErr := newUser.SetPassword(nu.Password)
	if passwordHashErr != nil {
		return nil, passwordHashErr
	}

	GetGravitar(newUser, nu.Email)
	return newUser, nil
}

// GetGravitar calculates the gravitar hash based on the string given and
// stores it for the user
func GetGravitar(user *User, str string) {
	photoURLHash := md5.Sum([]byte(strings.ToLower(strings.TrimSpace(str))))
	photoURLHashString := hex.EncodeToString(photoURLHash[:])
	user.PhotoURL = gravatarBasePhotoURL + photoURLHashString
}

//FullName returns the user's full name, in the form:
// "<FirstName> <LastName>"
//If either first or last name is an empty string, no
//space is put between the names. If both are missing,
//this returns an empty string
func (u *User) FullName() string {
	totalStringArr := []string{}
	if u.FirstName != "" {
		totalStringArr = append(totalStringArr, u.FirstName)
	}

	if u.LastName != "" {
		totalStringArr = append(totalStringArr, u.LastName)
	}

	return strings.Join(totalStringArr, " ")
}

//SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return err
	}
	u.PassHash = passHash
	return nil
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	err := bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

//ApplyUpdates applies the updates to the user. An error
//is returned if the updates are invalid
func (u *User) ApplyUpdates(updates *Updates) error {
	// Sure hope there isn't a catch to this function. I don't think
	// it said to modify Updates in any way, and if it doesn't change then
	// this is valid.
	if updates.FirstName != "" {
		u.FirstName = updates.FirstName
	}

	if updates.LastName != "" {
		u.LastName = updates.LastName
	}

	return nil
}
