package main

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	pb "gitlab.infra.online.net/paas/carbon/api/grpc/v1"
	"gitlab.infra.online.net/paas/carbon/api/server"
	"gitlab.infra.online.net/paas/carbon/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	util.InitLogging()
	err := util.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	grpcPort := viper.GetString("gateway.backend_port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUsageImpactServer(s, server.NewUsageServer())
	reflection.Register(s)

	log.Infof("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
