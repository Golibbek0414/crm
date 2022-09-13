package grpc

import (
	"context"
	"github.com/Golibbek0414/crmprotos/schedulepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

)
func NewScheduleServiceClient(ctx context.Context, url string) (schedulepb.NewScheduleServiceClient, error){
	conn, err:=grpc.DialContext(
		ctx, url,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err!=nil{
		return nil, err
	}
	return schedulepb.NewScheduleServiceClient(conn),nil
}