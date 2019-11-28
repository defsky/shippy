package main

import (
	"context"
	
	"github.com/globalsign/mgo"

	pb "shippy/vessel-service/proto/vessel"
)

// 定义货船服务
type handler struct {
	session *mgo.Session
}

// vesell-service/handler.go
func (h *handler) GetRepo() Repository {
    return &VesselRepository{h.session.Clone()}
}

// 实现服务端
func (s *handler) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
    // 调用内部方法查找
    v, err := s.GetRepo().FindAvailable(spec)
    if err != nil {
        return err
    }
	resp.Vessel = v
	
    return nil
}

// 实现微服务的服务端
func (h *handler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
    defer h.GetRepo().Close()
    if err := h.GetRepo().Create(req); err != nil {
        return err
    }
    resp.Vessel = req
    resp.Created = true
    return nil
}