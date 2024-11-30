package api

import (
	pb "github.com/shillaker/scw-environmental-footprint/api/grpc/v1"
	"github.com/shillaker/scw-environmental-footprint/model"
)

func ImpactConfigPbToModel(pbConfig *pb.ImpactConfig) model.ImpactConfig {
	return model.ImpactConfig{
		Backend: pbConfig.GetBackend(),
	}
}

func UsagePbToModel(usage *pb.CloudUsage) model.CloudUsageAmount {
	return model.CloudUsageAmount{
		TimeSeconds:    usage.GetTimeSeconds(),
		Count:          usage.GetCount(),
		LoadPercentage: usage.GetLoadPercentage(),
		MemoryMiB:      usage.GetMemoryMiB(),
		MilliVCPU:      usage.GetMilliVCPU(),
		Region:         usage.GetRegion(),
	}
}

func InstancePbToModel(instance *pb.Instance) model.Instance {
	return model.Instance{
		Type:        instance.GetType(),
		Description: instance.GetDescription(),
	}
}

func ElasticMetalPbToModel(em *pb.ElasticMetal) model.ElasticMetal {
	return model.ElasticMetal{
		Type:        em.GetType(),
		Description: em.GetDescription(),
	}
}

func KubernetesControlPlanePbToModel(cp *pb.KubernetesControlPlane) model.KubernetesControlPlaneType {
	return model.KubernetesControlPlaneType{
		Type:        cp.GetType(),
		Description: cp.GetDescription(),
	}
}

func KubernetesPoolPbToModel(cp *pb.KubernetesPool) model.KubernetesPool {
	return model.KubernetesPool{
		Instance: InstancePbToModel(cp.GetInstance()),
		Count:    cp.GetCount(),
	}
}

func InstanceToPb(instance model.Instance) *pb.Instance {
	return &pb.Instance{
		Type:        instance.Type,
		Description: instance.Description,
	}
}

func ElasticMetalToPb(em model.ElasticMetal) *pb.ElasticMetal {
	return &pb.ElasticMetal{
		Type:        em.Type,
		Description: em.Description,
	}
}

func KubernetesControlPlaneTypeToPb(cp model.KubernetesControlPlaneType) *pb.KubernetesControlPlane {
	return &pb.KubernetesControlPlane{
		Type:        cp.Type,
		Description: cp.Description,
	}
}

func StoragePbToModel(storage *pb.Storage) model.Storage {
	return model.Storage{
		Type: storage.GetType(),
	}
}

func ImpactToPb(impact model.Impact) *pb.Impact {
	return &pb.Impact{
		Use:         impact.Use,
		Manufacture: impact.Manufacture,
		Unit:        impact.Unit,
	}
}

func EquivalentToPb(equivalent model.EquivalentCO2E) *pb.EquivalentCO2E {
	return &pb.EquivalentCO2E{
		Thing:  equivalent.Thing,
		Amount: equivalent.Amount,
	}
}

func EquivalentsToPb(equivalents []model.EquivalentCO2E) []*pb.EquivalentCO2E {
	var pbEquivalents []*pb.EquivalentCO2E

	for _, equivalent := range equivalents {
		pbEquivalents = append(pbEquivalents, EquivalentToPb(equivalent))
	}

	return pbEquivalents
}
