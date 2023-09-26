package androidinfo

import (
	"sync"
)

var lock = &sync.Mutex{}

type token struct {
	token_info string
}

var tokenInstance *token

func GetInstance() *token {
	if tokenInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if tokenInstance == nil {
			tokenInstance = &token{}
		}
	}
	return tokenInstance
}

func (t *token) GetToken() string {
	return t.token_info
}

func (t *token) SetToken(token_info string) string {
	t.token_info = token_info
	return "success"
}
