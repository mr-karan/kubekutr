package models

// Config represents the structure to hold configuration loaded from an external data source.
type Config struct {
	Deployments []Deployment `koanf:"deployments"`
	Services    []Service    `koanf:"services"`
	Ingresses   []Ingress    `koanf:"ingresses"`
}

// Deployment represents
type Deployment struct {
	Name       string      `koanf:"name"`
	Replicas   string      `koanf:"replicas"`
	Containers []Container `koanf:"containers"`
}

// Container represents
type Container struct {
	Name          string `koanf:"name"`
	Image         string `koanf:"image"`
	EnvSecret     string `koanf:"envSecret"`
	Container     string `koanf:"container"`
	PortInt       string `koanf:"port"`
	Command       string `koanf:"command"`
	Args          string `koanf:"arg"`
	ConfigMapName string `koanf:"configmap"`
}

type Service struct {
	Name       string `koanf:"name"`
	Port       int    `koanf:"port"`
	TargetPort int    `koanf:"targetPort"`
	Type       string `koanf:"type"`
}

type Ingress struct {
	Name  string        `koanf:"name"`
	Class string        `koanf:"class"`
	Paths []IngressPath `koanf:"ingressPaths"`
}

type IngressPath struct {
	Path    string `koanf:"path"`
	Service string `koanf:"service"`
	Port    string `koanf:"port"`
}
