package main

import (
	"fmt"
	"net/http"
	"os"
)

func Handler_header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "write header! ")
	fmt.Println(r.Header.Get("Accept-Language"))
	w.Header().Set("Accept-Language", r.Header.Get("Accept-Language"))
	fmt.Printf("ClientIP: %s \n", r.RemoteAddr)
	fmt.Printf("VERSION: %s \n", os.Getenv("VERSION"))
}

func Handler_healthz(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "healthz")

	//task1
	fmt.Printf("########################HTTP Req Header##########################")
	for headerKey, headerValues := range r.Header {
		fmt.Println(headerKey)
		fmt.Println(headerValues)
		for _, value := range headerValues {
			w.Header().Add(headerKey, value)
		}
	}

	//task2
	fmt.Printf("########################VERSION##########################\n")
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	fmt.Printf("VERSION: %s \n", os.Getenv("VERSION"))

	//task3
	fmt.Printf("########################ClientIP##########################\n")
	fmt.Printf("ClientIP: %s \n", r.RemoteAddr)

	//task 4
	w.WriteHeader(200)
	fmt.Printf("HTTP return:%d \n", 200)

	fmt.Printf("*************************HTTP Res Header***********************************\n")
	for headerKey, headerValues := range w.Header().Clone() {
		fmt.Println(headerKey)
		fmt.Println(headerValues)
	}

}

func main() {
	http.HandleFunc("/", Handler_header)
	http.HandleFunc("/healthz", Handler_healthz)
	http.ListenAndServe("0.0.0.0:80", nil)
	//fmt.Printf("Web end! \n")
}
