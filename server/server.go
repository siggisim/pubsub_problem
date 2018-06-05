package main

import (
	//"context"
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	//gc "cloud.google.com/go/pubsub"

	pb "api/service"
)

const (
	infinity = time.Duration(math.MaxInt64)
)

var (
	port = flag.Int("port", 8000, "The server port")
)

type serviceServer struct {
	//client *gc.Client
}

func NewServiceServer() *serviceServer {
	// ctx := context.Background()
	//
	// client, err := gc.NewClient(ctx, "test")
	// if err != nil {
	// 	return nil
	// }

	//return &serviceServer{client}
	return &serviceServer{}
}

func (s *serviceServer) TestRPC(ctx context.Context, registerRequest *pb.TestRPCRequest) (*pb.TestRPCReply, error) {

	return &pb.TestRPCReply{}, nil
}

func main() {
	flag.Parse()

	var opts []grpc.ServerOption

	keepAliveOpt := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     infinity,
		MaxConnectionAge:      infinity,
		MaxConnectionAgeGrace: infinity,
		Time:    25 * time.Second,
		Timeout: 5 * time.Second,
	})

	keepAliveEnforcementPolicyOpt := grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		MinTime:             5 * time.Minute,
		PermitWithoutStream: false,
	})

	opts = []grpc.ServerOption{keepAliveOpt, keepAliveEnforcementPolicyOpt}

	grpcServer := grpc.NewServer(opts...)

	var s pb.ServiceServer = NewServiceServer()
	pb.RegisterServiceServer(grpcServer, s)

	addr := fmt.Sprintf(":%d", *port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server Starting...")

	grpcServer.Serve(lis)
}
