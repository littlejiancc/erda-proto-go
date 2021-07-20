// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: expression.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/expression/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// ExpressionService expression.proto
	ExpressionService() pb.ExpressionServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		expressionService: pb.NewExpressionServiceClient(cc),
	}
}

type serviceClients struct {
	expressionService pb.ExpressionServiceClient
}

func (c *serviceClients) ExpressionService() pb.ExpressionServiceClient {
	return c.expressionService
}

type expressionServiceWrapper struct {
	client pb.ExpressionServiceClient
	opts   []grpc1.CallOption
}

func (s *expressionServiceWrapper) GetExpression(ctx context.Context, req *pb.GetExpressionRequest) (*pb.GetExpressionResponse, error) {
	return s.client.GetExpression(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
