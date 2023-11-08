package http_native

import (
	"net/http"
	"odo_office/pkg/api"
	"sync"
	"time"
)

type job struct {
	Num int
	Url string
}

type resultChan struct {
	Num          int
	IsValid      bool
	ContentLen   int64
	ResponseTime time.Duration
}

type urlService struct {
	numUrls int
	results chan resultChan
}

func NewUrlService() (api.UrlService, error) {
	return &urlService{}, nil
}

func (us *urlService) ProcessMany(input api.FileProcessResult) ([]api.UrlProcessResult, error) {
	numUrls := len(input.Lines)
	us.results = make(chan resultChan, numUrls)
	wg := &sync.WaitGroup{}
	wg.Add(numUrls)
	resultTable := make([]api.UrlProcessResult, numUrls)
	for i := 0; i < numUrls; i++ {
		go us.process(job{Num: i, Url: input.Lines[i]}, wg)
	}

	go func() {
		for {
			select {
			case result := <-us.results:
				resultTable[result.Num] = api.UrlProcessResult{ContentLength: result.ContentLen, ProcessTime: result.ResponseTime, IsValid: result.IsValid}
			}
		}
	}()
	wg.Wait()
	return resultTable, nil
}

func (us *urlService) process(job job, wg *sync.WaitGroup) error {
	defer wg.Done()
	prev := time.Now()
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		us.results <- resultChan{IsValid: false, Num: job.Num}
		return err
	}
	total := time.Since(prev)
	us.results <- resultChan{Num: job.Num, ContentLen: resp.ContentLength, ResponseTime: total, IsValid: true}
	return nil
}
