package client

import (
	pb "github.com/Azure/OperationContainer/api/v1"
	"github.com/Azure/aks-middleware/interceptor"

	log "log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewClient returns a client that has all the interceptors registered.
func NewClient(remoteAddr string, options interceptor.ClientInterceptorLogOptions) (pb.OperationContainerClient, error) {
	conn, err := grpc.Dial(
		remoteAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			interceptor.DefaultClientInterceptors(options)...,
		),
	)
	if err != nil {
		// logging for transparency, error handled by retry interceptor
		log.Error("did not connect: " + err.Error())
	}

	return pb.NewOperationContainerClient(conn), err
	// TODO: Figure out how to close the connection when the program exits.
	// defer conn.Close()
}
