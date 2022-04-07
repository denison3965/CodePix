package main

import (
	"os"

	"github.com/denison3965/CodePix/application/grpc"
	"github.com/denison3965/CodePix/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)

}
