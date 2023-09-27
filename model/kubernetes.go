package model

import "fmt"

type KubernetesControlPlaneType struct {
	Type        string
	Description string
}

type KubernetesControlPlane struct {
	Instance InstanceBaseServer
	Replicas int32
}

type KubernetesPool struct {
	Instance Instance
	Count    int32
}

// Dedicated control plane offers:
// https://www.scaleway.com/en/docs/containers/kubernetes/reference-content/kubernetes-control-plane-offers/
const (
	KubernetesControlPlaneTypeMutualized  = "mutualized"   // 4GB, 1vCPU, 1 replica
	KubernetesControlPlaneTypeDedicated4  = "dedicated-4"  // 4GB, 2vCPU, 2 replicas
	KubernetesControlPlaneTypeDedicated8  = "dedicated-8"  // 8GB, 2vCPU, 2 replicas
	KubernetesControlPlaneTypeDedicated16 = "dedicated-16" // 16GB, 4vCPU, 2 replicas
)

func KubernetesControlPlaneToString(cp KubernetesControlPlane) string {
	return fmt.Sprintf("%v vCPU, %v GiB, %v replicas", cp.Instance.VCpus, cp.Instance.RamGiB, cp.Replicas)
}

var KubernetesControlPlaneMapping = map[string]KubernetesControlPlane{
	KubernetesControlPlaneTypeMutualized: {
		Instance: InstanceBaseServer{
			VCpus:  1,
			RamGiB: 4,
			Server: BasePlay2Host,
		},
		Replicas: 1,
	},
	KubernetesControlPlaneTypeDedicated4: {
		Instance: InstanceBaseServer{
			VCpus:  2,
			RamGiB: 4,
			Server: BasePlay2Host,
		},
		Replicas: 2,
	},
	KubernetesControlPlaneTypeDedicated8: {
		Instance: InstanceBaseServer{
			VCpus:  2,
			RamGiB: 8,
			Server: BasePlay2Host,
		},
		Replicas: 2,
	},
	KubernetesControlPlaneTypeDedicated16: {
		Instance: InstanceBaseServer{
			VCpus:  4,
			RamGiB: 16,
			Server: BasePlay2Host,
		},
		Replicas: 2,
	},
}
