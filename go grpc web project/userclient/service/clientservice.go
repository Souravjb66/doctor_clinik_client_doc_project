package service

import (
	db "grpcweb/database"
	"grpcweb/model"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	
)

func SaveClient(client *model.Client)int{
	connect:=db.Mdb
	collection:=connect.Db.Database(connect.Name).Collection(connect.First)
	result,err:=collection.InsertOne(connect.Ctx,client)
	if err!=nil{
		log.Println(err)
		return http.StatusBadRequest
	}
	log.Println(result)
	return http.StatusOK
}
func GetClients()*[]model.Client{
	connect:=db.Mdb
	collection:=connect.Db.Database(connect.Name).Collection(connect.First)
	result,err:=collection.Find(connect.Ctx,bson.M{})
	
	if err!=nil{
		log.Println(err)
		return nil

	}
	defer result.Close(connect.Ctx)
	var client []model.Client
	for result.Next(connect.Ctx){
		var nclient model.Client
		err:=result.Decode(&nclient)
		if err!=nil{
			log.Println(err)
			continue
		}
		client=append(client,nclient)

	}
	return &client

}
func GetOneClient(phoneno *string)*model.Client{
	connect:=db.Mdb
	filter:=bson.M{"phoneno":*phoneno}
	collection:=connect.Db.Database(connect.Name).Collection(connect.First)
	client:=model.Client{}
	result:=collection.FindOne(connect.Ctx,filter).Decode(&client)
	log.Println(result)
	return &client
	
}
func UpdateClientPassword(phoneno *string,password *string)int{
	connect:=db.Mdb
	filter:=bson.M{"phoneno":*phoneno}
	update:=bson.M{
		"$set":bson.M{"password":*password}}
	collection:=connect.Db.Database(connect.Name).Collection(connect.First)
	result,err:=collection.UpdateOne(connect.Ctx,filter,update)
	if err!=nil{
		log.Println(err)
		return http.StatusBadRequest

	}
	log.Println(result)
	return http.StatusOK



}
func UpdateClientEmail(phoneno *string,email *string)int{
	connect:=db.Mdb
	filter:=bson.M{"phoneno":*phoneno}
	update:=bson.M{
		"$set":bson.M{"email":*email}}
	collection:=connect.Db.Database(connect.Name).Collection(connect.First)
	result,err:=collection.UpdateOne(connect.Ctx,filter,update)
	if err!=nil{
		log.Println(err)
		return http.StatusBadRequest
	}
	log.Println(result)
	return http.StatusOK

}
