package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/shillaker/scw-environmental-footprint/api/grpc/v1"
	"github.com/shillaker/scw-environmental-footprint/util"
)

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	backendServerHost := viper.GetString("gateway.backend_host")
	backendServerPort := viper.GetString("gateway.backend_port")
	gatewayServerPort := viper.GetString("gateway.port")

	grpcServerEndpoint := fmt.Sprintf("%v:%v", backendServerHost, backendServerPort)
	log.Infof("registering gRPC server at %v", grpcServerEndpoint)

	// Register usage backend
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterUsageImpactHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	gatewayListenAddress := fmt.Sprintf(":%v", gatewayServerPort)
	log.Infof("gateway listening at %v", gatewayListenAddress)

	// Start proxy server
	srv := &http.Server{
		Addr:    gatewayListenAddress,
		Handler: cors(mux),
	}
	return srv.ListenAndServe()
}

func main() {
	util.InitLogging()
	err := util.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	if err = run(); err != nil {
		log.Fatal(err)
	}
}
