package mapping

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/shillaker/scw-environmental-footprint/data"
	"github.com/shillaker/scw-environmental-footprint/model"
	"github.com/shillaker/scw-environmental-footprint/util"
)

type ScwMapper struct {
	Reader *data.DataReader
}

func (s *ScwMapper) ListInstances() []model.Instance {
	var instances []model.Instance
	for instanceName, server := range s.Reader.InstancesData {
		instances = append(instances, model.Instance{
			Type:        instanceName,
			Description: fmt.Sprintf("%v (%v)", instanceName, model.InstanceToString(server)),
		})
	}

	return instances
}

func (s *ScwMapper) ListElasticMetal() []model.ElasticMetal {
	var ems []model.ElasticMetal
	for emName, server := range s.Reader.ElasticMetalData {
		ems = append(ems, model.ElasticMetal{
			Type:        emName,
			Description: fmt.Sprintf("%v (%v)", emName, model.ServerToString(server)),
		})
	}

	return ems
}

func (s *ScwMapper) ListKubernetesControlPlanes() []model.KubernetesControlPlaneType {
	var cpTypes []model.KubernetesControlPlaneType
	for cpName, cp := range model.KubernetesControlPlaneMapping {
		cpTypes = append(cpTypes, model.KubernetesControlPlaneType{
			Type:        cpName,
			Description: fmt.Sprintf("%v (%v)", cpName, model.KubernetesControlPlaneToString(cp)),
		})
	}

	return cpTypes
}

func (s *ScwMapper) doServerUsage(server model.Server, cloudUsage model.CloudUsageAmount, hostShare float32) ([]model.ServerUsage, error) {
	serverUsageAmount := model.DefaultUsage(cloudUsage.TimeSeconds)
	serverUsageAmount.LoadPercentage = float32(cloudUsage.LoadPercentage)
	serverUsageAmount.Region = cloudUsage.Region

	var usage []model.ServerUsage
	var i int32

	for i = 0; i < cloudUsage.Count; i++ {
		usage = append(usage, model.ServerUsage{
			Server:    server,
			Usage:     serverUsageAmount,
			HostShare: hostShare,
		})
	}

	return usage, nil
}

func (s *ScwMapper) MapElasticMetalUsage(em model.ElasticMetal, cloudUsage model.CloudUsageAmount) ([]model.ServerUsage, error) {
	log.Debugf("Calculating usage for elastic metal type %v", em.Type)

	server, exists := s.Reader.ElasticMetalData[em.Type]
	if !exists {
		return nil, util.ErrNoMappingFound
	}

	return s.doServerUsage(server, cloudUsage, 1)
}

func (s *ScwMapper) MapInstanceUsage(instance model.Instance, cloudUsage model.CloudUsageAmount) ([]model.ServerUsage, error) {
	log.Debugf("Calculating usage for instance type %v", instance.Type)

	instanceBase, exists := s.Reader.InstancesData[instance.Type]
	if !exists {
		return nil, util.ErrNoMappingFound
	}

	return s.doServerUsage(instanceBase.Server, cloudUsage, instanceBase.GetHostShare())
}

func (s *ScwMapper) MapKubernetesUsage(cpType model.KubernetesControlPlaneType, pools []model.KubernetesPool, cloudUsage model.CloudUsageAmount) ([]model.ServerUsage, error) {
	log.Debugf("Calculating usage for K8s cluster %v with %v pools", cpType.Type, len(pools))

	// Get control plane
	cp, exists := model.KubernetesControlPlaneMapping[cpType.Type]
	if !exists {
		return nil, util.ErrNoMappingFound
	}

	// Work out size of result
	resultSize := cp.Replicas
	for _, pool := range pools {
		resultSize += pool.Count
	}

	result := make([]model.ServerUsage, resultSize)

	// Create a usage with the required replicas
	cpUsage := cloudUsage
	cpUsage.Count = cp.Replicas

	// Calculate control plane impact
	cpResult, err := s.doServerUsage(cp.Instance.Server, cpUsage, cp.Instance.GetHostShare())
	if err != nil {
		return nil, err
	}

	resultSlice := result[:0]
	resultSlice = append(resultSlice, cpResult[:]...)

	// Iterate through pools
	for _, pool := range pools {
		poolInstance, exists := s.Reader.InstancesData[pool.Instance.Type]
		if !exists {
			return nil, util.ErrNoMappingFound
		}

		// Create a usage with the pool size
		poolUsage := cloudUsage
		poolUsage.Count = pool.Count

		// Calculate usage for the pool
		poolResult, err := s.doServerUsage(poolInstance.Server, poolUsage, poolInstance.GetHostShare())
		if err != nil {
			return nil, err
		}

		resultSlice = append(resultSlice, poolResult[:]...)
	}

	return result, err
}

func (s *ScwMapper) MapStorageUsage(storage model.Storage, cloudUsage model.CloudUsageAmount) ([]model.ServerUsage, error) {
	log.Debugf("Calculating usage for storage type %v", storage.Type)

	// TODO - map storage usage
	var usage []model.ServerUsage

	return usage, nil
}

func NewScwMapper() (*ScwMapper, error) {
	reader, err := data.NewDataReader()
	if err != nil {
		return nil, err
	}

	return &ScwMapper{
		Reader: reader,
	}, nil
}
