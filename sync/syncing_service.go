package sync

import "kosyncsrv/types"

type syncingService struct {
	repo types.Repo
}

func NewSyncingService(repo types.Repo) types.SyncingService{
	return &syncingService{repo: repo}
}

func (s *syncingService) GetProgress(username, documentId string) (*types.DocumentPosition, error) {
	ret, err := s.repo.GetDocumentPosition(username, documentId)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *syncingService) UpdateProgress(username string, documentPosition *types.DocumentPosition) (*int64, error) {
	timestamp, err := s.repo.UpdateDocumentPosition(username, documentPosition)
	if err != nil {
		return nil, err
	}
	return timestamp, nil
}
