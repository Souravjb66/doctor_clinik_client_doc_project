package main

import (
	mydb "grpcweb/database"
	proto "grpcweb/mygrpc"
	"log"
	"net"
	// "net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)
type server struct{
	proto.UnimplementedDoctorServer
}
func main(){
	mydb.Connect()
	db:=mydb.Mdb.Db
	defer db.Disconnect(mydb.Mdb.Ctx)
	listen,ln:=net.Listen("tcp",":8080")
	if ln!=nil{
		log.Panic("problem")
	}
	router:=mux.NewRouter()

	srv:=grpc.NewServer()
    proto.RegisterDoctorServer(srv,&server{})
	reflection.Register(srv)
	srv.Serve(listen)


	

}