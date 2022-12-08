package main

import "github.com/jrvldam/advent-code-2022/shared/stack"

const (
	ErrNotFound = StoreErr("could not found the stack")
)

type StoreErr string

func (e StoreErr) Error() string {
	return string(e)
}

type Store map[int]*stack.Stack

func (s Store) Search(idx int) (*stack.Stack, error) {
	content, ok := s[idx]
	if !ok {
		return nil, ErrNotFound
	}

	return content, nil
}

func (s Store) Add(idx int, newStack stack.Stack) {
	_, err := s.Search(idx)

	switch err {
	case ErrNotFound:
		s[idx] = &newStack
	default:
		return
	}
}

func NewStore() Store {
	return Store{}
}
