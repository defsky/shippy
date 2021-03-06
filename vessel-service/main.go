package main

import (
    "os"
	"log"
    
    "github.com/micro/go-micro"
    
	pb "shippy/vessel-service/proto/vessel"
)

const (
    DEFAULT_HOST = "localhost:27017"
)

func main() {

    // 获取容器设置的数据库地址环境变量的值
    dbHost := os.Getenv("DB_HOST")
    if dbHost == ""{
         dbHost = DEFAULT_HOST
    }

    // 创建于 MongoDB 的主会话，需在退出 main() 时候手动释放连接
    session, err := CreateSession(dbHost)
    
    if err != nil {
        log.Fatalf("create session error: %v\n", err)
    }
    defer session.Close()

    // 停留在港口的货船，先写死
    // vessels := []*pb.Vessel{
    //     {Id: "vessel-01", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
    // }
    // repo := &VesselRepository{vessels}

    server := micro.NewService(
        micro.Name("go.micro.srv.vessel"),
        micro.Version("latest"),
    )
    server.Init()

    // 将实现服务端的 API 注册到服务端
    pb.RegisterVesselServiceHandler(server.Server(), &handler{session})

    if err := server.Run(); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}