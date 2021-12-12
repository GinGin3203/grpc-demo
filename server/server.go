package main

import (
	"context"
	pb "github.com/GinGin3203/grpc-demo/proto"
	"github.com/GinGin3203/grpc-demo/server/model"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (server *usersServer) AddUser(ctx context.Context, req *pb.UserCreationRequest) (*pb.UserResponse, error) {
	log.Println(req)
	u, err := server.repo.
		AddUser(ctx,
			req.Role.String(),
			req.Name)

	if err != nil {
		return nil, err
	}
	return userResponse(u), nil

}
func (server *usersServer) DeleteUser(ctx context.Context, uid *pb.UserID) (*pb.UserResponse, error) {
	u, err := server.repo.
		DeleteUser(ctx, int(uid.Id))

	if err != nil {
		return nil, err
	}

	return userResponse(u), nil
}

func (server *usersServer) ChangeUserRole(ctx context.Context, request *pb.UserChangeRoleRequest) (*pb.UserResponse, error) {
	u, err := server.repo.
		ChangeUserRole(ctx,
			int(request.Uid.Id),
			request.NewRole.String())
	if err != nil {
		return nil, err
	}

	return userResponse(u), nil
}

func (server *usersServer) GetUser(ctx context.Context, uid *pb.UserID) (*pb.UserResponse, error) {
	u, err := server.repo.
		GetUser(ctx, int(uid.Id))
	if err != nil {
		return nil, err
	}
	return userResponse(u), nil
}

func (server *usersServer) GetAllUsers(ctx context.Context, _ *emptypb.Empty) (*pb.UserResponseList, error) {
	users, err := server.repo.
		GetAllUsers(ctx)

	if err != nil {
		return nil, err
	}

	return userResponseList(users), nil
}

func userResponse(user *model.UserRecord) *pb.UserResponse {
	return &pb.UserResponse{
		Uid:           &pb.UserID{Id: user.ID},
		Role:          pb.Role(pb.Role_value[user.Role]),
		Name:          user.Name,
		LastUpdatedAt: timestamppb.New(user.LastUpdatedAt),
	}
}

func userResponseList(users []*model.UserRecord) *pb.UserResponseList {
	uList := make([]*pb.UserResponse, 0, len(users))
	for _, u := range users {
		uList = append(uList, userResponse(u))
	}
	return &pb.UserResponseList{Users: uList}
}
