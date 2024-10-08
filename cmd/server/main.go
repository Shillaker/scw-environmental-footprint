package main

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	pb "github.com/shillaker/scw-environmental-footprint/api/grpc/v1"
	"github.com/shillaker/scw-environmental-footprint/api/server"
	"github.com/shillaker/scw-environmental-footprint/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := util.InitConfig()
	util.InitLogging()

	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	grpcPort := viper.GetString("gateway.backend_port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	ser, err := server.NewUsageServer()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	pb.RegisterUsageImpactServer(s, ser)
	reflection.Register(s)

	log.Infof("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
