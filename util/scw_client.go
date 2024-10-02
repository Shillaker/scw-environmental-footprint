package util

import (
	"context"
	"fmt"
	"strings"

	as "github.com/scaleway/scaleway-sdk-go/api/applesilicon/v1alpha1"
	bm "github.com/scaleway/scaleway-sdk-go/api/baremetal/v1"
	ddx "github.com/scaleway/scaleway-sdk-go/api/dedibox/v1"
	instance "github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/shillaker/scw-environmental-footprint/model"
)

type SCWClient struct {
	cli *scw.Client
}

func NewClient(ctx context.Context) (*SCWClient, error) {
	accessKey := viper.GetString("scw.access_key")
	secretKey := viper.GetString("scw.secret_key")

	if len(accessKey) == 0 || len(secretKey) == 0 {
		return nil, fmt.Errorf("Scaleway access key and/or secret key not set")
	}

	cli, err := scw.NewClient(
		scw.WithAuth(accessKey, secretKey),
		scw.WithDefaultRegion(DefaultRegion),
	)

	if err != nil {
		return nil, err
	}

	log.Infof("initialised SCW client with key %s", accessKey)

	client := &SCWClient{
		cli: cli,
	}

	return client, err
}

func (s *SCWClient) ListInstanceServerTypes(ctx context.Context) (map[string]*instance.ServerType, error) {
	api := instance.NewAPI(s.cli)

	offers := map[string]*instance.ServerType{}
	for _, zone := range InstanceZones {
		log.Debugf("requesting instance types in %s", zone)
		resp, err := api.ListServersTypes(&instance.ListServersTypesRequest{
			Zone: zone,
		})

		if err != nil {
			return nil, err
		}

		for instanceName, serverType := range resp.Servers {
			offers[instanceName] = serverType
		}
	}

	return offers, nil
}

func (s *SCWClient) ListInstanceVMs(ctx context.Context) (map[string]model.VirtualMachine, error) {
	serverTypes, err := s.ListInstanceServerTypes(ctx)

	if err != nil {
		return nil, err
	}

	servers := map[string]model.VirtualMachine{}

	for instanceName, serverType := range serverTypes {
		instanceVm := model.VirtualMachine{
			Type:   instanceName,
			SsdGb:  model.MinimumInstanceBlockVolume,
			VCpus:  serverType.Ncpus,
			Gpus:   uint32(*serverType.Gpu),
			RamGiB: uint32(serverType.RAM / 1024 / 1024 / 1024),
		}

		instanceName = strings.ToLower(instanceName)

		if strings.HasPrefix(instanceName, "coparm") {
			instanceVm.Server = model.BaseCopArm1Host
		} else if strings.HasPrefix(instanceName, "dev1") {
			instanceVm.Server = model.BaseDev1Host
		} else if strings.HasPrefix(instanceName, "ent1") {
			instanceVm.Server = model.BaseEnt1Host
		} else if strings.HasPrefix(instanceName, "gp1") {
			instanceVm.Server = model.BaseGp1Host
		} else if strings.HasPrefix(instanceName, "pop2-hm") {
			instanceVm.Server = model.BasePop2HmHost
		} else if strings.HasPrefix(instanceName, "pop2-hc") {
			instanceVm.Server = model.BasePop2HcHost
		} else if strings.HasPrefix(instanceName, "pop2") { // Must come after pop2-hm/hc
			instanceVm.Server = model.BasePop2Host
		} else if strings.HasPrefix(instanceName, "pro2") {
			instanceVm.Server = model.BasePro2Host
		} else if strings.HasPrefix(instanceName, "play2") {
			instanceVm.Server = model.BasePlay2Host
		} else if strings.HasPrefix(instanceName, "ent1") {
			instanceVm.Server = model.BaseEnt1Host
		} else if strings.HasPrefix(instanceName, "stardust1") {
			instanceVm.Server = model.BaseStardust1Host
		} else if strings.HasPrefix(instanceName, "h100") {
			instanceVm.Server = model.BaseH100Host
		} else if strings.HasPrefix(instanceName, "l4") {
			instanceVm.Server = model.BaseL4Host
		} else if strings.HasPrefix(instanceName, "gpu-3070") {
			instanceVm.Server = model.BaseRenderSHost
		} else {
			log.Warnf("Skipping instance, no mapping for type %s", instanceName)
			continue
		}

		servers[instanceName] = instanceVm
	}

	return servers, nil
}

func (s *SCWClient) ListElasticMetalOffers(ctx context.Context) (map[string]*bm.Offer, error) {
	api := bm.NewAPI(s.cli)

	offers := map[string]*bm.Offer{}
	for _, zone := range ElasticMetalZones {
		log.Debugf("requesting em offers in %s", zone)
		resp, err := api.ListOffers(&bm.ListOffersRequest{
			Zone: zone,
		})

		if err != nil {
			return nil, err
		}

		for _, offer := range resp.Offers {
			offers[offer.Name] = offer
		}
	}

	return offers, nil
}

func gbToMb(gbIn scw.Size) uint32 {
	return uint32(gbIn) / 1e3
}

