package model

type K8sControlPlane struct {
	Type        string
	Description string
}

var K8sControlPlaneMapping = map[string]Server{}
