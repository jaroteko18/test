package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ExerciseBasic struct {
	DTime time.Time
	// GetDTime string
}

var Profile struct {
	ID       bson.ObjectId
	Name     string
	Birthday string
	Age      int
	Parent   []string
}

var File struct {
	ID      int
	Fname   string
	LINK    string
	IsError bool
	Message string
	UD      string
	UB      string
}

var User struct {

	// DATE string `json:"date,omitempty"`

	LOGIN      string
	NAME       string
	COMPANY    string
	CREATED_AT string
	UPDATED_AT string
}

// // Person a
// var ExerciseBasic struct {
// 	ID        bson.ObjectId
// 	Nama      string
// 	Birthdate string
// 	Age       int
// 	Parent    []string
// 	// NAMA string `json:"name,omitempty"`
// 	// UMUR int    `json:"umur,omitempty"`
// }
