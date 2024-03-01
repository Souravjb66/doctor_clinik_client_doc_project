package model

type Client struct{
	Id int `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Phoneno string `json:"phoneno" bson:"phoneno,omitempty,min=10"`
	Address string `json:"address" bson:"address,omitempty"`

}