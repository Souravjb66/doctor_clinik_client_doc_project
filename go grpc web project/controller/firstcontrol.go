package controller

import (
	"encoding/json"
	"grpcweb/model"
	"net/http"
	service2 "grpcweb/doctorclient/service"
	service1 "grpcweb/userclient/service"
)

func SaveUser(w http.ResponseWriter,r *http.Request){
	client:=model.Client{}
	json.NewDecoder(r.Body).Decode(&client)
	result:=service1.SaveClient(&client)
	if result==400{
		json.NewEncoder(w).Encode("request not valid")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("succesful")

}
func SaveDoctor(w http.ResponseWriter,r *http.Request){
	doc:=model.Doctor{}
	json.NewDecoder(r.Body).Decode(&doc)
	result:=service2.SaveDoctor(&doc)
	if result==400{
		json.NewEncoder(w).Encode("bad request")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("successful")

}
func UpdateClientEmail(w http.ResponseWriter,r *http.Request){
	type updateclient struct{
		Newemail string `json:"newemail" bson:"newemail"`
		Phoneno string `json:"phoneno" bson:"phoneno"`
	}
	forupdate:=updateclient{}
	json.NewDecoder(r.Body).Decode(&forupdate)
	result:=service1.UpdateClientEmail(&forupdate.Phoneno,&forupdate.Newemail)
	if result==400{
		json.NewEncoder(w).Encode("bad request")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("successfuly updated")
}
func UpdateClientPassword(w http.ResponseWriter,r *http.Request){
	type updatepassword struct{
		Phoneno string `json:"phoneno" bson:"phoneno"`
		Password string `json:"password" bson:"password"`
	}
	forupdate:=updatepassword{}
	json.NewDecoder(r.Body).Decode(&forupdate)
	result:=service1.UpdateClientPassword(&forupdate.Phoneno,&forupdate.Password)
	if result==400{
		json.NewEncoder(w).Encode("bad request")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("succesfuly updated")
}
func UpdateDoctorClinic(w http.ResponseWriter,r *http.Request){
	type updateclinic struct{
		Clinic string `json:"clinic" bson:"clinic"`
		Phoneno string `json:"phoneno" bson:"phoneno"`
	}
	forupdate:=updateclinic{}
	json.NewDecoder(r.Body).Decode(&forupdate)
	result:=service2.UpdateDoctorClinic(&forupdate.Phoneno,&forupdate.Clinic)
	if result==404{
		json.NewEncoder(w).Encode("bad request 404")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("successfuly clinic updated")
}
func UpdateDoctorPhoneno(w http.ResponseWriter,r *http.Request){
	type updatephoneno struct{
		Newphoneno string `json:"newphoneno" bson:"newphoneno"`
		Email string `json:"email" bson:"email"`
	}
	forupdate:=updatephoneno{}
	json.NewDecoder(r.Body).Decode(&forupdate)
	result:=service2.UpdateDoctorPhoneno(&forupdate.Email,&forupdate.Newphoneno)
	if result==400{
		json.NewEncoder(w).Encode("bad request 400")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("successfuly update phoneno")
}
func UpdateDoctorEmail(w http.ResponseWriter,r *http.Request){
	type updateemail struct{
		Email string `json:"email" bson:"email"`
		Phoneno string `json:"phoneno" bson:"phoneno"`
	}
	forupdate:=updateemail{}
	json.NewDecoder(r.Body).Decode(&forupdate)
	result:=service2.UpdateDoctorEmail(&forupdate.Phoneno,&forupdate.Email)
	if result==400{
		json.NewEncoder(w).Encode("bad request 400")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("successfuly update email")
}