package entity

type OnSaveCallback func(fileName string) error
type OnDeleteCallback func(fileName string) error
