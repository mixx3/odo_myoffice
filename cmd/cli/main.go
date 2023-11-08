package main

import (
	"flag"
	"fmt"
	"odo_office/pkg/file_service/bufio_impl"
	"odo_office/pkg/url_service/http_native"
)

func main() {
	var name string
	flag.StringVar(&name, "filename", "", "filename")
	flag.Parse()
	fmt.Println(name)
	fileService, err := bufio_impl.NewFileService(name)
	if err != nil {
		fmt.Println("input err")
	}
	urlService, err := http_native.NewUrlService()
	data, err := fileService.ReadFile()
	if err != nil {
		fmt.Println("file err")
	}
	out, err := urlService.ProcessMany(data)
	if err != nil {
		fmt.Println("http err")
	}
	for _, line := range out {
		fmt.Println(line.ContentLength, line.ProcessTime)
	}
}
