package models

// Config represents the structure to hold configuration loaded from an external data source.
type Config struct {
	Workloads []Workload `koanf:"workloads"`
}

// Workload represents the structure to represent all configs and resources to deploy an application.
type Workload struct {
	Name         string        `koanf:"name"`
	Deployments  []Deployment  `koanf:"deployments"`
	Services     []Service     `koanf:"services"`
	Ingresses    []Ingress     `koanf:"ingresses"`
	StatefulSets []StatefulSet `koanf:"statefulsets"`
}

// Deployment represents configuration options for the Deployment spec.
type Deployment struct {
	Name       string      `koanf:"name"`
	Replicas   string      `koanf:"replicas"`
	Containers []Container `koanf:"containers"`
	Labels     []Identifer `koanf:"labels"`
	Volumes    []Volume    `koanf:"volumes"`
}

// StatefulSet represents configuration options for StatefulSet spec.
type StatefulSet struct {
	Name        string      `koanf:"name"`
	ServiceName string      `koanf:"serviceName"`
	Containers  []Container `koanf:"containers"`
	Labels      []Identifer `koanf:"labels"`
	Volumes     []Volume    `koanf:"volumes"`
}

// Container represents configuration options for the Container spec in a Pod definition.
type Container struct {
	CreateService      bool          `koanf:"createService"`
	Name               string        `koanf:"name"`
	Image              string        `koanf:"image"`
	EnvSecret          string        `koanf:"envSecret"`
	EnvVars            []EnvVar      `koanf:"envVars"`
	Container          string        `koanf:"container"`
	Ports              []Port        `koanf:"ports"`
	Command            string        `koanf:"command"`
	Args               string        `koanf:"args"`
	VolumeMounts       []VolumeMount `koanf:"volumeMounts"`
	RequestsCPU        string        `koanf:"cpuRequests"`
	RequestsMemory     string        `koanf:"memoryRequests"`
	LimitsCPU          string        `koanf:"cpuLimits"`
	LimitsMemory       string        `koanf:"memoryLimits"`
	ReadinessProbePort string        `koanf:"readinessProbePort"`
	ReadinessProbePath string        `konaf:"readinessProbePath"`
	LivenessProbePort  string        `koanf:"livenessProbePort"`
	LivenessProbePath  string        `konaf:"livenessProbePath"`
}

// Service represents configuration options for Service spec.
type Service struct {
	Name      string      `koanf:"name"`
	Ports     []Port      `koanf:"ports"`
	Type      string      `koanf:"type"`
	Labels    []Identifer `koanf:"labels"`
	Selectors []Identifer `koanf:"selectors"`
	Headless  bool        `koanf:"headless"`
}

// Ingress represents configuration options for Ingress spec.
type Ingress struct {
	Name        string        `koanf:"name"`
	Class       string        `koanf:"class"`
	Paths       []IngressPath `koanf:"ingressPaths"`
	Annotations []Annotation  `koanf:"annotations"`
	Labels      []Identifer   `koanf:"labels"`
}

// IngressPath represents the definition for `paths` specified in Ingress.
type IngressPath struct {
	Path    string `koanf:"path"`
	Service string `koanf:"service"`
	Port    string `koanf:"port"`
}

// Annotation represents the name of annotation value.
type Annotation struct {
	Name string `koanf:"name"`
}

// Port represents the structure for defining ports in services
type Port struct {
	Name       string `koanf:"name"`
	Port       string `koanf:"port"`
	TargetPort string `koanf:"targetPort"`
	Protocol   string `koanf:"protocol"`
}

// Identifer represents the kv pair for a label.
type Identifer struct {
	Name string `koanf:"name"`
}

// EnvVar represents the env variables to be used in Pod definition.
type EnvVar struct {
	Name  string `koanf:"name"`
	Value string `koanf:"value"`
}

// VolumeMount represents the options for mounting volume in a pod.
type VolumeMount struct {
	MountPath string `koanf:"mountPath"`
	SubPath   string `koanf:"subPath"`
	Name      string `koanf:"name"`
}

// Volume represnts the option for Volume attached to a pod. Currently only supports
// ConfigMap as the source.
type Volume struct {
	Name string `koanf:"name"`
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
