package main

import (
    "log"
    "os"
    "errors"
    "context"

    "github.com/micro/go-micro"
    "github.com/micro/go-micro/server"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/metadata"

    pb "shippy/consignment-service/proto/consignment"
    vesselPb "shippy/vessel-service/proto/vessel"
    userPb "shippy/user-service/proto/user"
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
        log.Fatalf("create datastore session error: %v\n", err)
    }
    defer session.Close()

    server := micro.NewService(
        // 必须和 consignment.proto 中的 package 一致
        micro.Name("go.micro.srv.consignment"),
        micro.Version("latest"),
        micro.WrapHandler(AuthWrapper),
    )

    // 解析命令行参数
    server.Init()

    // 作为 vessel-service 的客户端
    vClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())

    // 将 server 作为微服务的服务端
    pb.RegisterShippingServiceHandler(server.Server(), &handler{session, vClient})

    if err := server.Run(); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// AuthWrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user-service 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
    return func(ctx context.Context, req server.Request, resp interface{}) error {
        meta, ok := metadata.FromContext(ctx)
        if !ok {
            return errors.New("no auth meta-data found in request")
        }

        // Note this is now uppercase (not entirely sure why this is...)
        token := meta["Token"]

        if len(token) <= 0 {
            return errors.New("token string can not be null")
        }
        
        // Auth
        authClient := userPb.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
        authResp, err := authClient.ValidateToken(context.Background(), &userPb.Token{
            Token: token,
        })
        log.Println("Auth Resp:", authResp)
        if err != nil {
            return err
        }
        err = fn(ctx, req, resp)
        return err
    }
}