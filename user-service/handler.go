package main

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"unicode/utf8"

	pb "shippy/user-service/proto/user"
)

type handler struct {
	repo         Repository
	tokenService *TokenService
}

func (h *handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	if utf8.RuneCountInString(req.Name) <= 0 {
		return errors.New("user name can not be null")
	}
	if len(req.Email) <= 0 {
		return errors.New("user email can not be null")
	}

	u, err := h.repo.GetByEmail(req)
	if u.Email != "" {
		return errors.New("user email has been used")
	}

	if len(req.Password) <= 6 {
		return errors.New("length of user password can not be less than 6")
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPwd)

	if err := h.repo.Create(req); err != nil {
		return err
	}

	resp.User = req

	return nil
}
func (h *handler) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	u, err := h.repo.Get(req.Id)
	if err != nil {
		return err
	}

	resp.User = u

	return nil
}
func (h *handler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}

	resp.Users = users

	return nil
}
func (h *handler) Auth(ctx context.Context, req *pb.User, resp *pb.Response) error {
	u, err := h.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}

	t, err := h.tokenService.Encode(u)
	if err != nil {
		return err
	}
	//resp.Token = t

	resp.Token = &pb.Token{
		Token: t, //"`x_2nam",
		Valid: true,
	}

	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Response) error {
	// Decode token
	claims, err := h.tokenService.Decode(req.Token)
	if err != nil {
		//return err
		resp.Errors = []*pb.Error{
			&pb.Error{
				Code:        -1,
				Description: err.Error(),
			},
		}

		return err
	}

	log.Println(claims)

	if claims.User.Id == "" {
		err := errors.New("invalid user id")
		resp.Errors = []*pb.Error{
			&pb.Error{
				Code:        -2,
				Description: err.Error(),
			},
		}

		return err
	}

	claims.User.Id = ""
	claims.User.Password = ""
	resp.User = claims.User
	resp.Token = &pb.Token{Valid: true}

	return nil
}
