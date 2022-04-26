package auth

import (
	"errors"
	"kosyncsrv/types"
	"log"
)

type authService struct {
	repo types.Repo
}

func NewAuthService(repo types.Repo) types.AuthorizationService {
	return &authService{repo: repo}
}

func (a *authService) RegisterUser(username, password string) (error, *string) {
	log.Printf("Register user. Username: %+v", username)

	if err := a.repo.AddUser(username, password); err != nil {
		log.Printf("Could not register user. Error: %+v", err)
		return errors.New("Could not register user. PLease check the logs for more details"), nil
	} else {
		return nil, &username
	}
}

func (a *authService) AuthorizeUser(username, password string) types.AuthReturnCode {
	log.Printf("Authorize User: %+v", username)

	user, err := a.repo.GetUser(username)
	if err != nil {
		log.Printf("Could not get User: %+v. Error: %+v", username, err)
		return types.Forbidden
	}
	if password != user.Password {
		log.Printf("Invalid Password for User: %+v.", username)
		return types.Unauthorized
	}
	return types.Allowed
}
