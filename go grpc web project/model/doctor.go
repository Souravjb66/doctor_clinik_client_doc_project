package model

type Doctor struct{
	Id int `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Phoneno string `json:"phoneno" bson:"phoneno"`
	Species string `json:"species" bson:"species"`
	Clinic string `json:"clinic" bson:"clinic"`
}