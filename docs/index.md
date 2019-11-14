---
layout: default
title: Getting Started
nav_order: 1 
---

# Getting Started

## Requirements
* Docker - [https://docs.docker.com/install/](https://docs.docker.com/install/)
* Shipyard - [https://shipyard.demo.gs](https://shipyard.demo.gs)

## Starting Kubernetes and Installing Consul Service Mesh

Once you have installed `Docker` and `Shipyard` you can use `Shipyard` to start a Kubernetes cluster with Consul Service Mesh pre-installed.

Run the following command in your terminal:

```shell
yard up --consul-values $PWD/consul-values.yml

     _______. __    __   __  .______   ____    ____  ___      .______       _______  
    /       ||  |  |  | |  | |   _  \  \   \  /   / /   \     |   _  \     |       \ 
   |   (----`|  |__|  | |  | |  |_)  |  \   \/   / /  ^  \    |  |_)  |    |  .--.  |
    \   \    |   __   | |  | |   ___/    \_    _/ /  /_\  \   |      /     |  |  |  |
.----)   |   |  |  |  | |  | |  |          |  |  /  _____  \  |  |\  \----.|  .--.  |
|_______/    |__|  |__| |__| | _|          |__| /__/     \__\ | _| `._____||_______/ 


Version: 0.2.11

## Creating K8s cluster in Docker and installing Consul

#...

### Setup complete:

To interact with Kubernetes set your KUBECONFIG environment variable
export KUBECONFIG="$HOME/.shipyard/yards/shipyard/kubeconfig.yml"

Consul can be accessed at: http://localhost:8500
Kubernetes dashboard can be accessed at: http://localhost:8443

To expose Kubernetes pods or services use the 'yard expose' command. e.g.
yard expose --service-name svc/myservice --port 8080:8080

When finished use "yard down" to cleanup and remove resources
```

You should will not be able to see the Consul UI at [http://localhost:8500](http://localhost:8500), and the Kubernetes dashboard at [http://localhost:8443](http://localhost:8443)

### Consul UI
![](images/getting_started/consul_ui.png)

### Kubernetes Dashboard
![](images/getting_started/k8s_ui.png)

## Development Environment
Shipyard comes bundled with a built in development environment, you can of course use your own IDE and terminal if you have the tools installed but for the purposes of this workshop we are going to be using the built in tools like Go, KubeCtl, Consul, etc.

To start the built in IDE run the following command in your terminal:

``` shell
yard vscode

     _______. __    __   __  .______   ____    ____  ___      .______       _______  
    /       ||  |  |  | |  | |   _  \  \   \  /   / /   \     |   _  \     |       \ 
   |   (----`|  |__|  | |  | |  |_)  |  \   \/   / /  ^  \    |  |_)  |    |  .--.  |
    \   \    |   __   | |  | |   ___/    \_    _/ /  /_\  \   |      /     |  |  |  |
.----)   |   |  |  |  | |  | |  |          |  |  /  _____  \  |  |\  \----.|  .--.  |
|_______/    |__|  |__| |__| | _|          |__| /__/     \__\ | _| `._____||_______/ 


Version: 0.2.11

## Starting VSCode in a browser
Starting VS Code

When finished you can stop the VSCode server using the command:
docker kill vscode

c294f9f3d42c17ff9135ec3bd17e8951a4c5086290b51bbfe6e418e3dfca14ed
```

![](images/getting_started/vscode.png)

## Running the demo application
Now that Kubernetes and Consul are running you can install the example application.