package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	PORT = ":8081"
	PRIV_KEY = "./private_key"
	PUBLIC_KEY = "./public_key"
)

func ScheduleOnce(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ScheduleOnce message received")
	body, err := ioutil.ReadAll(r.Body)
	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	if r.Method == "POST" {
		fmt.Println("Post message received")
		for k, v := range r.Form {
			fmt.Println("key :", k)
			fmt.Println("value :", strings.Join(v, ""))
		}
		f, err := os.OpenFile("sologfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("error opening file: %v", err)
			return
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(string(body))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("HTTP status code returned!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/scheduleonce", ScheduleOnce)
	// http.ListenAndServe(PORT, mux)
	http.ListenAndServeTLS(PORT, PUBLIC_KEY, PRIV_KEY, mux)
}
