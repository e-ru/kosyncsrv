package auth

import "kosyncsrv/types"

type authService struct {
	repo types.Repo
}

func NewAuthService(repo types.Repo) types.AuthorizationService {
	return &authService{repo: repo}
}

func (a *authService) RegisterUser(username, password string) (bool, string) {
	if success := a.repo.AddUser(username, password); success {
		return success, username
	} else {
		return success, "User already exists"
	}
}

func (a *authService) AuthorizeUser() {

}
