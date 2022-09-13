package grpc

import (
	"context"
	"github.com/Golibbek0414/crmprotos/journalpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
func NewJournalServiceClient(ctx context.Context, url string) (journalpb.JournalServiceClient, error) {
	conn, err := grpc.DialContext(
		ctx, url, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return journalpb.NewJournalServiceClient(conn), nil
}

