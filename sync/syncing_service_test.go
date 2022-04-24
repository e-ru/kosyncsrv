package sync_test

import (
	"errors"
	"kosyncsrv/sync"
	"kosyncsrv/test/mocks"
	"kosyncsrv/test/utils"
	"kosyncsrv/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Synth_Service_Get_Progress(t *testing.T) {
	// GIVEN
	username := "username"
	documentId := "documentId"

	testcases := []struct {
		name    string
		docPos  *types.DocumentPosition
		err     error
		wantErr bool
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

func Test_Synth_Service_Update_Progress(t *testing.T) {
	// GIVEN
	username := "username"
	documentId := "documentId"

	testcases := []struct {
		name      string
		docPos    *types.DocumentPosition
		timestamp *int64
		err       error
		wantErr   bool
	}{
		{name: "Update Progress Successful", docPos: &types.DocumentPosition{DocumentID: documentId, Percentage: 10.0, Progress: "prog", Device: "dev", DeviceID: "1"}, timestamp: utils.PtrInt64(0), wantErr: false},
		{name: "Update Progress Unsuccessful", err: errors.New("Could not update Doc Position"), wantErr: true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			repo := new(mocks.MockedRepo)
			repo.On("UpdateDocumentPosition", username, testcase.docPos).Return(testcase.timestamp, testcase.err)

			// WHEN
			syncService := sync.NewSyncingService(repo)
			res, err := syncService.UpdateProgress(username, testcase.docPos)

			// THEN
			repo.AssertExpectations(t)

			if (err != nil) != testcase.wantErr {
				t.Errorf("syncService.UpdateProgress error = %v, wantErr %v", err, testcase.wantErr)
			} else if res != nil {
				assert.Equal(t, utils.PtrInt64(0), res)
			}
		})
	}
}
