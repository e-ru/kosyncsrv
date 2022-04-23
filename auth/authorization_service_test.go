package auth_test

import (
	"kosyncsrv/auth"
	"kosyncsrv/test/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
)

func ptrString(str string) *string {
	return &str
}

func Test_Auth_Service_Register_User_New_User(t *testing.T) {
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
