package service

import (
	"github.com/agbankar/navigation-service/internal/model"
	"sync"
)

type Visitor interface {
	Visit(u *model.User) error
	GetUniqueVisits(url string) int
}

type VisitorService struct {
	PageVisits map[string]model.PageDetails
	Lock       *sync.RWMutex
}

var VoidStruct model.EmptyStruct

func (s *VisitorService) Visit(u *model.User) error {
	return s.write(u)
}

func (s *VisitorService) GetUniqueVisits(url string) int {
	visits := s.read(url)
	if visits != nil {
		return visits.Counter
	}
	return 0
}
func (s *VisitorService) write(u *model.User) error {
	pageData := s.read(u.Url)
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if pageData != nil {
		_, ok := pageData.UserIds[u.UserId]
		if !ok {
			ids := pageData.UserIds
			ids[u.UserId] = VoidStruct
			data := model.PageDetails{
				Counter: pageData.Counter + 1,
				UserIds: ids,
			}
			s.PageVisits[u.Url] = data
		}
		return nil
	}

	m := make(map[string]model.EmptyStruct)
	m[u.UserId] = VoidStruct
	s.PageVisits[u.Url] = model.PageDetails{
		UserIds: m,
		Counter: 1,
	}
	return nil

}

func (s *VisitorService) read(uid string) *model.PageDetails {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	visit, ok := s.PageVisits[uid]
	if !ok {
		return nil
	}
	return &visit
}
