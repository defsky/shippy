package main

import (
    "os"
	"log"
    
    "github.com/micro/go-micro"
    _ "github.com/micro/go-plugins/registry/mdns"
    
    pb "shippy/user-service/proto/user"
)
func main(){
    db, err := ConnectDB(os.Getenv("SQL_DRV"), os.Getenv("SQL_DSN"))
    defer db.Close()
    
    if err != nil {
        log.Fatalln("Connect DB error:", err)
    }
    
    repo := &UserRepository{db}
    
    server := micro.NewService(
        micro.Name("go.micro.srv.user"),
        micro.Version("latest"),
    )
    server.Init()

    // 将实现服务端的 API 注册到服务端
    pb.RegisterUserServiceHandler(server.Server(), &handler{repo, &TokenService{repo}})

    if err := server.Run(); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}