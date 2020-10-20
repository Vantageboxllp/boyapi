package main

import (
    "fmt"
	"net/http"
	"io/ioutil"
	"log"
)
var dataresponseString string
var catresponseString string

func maindata(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, dataresponseString)
}

func catdata(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, catresponseString)
}

func api(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "loaderio-4d317a2a744ddd08c2729a693e987c51\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
   dataurl := "https://raw.githubusercontent.com/Vantageboxllp/boyapi/master/datas.json"
   caturl := "https://raw.githubusercontent.com/Vantageboxllp/boyapi/master/slider_images.json"

    fmt.Println("fetching video data json from github..")
	response, err := http.Get(dataurl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	dataresponseString = string(responseData)
	fmt.Println("Main data fetched successfully..")

	fmt.Println("fetching Category data json from github..")
	response2, err2 := http.Get(caturl)
	if err != nil {
		log.Fatal(err2)
	}
	defer response2.Body.Close()
	responseData2, err2 := ioutil.ReadAll(response2.Body)
	if err != nil {
		log.Fatal(err2)
	}
	catresponseString = string(responseData2)
	fmt.Println("Cat data fetched successfully..")
	http.HandleFunc("/maindata", maindata)
	http.HandleFunc("/catdata", catdata)
    http.HandleFunc("/loaderio-4d317a2a744ddd08c2729a693e987c51/", api)
    http.HandleFunc("/headers", headers)
    fmt.Println("server running successfully..")
    http.ListenAndServe(":8080", nil)
}
