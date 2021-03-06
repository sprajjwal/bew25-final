package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

func getFileName(url string, fileFlag string) string {
	var name string

	if fileFlag != "default_value" {
		name = fileFlag + ".jpg"
	} else {
		split := strings.Split(url, "/")
		name = split[len(split)-1]
	}
	if _, err := os.Stat(name); err == nil {

		id, _ := uuid.NewV4()
		return id.String() + ".jpg"
	}
	return name
}

func main() {

	var urlFlag = flag.String("url", "http://i.imgur.com/m1UIjW1.jpg", "URL to get image from")
	var fileFlag = flag.String("file", "default_value", "Name for the downloaded image")

	flag.Parse()

	response, err := http.Get(*urlFlag)
	if err != nil {
		log.Fatal("err1", err)
	}

	defer response.Body.Close()

	file, err := os.OpenFile(getFileName(*urlFlag, *fileFlag), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal("err2", err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal("err3", err)
	}

	fmt.Println("Downloaded: ", *urlFlag)
}
