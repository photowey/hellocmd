package repository

import (
	"sync"
)

var (
	_repositoryRegistry = make(map[string]any)
	_lock               sync.Mutex

	_engine Engine = (*engine)(nil)
)

func init() {
	initEngine()

	initRepository()
}

func initEngine() {
	_engine = newEngine()
}

func initRepository() {}

// Engine - repository engine
type Engine interface {
}

type engine struct{}

// ---------------------------------------------------------------- private

func newEngine() Engine {
	return &engine{}
}
