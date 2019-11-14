---
layout: default
title: Getting Started
nav_order: 1 
---

# Getting Started

In this section you will learn how to set up your system ready for the workshop.

## Requirements

* Docker - [https://docs.docker.com/install/](https://docs.docker.com/install/)
* Shipyard - [https://shipyard.demo.gs](https://shipyard.demo.gs)

## Clone the example repo

The code repository has source files and examples which will be used by this workshop, before continuing clone this repo.

```shell
git clone https://github.com/hashicorp/consul-service-mesh-for-developers.git
cd consul-service-mesh-for-developers
```

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

Open a new terminal in the IDE ``Ctrl-Shift-` ``

The settings for `kubectl` and `consul` are already configured for you, give this a quick test.

### Getting all running pods `kubectl get pods`

```shell
kubectl get pods
NAME                                                              READY   STATUS    RESTARTS   AGE
consul-consul-connect-injector-webhook-deployment-866c55c88bjh7   1/1     Running   0          45m
consul-consul-server-0                                            1/1     Running   0          45m
consul-consul-fqn7l                                               1/1     Running   0          45m
```

### Display Consul members `consul members`
```shell
consul members
Node                    Address         Status  Type    Build  Protocol  DC   Segment
consul-consul-server-0  10.42.0.9:8301  alive   server  1.6.1  2         dc1  <all>
k3d-shipyard-server     10.42.0.6:8301  alive   client  1.6.1  2         dc1  <default>
```

The folder `1_getting_started` contains a number of Kubernetes configuration files, this will install the demo application and Jaeger for tracing into your Kubernetes cluster.

```shell
kubectl apply -f ./1_getting_started
deployment.apps/api-deployment-v1 created
deployment.apps/payment-deployment-v1 created
service/payment-service created
service/web-service created
deployment.apps/web-deployment created
deployment.apps/jaeger created
service/jaeger-service created
```

You can now create load balancers for the `web-service` and `jaeger-servce` so you can view them in your browser. Run the following commands in your terminal:

```
yard expose --service-name svc/jaeger --port 16686:16686

     _______. __    __   __  .______   ____    ____  ___      .______       _______  
    /       ||  |  |  | |  | |   _  \  \   \  /   / /   \     |   _  \     |       \ 
   |   (----`|  |__|  | |  | |  |_)  |  \   \/   / /  ^  \    |  |_)  |    |  .--.  |
    \   \    |   __   | |  | |   ___/    \_    _/ /  /_\  \   |      /     |  |  |  |
.----)   |   |  |  |  | |  | |  |          |  |  /  _____  \  |  |\  \----.|  .--.  |
|_______/    |__|  |__| |__| | _|          |__| /__/     \__\ | _| `._____||_______/ 


Version: 0.2.11

## Expose service: svc/jaeger ports: 16686:16686 using network null
b920105a8415509ff627e209fb17b9c4385989ba554db75b58cf8ca257a798bd
```

```
yard expose --service-name svc/web --port 9090:9090

     _______. __    __   __  .______   ____    ____  ___      .______       _______  
    /       ||  |  |  | |  | |   _  \  \   \  /   / /   \     |   _  \     |       \ 
   |   (----`|  |__|  | |  | |  |_)  |  \   \/   / /  ^  \    |  |_)  |    |  .--.  |
    \   \    |   __   | |  | |   ___/    \_    _/ /  /_\  \   |      /     |  |  |  |
.----)   |   |  |  |  | |  | |  |          |  |  /  _____  \  |  |\  \----.|  .--.  |
|_______/    |__|  |__| |__| | _|          |__| /__/     \__\ | _| `._____||_______/ 


Version: 0.2.11

## Expose service: svc/web ports: 9090:9090 using network null
6c9582667d2c90beec516ce89339e0e355816ec36b014fd4a736c6b256961a91
```

When you now view the web service in your browser at [http://localhost:9090/ui](http://localhost:9090/ui), you will see the UI for `Fake Service`. Fake Service simulates complex service topologies. In this example, you have two tier system, `Web` calls an upstream service `API`. All of this traffic is flowing over the service mesh.

![](images/getting_started/web.png)

Fake Service is not that fake though, it also emits metrics and tracing data which is capture by `Jaeger`. We will learn more about how tracing works inside your application and in the service mesh in the next section. For now you can look at the dashboard by pointing your browser at [http://localhost:16686/search](http://localhost:16686/search)

![](images/getting_started/jaeger.png)

## Summary

In this section you have learned how to set up a simple application in a development environment. In the next section we will start to investigate the capabilities of the Service Mesh and what they means for us developers.