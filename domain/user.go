package domain

import (
	"errors"
)

/**
 * User:
 * Base user object for the service
 */
type User struct {
	ID       uint64
	Username string
	Email    string
}

/**
 * ValidUser
 * @param user - user struct to validate
 * @return - whether the user is valid or not
 */
func ValidUser(user *User) (isValid bool, err error) {
	valid, err := validateID(user.ID)
	if !valid {
		return false, err
	}
	valid, err = validateUsername(user.Username)
	if !valid {
		return false, err
	}
	valid, err = validateEmail(user.Email)
	if !valid {
		return false, err
	}
	return true, nil
}

func validateID(id uint64) (bool, error) {
	return true, nil // probably do a db lookup for uniqueness or something
}

func validateUsername(name string) (bool, error) {
	if len(name) <= 0 {
		return false, errors.New("username must not be empty")
	}
	return true, nil
}

func validateEmail(email string) (bool, error) {
	if len(email) <= 0 {
		return false, errors.New("username must not be empty")
	}
	return true, nil
}
