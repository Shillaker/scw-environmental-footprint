package api

import (
	pb "gitlab.infra.online.net/paas/carbon/api/grpc/v1"
	"gitlab.infra.online.net/paas/carbon/model"
)

func UsagePbToModel(usage *pb.CloudUsage) model.CloudUsageAmount {
	return model.CloudUsageAmount{
		TimeSeconds:    usage.TimeSeconds,
		Count:          usage.Count,
		LoadPercentage: usage.LoadPercentage,
		MemoryMiB:      usage.MemoryMiB,
		MilliVCPU:      usage.MilliVCPU,
		Region:         usage.Region,
	}
}

func InstancePbToModel(instance *pb.Instance) model.Instance {
	return model.Instance{
		Type:        instance.Type,
		Description: instance.Description,
	}
}

func ElasticMetalPbToModel(em *pb.ElasticMetal) model.ElasticMetal {
	return model.ElasticMetal{
		Type:        em.Type,
		Description: em.Description,
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
		Type: storage.Type,
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
