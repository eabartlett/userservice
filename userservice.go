package userservice

import (
	"github.com/eabartlett/userservice/domain"
)

/* CreateUser for creating user in service
 * @user - new user to create in service
 * @return - error if error creating user, nil otherwise
 */
func CreateUser(user *domain.User) (err error) {
	_, err = domain.ValidUser(user)
	return err
}

/* GetUserByID
 * @id - user id of user to fetch
 * @return - pointer to user object found for given id,
 * error if there was an error retrieving user - is nil otherwise
 */
func GetUserByID(id uint64) (user *domain.User, err error) {
	user = &domain.User{ID: 1, Username: "barkledad", Email: "eb21@sbcglobal.net"}
	return user, nil
}
