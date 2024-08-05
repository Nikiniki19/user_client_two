package handler

import (
	"context"
	"log"
	"time"
	pb "userclientservice/proto"
)

func FetchUser(grpcclient1 pb.Client2RequestClient, user *pb.Id) (*pb.UserResponse2, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	res, err := grpcclient1.FetchUser(ctx, user)
	if err != nil {
		log.Fatalf("could not create the user: %v", err)
	}
	return res, nil
}

	