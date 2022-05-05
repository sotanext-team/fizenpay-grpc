package erpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetConnection return GRPC Connection of input url
// It will create connection if it is closed and return connection if connected
func GetConnection(hostURL string, timeout time.Duration) (*grpc.ClientConn, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()
	clientConn, err := grpc.DialContext(
		ctx,
		hostURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	return clientConn, nil
}

// Time2TimePb converts Golang *time.Time to *timestamppb.Timestamp to use in gRPC response
func Time2TimePb(t *time.Time) *timestamppb.Timestamp {
	if t != nil {
		return timestamppb.New(*t)
	}
	return nil
}
