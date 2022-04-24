package sync_test

import (
	"errors"
	"kosyncsrv/sync"
	"kosyncsrv/test/mocks"
	"kosyncsrv/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Synth_Service_Get_Progress(t *testing.T) {
	// GIVEN
	username := "username"
	documentId := "documentId"

	testcases := []struct {
		name          string
		docPos    *types.DocumentPosition
		err	error
		wantErr   bool
	}{
		{name: "Get Progress Successful", docPos: &types.DocumentPosition{DocumentID: documentId, Percentage: 10.0, Progress: "prog", Device: "dev", DeviceID: "1"}, wantErr: false},
		{name: "Get Progress Unsuccessful", err: errors.New("Could not get Doc Position"), wantErr: true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			repo := new(mocks.MockedRepo)
			repo.On("GetDocumentPosition", username, documentId).Return(testcase.docPos, testcase.err)

			// WHEN
			syncService := sync.NewSyncingService(repo)
			res, err := syncService.GetProgress(username, documentId)
			// success, msg := authService.RegisterUser(username, password)

			// THEN
			repo.AssertExpectations(t)

			if (err != nil) != testcase.wantErr {
				t.Errorf("syncService.GetProgress error = %v, wantErr %v", err, testcase.wantErr)
			} else if res != nil {
				assert.Equal(t, testcase.docPos, res)
			}
		})
	}
}