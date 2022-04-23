package sync

import "kosyncsrv/types"

type syncingService struct {
	repo types.Repo
}

func NewSyncingService(repo types.Repo) types.SyncingService{
	return &syncingService{repo: repo}
}

func (s *syncingService) GetProgress() {

}

func (s *syncingService) UpdateProgress() {
	
}
