package mapping

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"gitlab.infra.online.net/paas/carbon/model"
	"gitlab.infra.online.net/paas/carbon/util"
)

type ScwMapper struct {
}

func (s *ScwMapper) ListInstances() []model.Instance {
	var instances []model.Instance
	for instanceName, server := range model.InstanceServerMapping {
		instances = append(instances, model.Instance{
			Type:        instanceName,
			Description: fmt.Sprintf("%v (%v)", instanceName, model.InstanceToString(server)),
		})
	}

	return instances
}

func (s *ScwMapper) ListElasticMetal() []model.ElasticMetal {
	var ems []model.ElasticMetal
	for emName, server := range model.ElasticMetalServerMapping {
		ems = append(ems, model.ElasticMetal{
			Type:        emName,
			Description: fmt.Sprintf("%v (%v)", emName, model.ServerToString(server)),
		})
	}

	return ems
}

func (s *ScwMapper) ListK8sControlPlanes() []model.K8sControlPlane {
	var cps []model.K8sControlPlane
	for cpName, cpServer := range model.K8sControlPlaneMapping {
		cps = append(cps, model.K8sControlPlane{
			Type:        cpName,
			Description: fmt.Sprintf("%v (%v)", cpName, model.ServerToString(cpServer)),
		})
	}

	return cps
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

	server, exists := model.ElasticMetalServerMapping[em.Type]
	if !exists {
		return nil, util.ErrNoMappingFound
	}

	return s.doServerUsage(server, cloudUsage, 1)
}

func (s *ScwMapper) MapInstanceUsage(instance model.Instance, cloudUsage model.CloudUsageAmount) ([]model.ServerUsage, error) {
	log.Debugf("Calculating usage for instance type %v", instance.Type)

	instanceBase, exists := model.InstanceServerMapping[instance.Type]
	if !exists {
		return nil, util.ErrNoMappingFound
	}

	return s.doServerUsage(instanceBase.Server, cloudUsage, instanceBase.GetHostShare())
}

func (s *ScwMapper) MapStorageUsage(storage model.Storage, cloudUsage model.CloudUsageAmount) ([]model.ServerUsage, error) {
	log.Debugf("Calculating usage for storage type %v", storage.Type)

	// TODO - map storage usage
	var usage []model.ServerUsage

	return usage, nil
}
