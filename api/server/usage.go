package server

import (
	"context"

	"gitlab.infra.online.net/paas/carbon/api"
	pb "gitlab.infra.online.net/paas/carbon/api/grpc/v1"
	"gitlab.infra.online.net/paas/carbon/impact"
	"gitlab.infra.online.net/paas/carbon/mapping"
	"gitlab.infra.online.net/paas/carbon/model"
)

type UsageServer struct {
	pb.UnimplementedUsageImpactServer
}

func NewUsageServer() *UsageServer {
	return &UsageServer{}
}

func (s *UsageServer) ListInstances(context.Context, *pb.EmptyRequest) (*pb.ListInstancesResponse, error) {
	mapper := mapping.ScwMapper{}
	instances := mapper.ListInstances()

	response := &pb.ListInstancesResponse{}
	for _, instance := range instances {
		response.Instances = append(response.Instances, api.InstanceToPb(instance))
	}

	return response, nil
}

func (s *UsageServer) ListElasticMetal(context.Context, *pb.EmptyRequest) (*pb.ListElasticMetalResponse, error) {
	mapper := mapping.ScwMapper{}
	ems := mapper.ListElasticMetal()

	response := &pb.ListElasticMetalResponse{}
	for _, em := range ems {
		response.ElasticMetals = append(response.ElasticMetals, api.ElasticMetalToPb(em))
	}

	return response, nil
}

func (s *UsageServer) ListKubernetesControlPlanes(context.Context, *pb.EmptyRequest) (*pb.ListKubernetesControlPlanesResponse, error) {
	mapper := mapping.ScwMapper{}
	cps := mapper.ListKubernetesControlPlanes()

	response := &pb.ListKubernetesControlPlanesResponse{}
	for _, cp := range cps {
		response.ControlPlanes = append(response.ControlPlanes, api.KubernetesControlPlaneTypeToPb(cp))
	}

	return response, nil
}

func (s *UsageServer) GetElasticMetalUsageImpact(ctx context.Context, request *pb.ElasticMetalUsageRequest) (*pb.CloudUsageImpactResponse, error) {
	em := api.ElasticMetalPbToModel(request.ElasticMetal)
	usage := api.UsagePbToModel(request.Usage)
	config := api.ImpactConfigPbToModel(request.Config)

	mapper := mapping.ScwMapper{}
	serverUsage, err := mapper.MapElasticMetalUsage(em, usage)

	if err != nil {
		return nil, err
	}

	return doCalculateImpact(config, serverUsage)
}

func (s *UsageServer) GetInstanceUsageImpact(ctx context.Context, request *pb.InstanceUsageRequest) (*pb.CloudUsageImpactResponse, error) {
	instance := api.InstancePbToModel(request.Instance)
	usage := api.UsagePbToModel(request.Usage)
	config := api.ImpactConfigPbToModel(request.Config)

	mapper := mapping.ScwMapper{}
	serverUsage, err := mapper.MapInstanceUsage(instance, usage)

	if err != nil {
		return nil, err
	}

	return doCalculateImpact(config, serverUsage)
}

func (s *UsageServer) GetKubernetesUsageImpact(ctx context.Context, request *pb.KubernetesUsageRequest) (*pb.CloudUsageImpactResponse, error) {
	cpType := api.KubernetesControlPlanePbToModel(request.ControlPlane)

	usage := api.UsagePbToModel(request.Usage)
	config := api.ImpactConfigPbToModel(request.Config)

	poolsPb := request.GetPools()
	pools := make([]model.KubernetesPool, len(poolsPb))
	for i, poolPb := range poolsPb {
		pools[i] = api.KubernetesPoolPbToModel(poolPb)
	}

	mapper := mapping.ScwMapper{}
	serverUsage, err := mapper.MapKubernetesUsage(cpType, pools, usage)

	if err != nil {
		return nil, err
	}

	return doCalculateImpact(config, serverUsage)
}

func doCalculateImpact(config model.ImpactConfig, serverUsage []model.ServerUsage) (*pb.CloudUsageImpactResponse, error) {
	calculator, err := impact.GetCalculator(config)
	if err != nil {
		return nil, err
	}

	res, err := calculator.CalculateServerImpact(serverUsage)
	if err != nil {
		return nil, err
	}

	response := &pb.CloudUsageImpactResponse{}
	response.Impacts = make(map[string]*pb.Impact, len(res.Impacts))

	for impactName, value := range res.Impacts {
		response.Impacts[impactName] = api.ImpactToPb(value)
	}

	response.EquivalentsManufacture = api.EquivalentsToPb(res.EquivalentsManufacture)
	response.EquivalentsUse = api.EquivalentsToPb(res.EquivalentsUse)

	return response, nil
}
