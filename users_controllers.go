package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {

	var users Users
	var arr_user []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select * from users")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.Name, &users.Email); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_user = append(arr_user, users)
			// log.Fatal(arr_user)
			// json.NewEncoder(w).Encode(arr_user)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	// log.Fatal(response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func insertUsersMultipart(w http.ResponseWriter, r *http.Request) {
	var response Response
	var users Users

	db := connect()
	defer db.Close()

	json.NewDecoder(r.Body).Decode(&users)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// name := r.Form.Get("name")
	// email := r.Form.Get("email")

	// log.Print(name)
	// log.Print(name)

	_, err = db.Exec("INSERT INTO users (nama, email) values (?,?)",
		users.Name,
		users.Email,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Add"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	var response Response
	var users Users

	db := connect()
	defer db.Close()

	json.NewDecoder(r.Body).Decode(&users)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DELETE FROM users WHERE id = ? ",
		users.Id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Delete"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
