package main

type Student struct { //
	Id        string `form:"id" json:"id"` //
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
}

type Response struct { // base response
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Student
}