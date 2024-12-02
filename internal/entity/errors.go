package entity

import "errors"

var (
	ErrEmptyID  = errors.New("empty ID")
	ErrNotFound = errors.New("not found")
)
