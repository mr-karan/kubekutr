package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"zerodha.tech/kubekutr/models"
)

const (
	defaultMemoryLimit    = "256Mi"
	defaultMemoryRequests = "128Mi"
	defaultCPULimit       = "500m"
	defaultCPURequests    = "250m"
	defaultPortName       = "-port"
)

// isInt checks if value is integer
func isInt(val interface{}) error {
	// the reflect value of the result
	value := (val).(string)
	// if the value passed in not int
	_, err := strconv.Atoi(value)
	if err != nil {
		return errors.New("Please enter a number")
	}
	return nil
}

func gatherBasicInfo() int {
	var (
		workloadsLen = 0
	)
	// Gather information from user.
	exitOnInterrupt(survey.AskOne(&survey.Input{
		Message: "How many workloads do you want to deploy?",
		Help:    "Workloads represent different application names.",
		Default: "1",
	}, &workloadsLen, survey.WithValidator(isInt)))
	return workloadsLen
}

func gatherWorkloadsInfo() (models.Workload, error) {
	var (
		wd = models.Workload{}
	)
	exitOnInterrupt(survey.AskOne(&survey.Input{
		Message: "What's the name of the application?",
	}, &wd.Name, survey.WithValidator(survey.Required)))
	// get deployments
	deployments, err := gatherDeploymentsInfo()
	if err != nil {
		return wd, fmt.Errorf("error while preparing resources for deployment: %v", err)
	}
	wd.Deployments = deployments
	return wd, nil
}

func gatherDeploymentsInfo() ([]models.Deployment, error) {
	var deploymentsLen = 0
	var deployments = []models.Deployment{}
	exitOnInterrupt(survey.AskOne(&survey.Input{
		Message: "How many deployments do you want to configure?",
		Help:    "Deployments represent different components of your application.",
		Default: "1",
	}, &deploymentsLen, survey.WithValidator(isInt)))
	for i := 0; i < deploymentsLen; i++ {
		var dep = models.Deployment{}
		var containersLen = 0
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "What's the name of deployment?",
			Help:    "Name of the deployment to be configured.",
			Default: "mydeployment",
		}, &dep.Name, survey.WithValidator(survey.Required)))
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "How many pod replicas do you want to run?",
			Help:    "Specify replica count for the deployment.",
			Default: "1",
		}, &dep.Replicas, survey.WithValidator(isInt)))
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "How many containers do you want to run in one pod?",
			Help:    "Containers form a part of a Pod. Typically each Pod is assosicated with one container.",
			Default: "1",
		}, &containersLen, survey.WithValidator(isInt)))
		for k := 0; k < containersLen; k++ {
			ctr, err := gatherContainerInfo()
			if err != nil {
				return nil, fmt.Errorf("Error while fetching info about containers: %v", err)
			}
			dep.Containers = append(dep.Containers, ctr)
		}
		deployments = append(deployments, dep)
	}
	return deployments, nil
}

func gatherContainerInfo() (models.Container, error) {
	var ctr = models.Container{
		Ports: make([]models.Port, 1),
	}
	exitOnInterrupt(survey.AskOne(&survey.Input{
		Message: "What's the name of the container?",
		Help:    "Name of the container to identify within a pod.",
	}, &ctr.Name, survey.WithValidator(survey.Required)))
	exitOnInterrupt(survey.AskOne(&survey.Input{
		Message: "What's the docker image name?",
	}, &ctr.Image, survey.WithValidator(survey.Required)))
	wantCreateService := true
	exitOnInterrupt(survey.AskOne(&survey.Confirm{
		Message: "Do you want to automatically create a Service?",
		Help:    "Service helps you expose your pods to other pods or external workloads",
		Default: true,
	}, &wantCreateService))
	if wantCreateService {
		ctr.CreateService = wantCreateService
	}
	exitOnInterrupt(survey.AskOne(&survey.Input{
		Message: "Specify the ports to expose from the container.",
		Default: "8000",
	}, &ctr.Ports[0].Port, survey.WithValidator(survey.Required)))
	wantComputeQuotas := false
	exitOnInterrupt(survey.AskOne(&survey.Confirm{
		Message: "Do you want to set resource requests and limits for this container? Skipping this will fall back to defaults which can be edited later.",
		Default: false,
	}, &wantComputeQuotas))
	if wantComputeQuotas {
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "Specify the memory requests.",
			Help:    "Requests describes the minimum amount of compute resources required. For example to put a 800MB memory request, you can specify it as `800Mi`",
			Default: "128Mi",
		}, &ctr.RequestsMemory, survey.WithValidator(survey.Required)))
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "Specify the cpu requests.",
			Help:    "Requests describes the minimum amount of compute resources required. For example to put half a core as requests, you can specify it as `500m`",
			Default: "250m",
		}, &ctr.RequestsCPU, survey.WithValidator(survey.Required)))
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "Specify the memory limits.",
			Help:    "Requests describes the minimum amount of compute resources required. For example to put a 800MB memory limit, you can specify it as `800Mi`",
			Default: "256Mi",
		}, &ctr.LimitsMemory, survey.WithValidator(survey.Required)))
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "Specify the cpu limits.",
			Help:    "Requests describes the minimum amount of compute resources required. For example to put half a core as limit, you can specify it as `500m`",
			Default: "500m",
		}, &ctr.LimitsCPU, survey.WithValidator(survey.Required)))
	} else {
		ctr.RequestsMemory = defaultMemoryRequests
		ctr.RequestsCPU = defaultCPURequests
		ctr.LimitsMemory = defaultMemoryLimit
		ctr.LimitsCPU = defaultCPULimit
	}
	wantReadinessCheck := false
	exitOnInterrupt(survey.AskOne(&survey.Confirm{
		Message: "Do you want to specify any periodic probe for container service readiness.",
		Help:    "Container willContainer will be removed from service endpoints if the probe fails.",
		Default: false,
	}, &wantReadinessCheck))
	if wantReadinessCheck {
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "Specify the HTTP API path where readiness probe is defined in your app.",
			Default: "/healthz",
		}, &ctr.ReadinessProbePath, survey.WithValidator(survey.Required)))
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "Specify the HTTP API port where readiness probe is defined in your app.",
			Default: "8000",
		}, &ctr.ReadinessProbePort, survey.WithValidator(survey.Required)))
	}
	wantLivenessCheck := false
	exitOnInterrupt(survey.AskOne(&survey.Confirm{
		Message: "Do you want to specify any periodic probe for container liveness.",
		Help:    "Container will be restarted if the probe fails.",
		Default: false,
	}, &wantLivenessCheck))
	if wantLivenessCheck {
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "Specify the HTTP API path where liveness probe is defined in your app.",
			Default: "/healthz",
		}, &ctr.LivenessProbePath, survey.WithValidator(survey.Required)))
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "Specify the HTTP API port where liveness probe is defined in your app.",
			Default: "8000",
		}, &ctr.LivenessProbePort, survey.WithValidator(survey.Required)))
	}
	return ctr, nil
}

func gatherOutputFileInfo() string {
	var (
		fileName = "kubekutr.yml"
	)
	// Gather information from user.
	exitOnInterrupt(survey.AskOne(&survey.Input{
		Message: "What should be the config file name?",
		Help:    "Override the default config filename",
		Default: "kubekutr.yml",
	}, &fileName, survey.WithValidator(survey.Required)))
	return fileName
}

func exitOnInterrupt(err error) error {
	if err == terminal.InterruptErr {
		fmt.Println("quitting kubekutr. bye...")
		os.Exit(0)
	}
	return err
}
