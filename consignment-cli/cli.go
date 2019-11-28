package main

import (
    "io/ioutil"
    "encoding/json"
    "errors"
    "log"
    "os"
    "context"

    microclient "github.com/micro/go-micro/client"
	//"github.com/micro/go-micro/cmd"
	//"github.com/micro/go-micro/metadata"

    // 导如 protoc 自动生成的包
    pb "shippy/consignment-service/proto/consignment"
)

const (
    DEFAULT_INFO_FILE = "consignment.json"
)

// 读取 consignment.json 中记录的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        return nil, err
    }
    var consignment *pb.Consignment
    err = json.Unmarshal(data, &consignment)
    if err != nil {
        return nil, errors.New("consignment.json file content error")
    }
    return consignment, nil
}

func main() {
    //cmd.Init()
    // 创建微服务的客户端，简化了手动 Dial 连接服务端的步骤
    client := pb.NewShippingServiceClient("ship.consignment", microclient.DefaultClient)

    // 在命令行中指定新的货物信息 json 文件
    infoFile := DEFAULT_INFO_FILE
    if len(os.Args) > 1 {
        infoFile = os.Args[1]
    }

    // 解析货物信息
    consignment, err := parseFile(infoFile)
    if err != nil {
        log.Fatalf("parse info file error: %v", err)
    }

    log.Printf("Consignment will be created: %+v", consignment)

    // 调用 RPC
    // 将货物存储到我们自己的仓库里
    resp, err := client.CreateConsignment(context.Background(), consignment)
    if err != nil {
        log.Fatalf("create consignment error: %v", err)
    }

    // 新货物是否托运成功
    log.Printf("created: %t", resp.Created)

    log.Printf("All created consignments:")
    resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
    if err != nil {
        log.Fatalf("get consignments error: %v", err)
    }

    log.Printf("Total count(%d):\n",len(resp.Consignments))
    for _, c := range resp.Consignments {
        log.Printf("%+v", c)
    }
}