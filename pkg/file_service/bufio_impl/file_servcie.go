package bufio_impl

import (
	"bufio"
	"odo_office/pkg/api"
	"os"
)

type fileService struct {
	fp       *os.File
	filename string
	lines    []string
}

func NewFileService(filename string) (api.FileService, error) {
	return &fileService{filename: filename}, nil
}

func (fs *fileService) open() error {
	file, err := os.Open(fs.filename)
	if err != nil {
		return err
	}
	fs.fp = file
	return nil
}

func (fs *fileService) close() error {
	fs.fp.Close()
	return nil
}

func (fs *fileService) ReadFile() (api.FileProcessResult, error) {
	err := fs.open()
	if err != nil {
		return api.FileProcessResult{IsValid: false}, err
	}
	defer fs.close()
	scanner := bufio.NewScanner(fs.fp)
	for scanner.Scan() {
		fs.lines = append(fs.lines, scanner.Text())
	}
	return api.FileProcessResult{Lines: fs.lines, IsValid: true}, scanner.Err()
}
