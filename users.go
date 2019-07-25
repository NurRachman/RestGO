package main

type Users struct {
	Id    int    `form:"id" json:"id"`
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
}

type Response struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Users `json:"data"`
}
