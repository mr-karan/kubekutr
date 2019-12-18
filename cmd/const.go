package cmd

const (
	BaseDir         = "base"
	DeploymentsDir  = "deployments"
	ServicesDir     = "services"
	IngressesDir    = "ingresses"
	StatefulsetsDir = "statefulsets"
	RBACDir         = "rbac"
)

var (
	// parentPaths = []string{BaseDir}
	subPaths = []string{DeploymentsDir, StatefulsetsDir, ServicesDir, RBACDir, IngressesDir}
)
