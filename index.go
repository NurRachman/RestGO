package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get_users", returnAllUsers).Methods("GET")
	router.HandleFunc("/store_users", insertUsersMultipart).Methods("POST")
	router.HandleFunc("/delete_users", deleteUsers).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
