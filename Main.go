package main

import (
	"log"
	"net/http"
	"task-back-end-cyberarmyid/config"
	"task-back-end-cyberarmyid/controllers/homecontroller"
)

func main() {
	config.ConnectDB()

	//Home
	http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/input", homecontroller.Input)
	http.HandleFunc("/listKelas", homecontroller.ListKelas)
	http.HandleFunc("/detail", homecontroller.Detail)

	log.Println("mantap")
	http.ListenAndServe(":80", nil)
}
