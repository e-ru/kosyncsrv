package auth_test

import (
	"kosyncsrv/auth"
	"kosyncsrv/test/mocks"
	"kosyncsrv/types"

	"testing"

	"github.com/stretchr/testify/assert"
)

func ptrString(str string) *string {
	return &str
}

func Test_Auth_Service_Register_User(t *testing.T) {
	// GIVEN
	username := "username"
	password := "password"

	testcases := []struct {
		name          string
		retSuccess    bool
		retSuccessMsg *string
		retFailureMsg *string
		wantSuccess   bool
	}{
		{name: "New User Successful", retSuccess: true, retSuccessMsg: ptrString(username), wantSuccess: true},
		{name: "New User Unsuccessful", retSuccess: false, retFailureMsg: ptrString("User already exists"), wantSuccess: false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			repo := new(mocks.MockedRepo)
			repo.On("AddUser", username, password).Return(testcase.retSuccess)

			// WHEN
			authService := auth.NewAuthService(repo)
			success, msg := authService.RegisterUser(username, password)

			// THEN
			repo.AssertExpectations(t)

			if testcase.wantSuccess {
				assert.True(t, success)
				assert.Equal(t, *testcase.retSuccessMsg, msg)
			} else {
				assert.False(t, success)
				assert.Equal(t, *testcase.retFailureMsg, msg)
			}
		})
	}
}

func Test_Auth_Service_Authorize_User(t *testing.T) {
	// GIVEN
	username := "username"
	password := "password"

	testcases := []struct {
		name          string
		userExists    bool
		user          *types.User
		retCode       types.AuthReturnCode
		retMsg *string
	}{
		{name: "Authorize User Successful", userExists: true, user: &types.User{Username: username, Password: password}, retCode: types.Allowed, retMsg: ptrString(username)},
		{name: "Authorize User Forbidden", userExists: false, retCode: types.Forbidden, retMsg: ptrString(username)},
		{name: "Authorize User Unauthorized", userExists: true, user: &types.User{Username: username, Password: "foo"}, retCode: types.Unauthorized, retMsg: ptrString(username)},
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
