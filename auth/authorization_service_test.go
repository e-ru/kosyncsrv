package auth_test

import (
	"errors"
	"kosyncsrv/auth"
	"kosyncsrv/test/mocks"
	"kosyncsrv/test/utils"
	"kosyncsrv/types"

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Auth_Service_Register_User(t *testing.T) {
	// GIVEN
	username := "username"
	password := "password"

	testcases := []struct {
		name string
		retSuccessMsg *string
		err     error
		wantErr bool
	}{
		{name: "New User Successful", retSuccessMsg: utils.PtrString(username), wantErr: false},
		{name: "New User Unsuccessful", err: errors.New("Db error"), wantErr: true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			repo := new(mocks.MockedRepo)
			repo.On("AddUser", username, password).Return(testcase.err)

			// WHEN
			authService := auth.NewAuthService(repo)
			err, msg := authService.RegisterUser(username, password)

			// THEN
			repo.AssertExpectations(t)

			if (err != nil) != testcase.wantErr {
				t.Errorf("authService.RegisterUser error = %v, wantErr %v", err, testcase.wantErr)
			} else {
				assert.Equal(t, testcase.retSuccessMsg, msg)
			}
			// if testcase.wantSuccess {
			// 	assert.NoError(t, err)
			// } else {
			// 	assert.Error(t, err, "Could not create user. Error: Db error")
			// }
		})
	}
}

func Test_Auth_Service_Authorize_User(t *testing.T) {
	// GIVEN
	username := "username"
	password := "password"

	testcases := []struct {
		name       string
		userExists bool
		user       *types.User
		retCode    types.AuthReturnCode
		retMsg     *string
	}{
		{name: "Authorize User Successful", userExists: true, user: &types.User{Username: username, Password: password}, retCode: types.Allowed, retMsg: utils.PtrString(username)},
		{name: "Authorize User Forbidden", userExists: false, retCode: types.Forbidden, retMsg: utils.PtrString(username)},
		{name: "Authorize User Unauthorized", userExists: true, user: &types.User{Username: username, Password: "foo"}, retCode: types.Unauthorized, retMsg: utils.PtrString(username)},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			repo := new(mocks.MockedRepo)
			repo.On("GetUser", username).Return(testcase.user, testcase.userExists)

			// WHEN
			authService := auth.NewAuthService(repo)
			returnCode, msg := authService.AuthorizeUser(username, password)

			// THEN
			repo.AssertExpectations(t)

			assert.Equal(t, testcase.retCode, returnCode)
			assert.Equal(t, *testcase.retMsg, msg)
		})
	}
}
