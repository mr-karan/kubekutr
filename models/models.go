package models

// Config represents the structure to hold configuration loaded from an external data source.
type Config struct {
	Workloads []Workload `koanf:"workloads" yaml:"workloads"`
}

// Workload represents the structure to represent all configs and resources to deploy an application.
type Workload struct {
	Name         string        `koanf:"name" yaml:"name"`
	Deployments  []Deployment  `koanf:"deployments" yaml:"deployments"`
	Services     []Service     `koanf:"services" yaml:"services"`
	Ingresses    []Ingress     `koanf:"ingresses" yaml:"ingresses"`
	StatefulSets []StatefulSet `koanf:"statefulsets" yaml:"statefulsets"`
}

// Deployment represents configuration options for the Deployment spec.
type Deployment struct {
	Name       string      `koanf:"name" yaml:"name"`
	Replicas   string      `koanf:"replicas" yaml:"replicas"`
	Containers []Container `koanf:"containers" yaml:"containers"`
	Labels     []Identifer `koanf:"labels" yaml:"labels"`
	Volumes    []Volume    `koanf:"volumes" yaml:"volumes"`
}

// StatefulSet represents configuration options for StatefulSet spec.
type StatefulSet struct {
	Name        string      `koanf:"name" yaml:"name"`
	ServiceName string      `koanf:"serviceName" yaml:"serviceName"`
	Containers  []Container `koanf:"containers" yaml:"containers"`
	Labels      []Identifer `koanf:"labels" yaml:"labels"`
	Volumes     []Volume    `koanf:"volumes" yaml:"volumes"`
}

// Container represents configuration options for the Container spec in a Pod definition.
type Container struct {
	CreateService      bool          `koanf:"createService" yaml:"createService"`
	Name               string        `koanf:"name" yaml:"name"`
	Image              string        `koanf:"image" yaml:"image"`
	EnvSecret          string        `koanf:"envSecret" yaml:"envSecret"`
	EnvVars            []EnvVar      `koanf:"envVars" yaml:"envVars"`
	Container          string        `koanf:"container" yaml:"container"`
	Ports              []Port        `koanf:"ports" yaml:"ports"`
	Command            string        `koanf:"command" yaml:"command"`
	Args               string        `koanf:"args" yaml:"args"`
	VolumeMounts       []VolumeMount `koanf:"volumeMounts" yaml:"volumeMounts"`
	RequestsCPU        string        `koanf:"cpuRequests" yaml:"cpuRequests"`
	RequestsMemory     string        `koanf:"memoryRequests" yaml:"memoryRequests"`
	LimitsCPU          string        `koanf:"cpuLimits" yaml:"cpuLimits"`
	LimitsMemory       string        `koanf:"memoryLimits" yaml:"memoryLimits"`
	ReadinessProbePort string        `koanf:"readinessPort" yaml:"readinessPort"`
	ReadinessProbePath string        `koanf:"readinessPath" yaml:"readinessPath"`
	LivenessProbePort  string        `koanf:"livenessPort" yaml:"livenessPort"`
	LivenessProbePath  string        `koanf:"livenessPath" yaml:"livenessPath"`
}

// Service represents configuration options for Service spec.
type Service struct {
	Name      string      `koanf:"name" yaml:"name"`
	Ports     []Port      `koanf:"ports" yaml:"ports"`
	Type      string      `koanf:"type" yaml:"type"`
	Labels    []Identifer `koanf:"labels" yaml:"labels"`
	Selectors []Identifer `koanf:"selectors" yaml:"selectors"`
	Headless  bool        `koanf:"headless" yaml:"headless"`
}

// Ingress represents configuration options for Ingress spec.
type Ingress struct {
	Name        string        `koanf:"name" yaml:"name"`
	Class       string        `koanf:"class" yaml:"class"`
	Paths       []IngressPath `koanf:"ingressPaths" yaml:"ingressPaths"`
	Annotations []Annotation  `koanf:"annotations" yaml:"annotations"`
	Labels      []Identifer   `koanf:"labels" yaml:"labels"`
}

// IngressPath represents the definition for `paths` specified in Ingress.
type IngressPath struct {
	Path    string `koanf:"path" yaml:"path"`
	Service string `koanf:"service" yaml:"service"`
	Port    string `koanf:"port" yaml:"port"`
}

// Annotation represents the name of annotation value.
type Annotation struct {
	Name string `koanf:"name" yaml:"name"`
}

// Port represents the structure for defining ports in services
type Port struct {
	Name       string `koanf:"name" yaml:"name"`
	Port       string `koanf:"port" yaml:"port"`
	TargetPort string `koanf:"targetPort" yaml:"targetPort"`
	Protocol   string `koanf:"protocol" yaml:"protocol"`
}

// Identifer represents the kv pair for a label.
type Identifer struct {
	Name string `koanf:"name" yaml:"name"`
}

// EnvVar represents the env variables to be used in Pod definition.
type EnvVar struct {
	Name  string `koanf:"name" yaml:"name"`
	Value string `koanf:"value" yaml:"value"`
}

// VolumeMount represents the options for mounting volume in a pod.
type VolumeMount struct {
	MountPath string `koanf:"mountPath" yaml:"mountPath"`
	SubPath   string `koanf:"subPath" yaml:"subPath"`
	Name      string `koanf:"name" yaml:"name"`
}

// Volume represents the option for Volume attached to a pod. Currently only supports
// ConfigMap as the source.
type Volume struct {
	Name string `koanf:"name" yaml:"name"`
}

// Resource is a set of common actions performed on Resource Types.
type Resource interface {
	GetMetaData() ResourceMeta
}

// ResourceMeta contains metadata for preparing resource manifests.
type ResourceMeta struct {
	Name         string
	Config       map[string]interface{}
	TemplatePath string
	Type         string
}
