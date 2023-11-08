package api

import "time"

type UrlProcessResult struct {
	IsValid       bool
	ContentLength int64
	ProcessTime   time.Duration
}

type FileProcessResult struct {
	IsValid bool
	Lines   []string
}

type UrlService interface {
	ProcessMany(input FileProcessResult) ([]UrlProcessResult, error)
}

type FileService interface {
	ReadFile() (FileProcessResult, error)
}
