package service

import (
	db "grpcweb/database"
	"grpcweb/model"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func SaveDoctor(doctor *model.Doctor)int{
	connect:=db.Mdb
	collection:=connect.Db.Database(connect.Name).Collection(connect.Secend)
	result,err:=collection.InsertOne(connect.Ctx,doctor)
	if err!=nil{
		log.Println(err)
		return http.StatusBadRequest
	}
	log.Println(result)
	return http.StatusOK

}
func GetDoctors()*[]model.Doctor{
	connect:=db.Mdb
	collection:=connect.Db.Database(connect.Name).Collection(connect.Secend)
	result,err:=collection.Find(connect.Ctx,bson.M{})
	if err!=nil{
		log.Println(err)
		return nil
	}
	defer result.Close(connect.Ctx)
	doctor:=[]model.Doctor{}
	for result.Next(connect.Ctx){
		var doc model.Doctor
		err:=result.Decode(&doc)
		if err!=nil{
			log.Println(err)
			continue
		}
		doctor=append(doctor,doc)
	}
	return &doctor
	
}
func GetOneDoctorByPhoneno(phoneno *string)*model.Doctor{
	connect:=db.Mdb
	filter:=bson.M{"phoneno":*phoneno}
	doctor:=model.Doctor{}
	collection:=connect.Db.Database(connect.Name).Collection(connect.Secend)
	err:=collection.FindOne(connect.Ctx,filter).Decode(&doctor)
	if err!=nil{
		log.Println(err)
		
	}
	return &doctor

}
func GetDoctorBySpecies(species *string)*[]model.Doctor{
	connect:=db.Mdb
	filter:=bson.M{"species":*species}
	collection:=connect.Db.Database(connect.Name).Collection(connect.Secend)
	cursor,err:=collection.Find(connect.Ctx,filter)
	if err!=nil{
		log.Println(err)
	}
	defer cursor.Close(connect.Ctx)
	doctor:=[]model.Doctor{}
	for cursor.Next(connect.Ctx){
		var doc model.Doctor
		err:=cursor.Decode(&doc)
		if err!=nil{
			log.Println(err)
			continue
		}
		doctor=append(doctor,doc)
		

	}
	return &doctor

}
func UpdateDoctorPhoneno(email *string,phoneno *string)int{
	connect:=db.Mdb
	filter:=bson.M{"email":*email}
	update:=bson.M{
		"$set":bson.M{"phoneno":*phoneno}}
	collection:=connect.Db.Database(connect.Name).Collection(connect.Secend)
	result,err:=collection.UpdateOne(connect.Ctx,filter,update)
	if err!=nil{
		log.Println(err)
		return http.StatusBadRequest
	}
	log.Println(result)
	return http.StatusOK
	
}
func UpdateDoctorEmail(phoneno *string,email *string)int{
	connect:=db.Mdb
	filter:=bson.M{"phoneno":*phoneno}
	update:=bson.M{
		"$set":bson.M{"email":*email}}
	collection:=connect.Db.Database(connect.Name).Collection(connect.Secend)
	result,err:=collection.UpdateOne(connect.Ctx,filter,update)
	if err!=nil{
		log.Println(err)
		return http.StatusBadRequest
	}
	log.Println(result)
	return http.StatusOK
}
func UpdateDoctorClinic(phoneno *string,clinic *string)int{
	connect:=db.Mdb
	filter:=bson.M{"phoneno":*phoneno}
	update:=bson.M{
		"$set":bson.M{"clinic":*clinic}}
	collection:=connect.Db.Database(connect.Name).Collection(connect.Secend)
	result,err:=collection.UpdateOne(connect.Ctx,filter,update)
	if err!=nil{
		log.Println(err)
		return http.StatusNotFound
	}
	log.Println(result)
	return http.StatusOK
}