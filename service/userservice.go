package userservice

import (
	"github.com/eabartlett/userservice/domain"
)

// User representing user service
type UserService domain.User

type Result struct {
	User *domain.User
}

type GetByIdArgs struct {
	UserId uint64
}

/* CreateUser for creating user in service
 * @user - new user to create in service
 * @return - error if error creating user, nil otherwise
 */
func (u *UserService) CreateUser(user *domain.User, result *Result) (err error) {
	_, err = domain.ValidUser(user)
	return err
}

/* GetUserByID
 * @id - user id of user to fetch
 * @return - pointer to user object found for given id,
 * error if there was an error retrieving user - is nil otherwise
 */
func (u *UserService) GetUserByID(args *GetByIdArgs, result *Result) (err error) {
	user := &domain.User{ID: args.UserId, Username: "barkledad", Email: "eb21@sbcglobal.net"}
	result.User = user
	return nil
}