func (s *SCWClient) ListElasticMetalServers(ctx context.Context) (map[string]model.Server, error) {
	allOffers, err := s.ListElasticMetalOffers(ctx)

	if err != nil {
		return nil, err
	}

	servers := map[string]model.Server{}

	for offerName, offer := range allOffers {
		server := model.Server{}

		offerCpu := offer.CPUs[0]
		server.Cpus = append(server.Cpus, model.Cpu{
			Name:        CleanCPUName(offerCpu.Name),
			CoreUnits:   offerCpu.CoreCount,
			Threads:     offerCpu.CoreCount, // Apple M CPUs have threads = cores
			FrequencyHz: offerCpu.Frequency,
			Units:       1,
		})

		for _, disk := range offer.Disks {
			diskType := strings.ToLower(disk.Type)
			if diskType == "nvme" {
				server.Ssds = append(server.Ssds, model.Ssd{
					CapacityMB: gbToMb(disk.Capacity),
					Units:      1,
				})
			} else {
				server.Hdds = append(server.Hdds, model.Hdd{
					CapacityMB: gbToMb(disk.Capacity),
					Units:      1,
				})
			}
		}

		for _, ram := range offer.Memories {
			server.Rams = append(server.Rams, model.Ram{
				CapacityMib: GibiBytesMultipliedByThousandsToMebibytes(ram.Capacity),
				FrequencyHz: ram.Frequency,
				Type:        ram.Type,
				Units:       1,
			})
		}

		server.PowerSupply = model.DefaultPowerSupply(800)
		server.Motherboard.Units = 1

		server.Product = model.ProductElasticMetal
		server.Name = offer.Name

		servers[offerName] = server
	}

	return servers, nil
}

func (s *SCWClient) ListAppleSiliconOffers(ctx context.Context) (map[string]*as.ServerType, error) {
	api := as.NewAPI(s.cli)

	offers := map[string]*as.ServerType{}

	for _, zone := range AppleSiliconZones {
		log.Debugf("requesting as offers in %s", zone)

		resp, err := api.ListServerTypes(&as.ListServerTypesRequest{
			Zone: zone,
		})
		if err != nil {
			return nil, err
		}

		for _, serverType := range resp.ServerTypes {
			offers[serverType.Name] = serverType
		}
	}

	return offers, nil
}

func (s *SCWClient) ListAppleSiliconServers(ctx context.Context) (map[string]model.Server, error) {
	allOffers, err := s.ListAppleSiliconOffers(ctx)
	if err != nil {
		return nil, err
	}

	servers := map[string]model.Server{}

	for offerName, serverType := range allOffers {
		server := model.Server{}

		server.Cpus = append(server.Cpus, model.Cpu{
			Name:        CleanCPUName(serverType.CPU.Name),
			CoreUnits:   serverType.CPU.CoreCount,
			Units:       1,
			FrequencyHz: uint32(serverType.CPU.Frequency),
		})

		ramType := ""
		if serverType.Memory.Type == "LPDDR5" {
			ramType = "ddr5"
		}

		server.Rams = append(server.Rams, model.Ram{
			CapacityMib: GibiBytesMultipliedByThousandsToMebibytes(serverType.Memory.Capacity),
			Units:       1,
			Type:        ramType,
		})

		if strings.ToLower(serverType.Disk.Type) == "ssd" {
			server.Ssds = append(server.Ssds, model.Ssd{
				Units:      1,
				CapacityMB: gbToMb(serverType.Disk.Capacity),
			})
		} else {
			server.Hdds = append(server.Hdds, model.Hdd{
				Units:      1,
				CapacityMB: gbToMb(serverType.Disk.Capacity),
			})
		}

		if serverType.Gpu.Count > 0 {
			server.Gpus = append(server.Gpus, model.Gpu{
				Name:  serverType.CPU.Name, // GPU + CPU all part of same SoC with M Macs
				Units: 1,
			})
		}

		server.PowerSupply = model.DefaultPowerSupply(800)
		server.Motherboard.Units = 1

		server.Product = model.ProductAppleSilicon
		server.Name = serverType.Name

		servers[offerName] = server
	}

	return servers, nil
}

func (s *SCWClient) ListDediboxOffers(ctx context.Context) (map[string]*ddx.Offer, error) {
	api := ddx.NewAPI(s.cli)

	offers := map[string]*ddx.Offer{}

	for _, zone := range DediboxZones {
		log.Debugf("requesting ddx offers in %s", zone)

		resp, err := api.ListOffers(&ddx.ListOffersRequest{
			Zone: zone,
		})

		if err != nil {
			return nil, err
		}

		for _, offer := range resp.Offers {
			offers[offer.Name] = offer
		}
	}

	return offers, nil
}

func (s *SCWClient) ListDediboxServers(ctx context.Context) (map[string]model.Server, error) {
	allOffers, err := s.ListDediboxOffers(ctx)
	if err != nil {
		return nil, err
	}

	res := map[string]model.Server{}

	for offerName, offer := range allOffers {
		server := model.Server{}

		for _, cpu := range offer.ServerInfo.CPUs {
			server.Cpus = append(server.Cpus, model.Cpu{
				Name:        CleanCPUName(cpu.Name),
				CoreUnits:   cpu.CoreCount,
				Threads:     cpu.ThreadCount,
				FrequencyHz: cpu.Frequency,
				Units:       1,
			})
		}

		for _, ram := range offer.ServerInfo.Memories {
			server.Rams = append(server.Rams, model.Ram{
				CapacityMib: uint32(ram.Capacity / 1024),
				Units:       1,
			})
		}

		for _, disk := range offer.ServerInfo.Disks {
			if disk.Type == "sata" || disk.Type == "sas" {
				server.Hdds = append(server.Hdds, model.Hdd{
					CapacityMB: gbToMb(disk.Capacity),
					Units:      1,
				})
			} else {
				server.Ssds = append(server.Ssds, model.Ssd{
					CapacityMB: gbToMb(disk.Capacity),
					Units:      1,
				})
			}
		}

		server.PowerSupply = model.DefaultPowerSupply(800)
		server.Motherboard.Units = 1

		server.Product = model.ProductDedibox
		server.Name = offer.Name

		res[offerName] = server
	}

	return res, nil
}
