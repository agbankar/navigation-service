package service

import (
	"fmt"
	"github.com/agbankar/navigation-service/internal/model"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestSequentialUsers(t *testing.T) {
	visitorService := &VisitorService{
		PageVisits: make(map[string]model.PageDetails),
		Lock:       &sync.RWMutex{},
	}
	for i := 0; i < 10; i++ {
		user := model.User{
			UserId: strconv.Itoa(i),
			Url:    "page-1",
		}
		visitorService.Visit(&user)
	}
	assert.Equal(t, 10, visitorService.GetUniqueVisits("page-1"))

}

func TestSameUser(t *testing.T) {
	visitorService := &VisitorService{
		PageVisits: make(map[string]model.PageDetails),
		Lock:       &sync.RWMutex{},
	}
	for i := 0; i < 10; i++ {
		user := model.User{
			UserId: strconv.Itoa(1),
			Url:    "page-1",
		}
		visitorService.Visit(&user)
	}
	assert.Equal(t, 1, visitorService.GetUniqueVisits("page-1"))

}

func TestConcurrentRead(t *testing.T) {
	visitorService := &VisitorService{
		PageVisits: make(map[string]model.PageDetails),
		Lock:       &sync.RWMutex{},
	}
	for i := 0; i < 10; i++ {
		user := model.User{
			UserId: strconv.Itoa(1),
			Url:    "page-1",
		}
		visitorService.Visit(&user)
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//time.Sleep(time.Second)
		go func() {
			defer wg.Done()
			url := "page-" + strconv.Itoa(i)
			fmt.Println(url)
			visits := visitorService.read(url)
			if visits != nil {

			}
		}()

	}
	wg.Wait()

}

func TestConcurrentWrite(t *testing.T) {
	visitorService := &VisitorService{
		PageVisits: make(map[string]model.PageDetails),
		Lock:       &sync.RWMutex{},
	}
	wg := sync.WaitGroup{}
	n := 100
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		wg.Add(1)
		go func() {
			defer wg.Done()

			url := "page-" + strconv.Itoa(i)
			user := &model.User{
				UserId: strconv.Itoa(i),
				Url:    url,
			}
			visitorService.Visit(user)
		}()
	}
	wg.Wait()
	assert.Equal(t, n, len(visitorService.PageVisits))
}
