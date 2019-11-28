package main

import (
    "github.com/pkg/errors"
	
	"github.com/globalsign/mgo"

	pb "shippy/vessel-service/proto/vessel"
)

const (
	DB_NAME = "shippy"
	CON_COLLECTION = "vessel"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(v *pb.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(CON_COLLECTION)
}

// 创建货轮
func (repo *VesselRepository) Create(v *pb.Vessel) error {
    return repo.collection().Insert(v)
}

// 查找可用的货轮
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessels []*pb.Vessel

	err := repo.collection().Find(nil).All(&vessels)

	if err != nil {
		return nil, err
	}
	
    // 选择最近一条容量、载重都符合的货轮
    for _, v := range vessels {
        if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
            return v, nil
        }
    }
    return nil, errors.New("No available vessel")
}

func (repo *VesselRepository) Close() {
	// Close() 会在每次查询结束的时候关闭会话
    // Mgo 会在启动的时候生成一个 "主" 会话
    // 你可以使用 Copy() 直接从主会话复制出新会话来执行，即每个查询都会有自己的数据库会话
    // 同时每个会话都有自己连接到数据库的 socket 及错误处理，这么做既安全又高效
    // 如果只使用一个连接到数据库的主 socket 来执行查询，那很多请求处理都会阻塞
    // Mgo 因此能在不使用锁的情况下完美处理并发请求
    // 不过弊端就是，每次查询结束之后，必须确保数据库会话要手动 Close
	// 否则将建立过多无用的连接，白白浪费数据库资源
	
	repo.session.Close()
}