package controller

import (
	"sync"
)

var lock = &sync.Mutex{}

type block struct {
	isBlock bool `default:"false"`
}

var tokenInstance *block

func GetBlock() *block {
	if tokenInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if tokenInstance == nil {
			tokenInstance = &block{}
		}
	}
	return tokenInstance
}

func (t *block) SetBlock(is_block bool) {
	t.isBlock = is_block
}

func (t *block) GetBlock() bool {
	return t.isBlock
}
