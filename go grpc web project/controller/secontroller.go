package controller

import (
	service2 "grpcweb/doctorclient/service"
	"net/http"
	"encoding/json"
	service1 "grpcweb/userclient/service"

)

func GetAllClient(w http.ResponseWriter,r *http.Request){
	result:=service1.GetClients()
	if *result==nil{
		json.NewEncoder(w).Encode("empty client")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*result)

}
func GetOneClient(w http.ResponseWriter,r http.Request){
	var phoneno string
	json.NewDecoder(r.Body).Decode(&phoneno)
	result:=service1.GetOneClient(&phoneno)
	if result==nil{
		json.NewEncoder(w).Encode("empty data")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*result)
}
func GetAllDoctor(w http.ResponseWriter,r *http.Request){
	result:=service2.GetDoctors()
	if *result==nil{
		json.NewEncoder(w).Encode("empty doctor")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*result)
}
func GetDoctorByPhone(w http.ResponseWriter,r http.Request){
	var phoneno string
	json.NewDecoder(r.Body).Decode(&phoneno)
	result:=service2.GetOneDoctorByPhoneno(&phoneno)
	if result==nil{
		json.NewEncoder(w).Encode("not found")
		return

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*result)

}
func GetDoctorBySpecies(w http.ResponseWriter,r *http.Request){
	var species string
	json.NewDecoder(r.Body).Decode(&species)
	result:=service2.GetDoctorBySpecies(&species)
	if result==nil{
		json.NewEncoder(w).Encode("not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*result)
}
