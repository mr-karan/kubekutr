# kubekutr

<img src="logo.png" alt="drawing" width="400"/>
<!-- ![](logo.png) -->

_Cookie cutter for Kubernetes resource manifests_

`kubekutr` lets you scaffold a [bespoke] configuration for Kubernetes resource manifests with an _opinionated_ GitOps directory structure. `kubekutr` is ideally meant to be used in combination with [kustomize](). 

## Motivation

`kustomize` is a great tool when it comes to declarative application management for manifests. There still exists a lot of manual scaffolding to create a `base` which defines your application state. `kubekutr` aims to solve the issue of writing these manifests manually by providing a very simple Go template rendering engine. 

### Non Goals

`kubekutr` doesn't aim to provide all 1000s of options of templating `yaml` files. More users mean every user will want to customise the `yaml` in some way or the other and this is where `kustomize` comes into picture. Users of `kubekutr` are encourage to use `kustomize` to create _patches_ on top of `bases` to apply any kind of customisation. `kubekutr`'s **only** goal is to create the _base_ directory.

## Usage

```bash
# create a new base

$ kubekutr -c config.toml scaffold -o myapp

# `myapp` is created with the GitOps structure
mydir
`-- base
    |-- deployments
    |   `-- app.yml
    |-- ingresses
    |   `-- app.yml
    |-- services
    |   `-- app.yml
    `-- statefulsets
        `-- app.yml
```

### Configuration

-   **[deployments]**

    -   **name**: Name of the deployment
    -   **replicas**: Represents the number of replicas for a `Pod`
    -   **labels**:
        - **name**: Represent the key value pair as a string. For eg: `"app.kubernetes.io/tier: cache"`
    -   **containers**: List of containers in a Pod
        - **name**: Unique name for a container
        - **image**: Docker image name
        - **portInt**: Number of port to expose from Container
        - **portName**: Human friendly name for a port
        - **command**: Entrypoint array
        - **args**: Arguments to the entrypoint
        - **envVars**: List of environment variables to set in the container
            - **name**: Name of environment variable
            - **value**: Value of environment variable
        - **volumeMounts**: Pod volumes to mount into the container's filesystem
            - **name**: Name of Volume
            - **mountPath**: Path within the container at which the volume should be mounted
            - **subPath**: Path within the volume from which the container's volume should be mounted.
    -   **volumes**: List of volumes defined for a deployment
            - **name**: Name of Volume

-   **[statefulsets]**

    -   **name**: Name of the statefulset
    -   **serviceName**: serviceName is the name of the service that governs this StatefulSet
    -   **labels**: (reference above)
    -   **containers**: (reference above)
    -   **volumes**:(reference above)

-   **[services]**

    -   **name**: Name of service
    -   **type**: Type of service. Can be one of `ClusterIP`, `NodePort`, `LoadBalancer`
    -   **port**: The port that will be exposed by this service
    -   **targetPort**: Number or name of the port to access on the pods targeted by the service
    -   **labels**: (reference above)
    -   **selectors**:
        - **name**:  Route service traffic to pods with label keys and values matching this selector

-   **[ingresses]**

    -   **name**: Name of ingress
    -   **ingressPaths**
        -   **path**: Path which map requests to backends
        -   **service**: Specifies the name of the referenced service
        -   **port**: Specifies the port of the referenced service
    -   **labels**: (reference above)
    -   **annotations**:
        - **name**:  Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata

### Improvements

This is still an alpha release. For a full list of things to improve, see unchecked items in [TODO](TODO.md).
Contributions welcome!
