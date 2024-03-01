package main

import (
	"log"
    "net/http"
	"context"
	proto "grpcweb/mygrpc"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// "encoding/json"
	
)
type docimp struct{
	strem proto.Doctor_DoctorOnClient
}
var imp docimp
func main(){
	router:=mux.NewRouter()

	connect,err:=grpc.Dial("localhost:8080",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Println(err)
		return
	}

	client:=proto.NewDoctorClient(connect)
	strem,er:=client.DoctorOn(context.TODO())
	if er!=nil{
		log.Println(er)

	}
	imp.strem=strem

	router.HandleFunc("/docin/{live}",Doclive)

	defer connect.Close()

    http.ListenAndServe(":8081",router)

}
func Doclive(w http.ResponseWriter,r *http.Request){
	// var live bool
	// ro1:=mux.Vars{}
	// live=ro1["live"]
	req:=[]*proto.Docavailable{
		{IsTrue: true},{IsTrue: true},{IsTrue: true},{IsTrue: true},
	}
	// json.NewDecoder(r.Body).Decode(&live)

	for _,re:=range req{
		mv:=imp.strem.Send(re)
		if mv!=nil{
			log.Println(mv)
			continue
		}
		
	}
	imp.strem.Recv()
}

