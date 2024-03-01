package database
import(
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
)

type DbInstance struct{
	Db *mongo.Client
	Name string
	First string
	Secend string
	Third string
	Ctx context.Context
	
}
var Mdb DbInstance
func Connect(){
	db,err:=mongo.Connect(context.TODO(),options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil{
		log.Panic(err)
	}

	Mdb.Ctx=context.TODO()
	Mdb.Db=db
	Mdb.Name="vetenary"
	Mdb.First="client"
	Mdb.Secend="doctor"
	Mdb.Third="payment"


}