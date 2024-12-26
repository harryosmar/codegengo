package main

import (
	"context"
	"flag"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/harryosmar/codegen-go/errors"
	gw "github.com/harryosmar/codegen-go/grpc/pb/example" // Update
	"github.com/harryosmar/codegen-go/middlewares"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"net"
	"net/http"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	go func() {
		grpcMetrics := grpc_prometheus.NewServerMetrics()
		grpcServer := grpc.NewServer(
			grpc.UnaryInterceptor(
				grpc_middleware.ChainUnaryServer(
					middlewares.LogMiddleware(),
				),
			),
			grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		)
		defer grpcServer.GracefulStop()

		gw.RegisterMyServiceServer(grpcServer, gw.NewMyService())
		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(k string) (string, bool) {
			if k == middlewares.HeaderRequestId {
				return k, true
			}
			return runtime.DefaultHeaderMatcher(k)
		}),
		runtime.WithErrorHandler(errors.ErrorHandlerFunc),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = gw.RegisterMyServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	gw.ForwardResponseMessage()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
