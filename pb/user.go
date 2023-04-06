package pb

import (
	"context"
	"log"
)

type UserService struct {
}

var times int

func (u UserService) Search(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	log.Printf("grpc request [%d]", times)
	defer func() {
		times++
	}()
	username := "李四"
	return &UserResponse{
		Id:       100,
		Username: &username,
		Area:     []string{"1003", "2048"},
	}, nil
}

func (u UserService) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}
