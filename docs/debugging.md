---
layout: default
title: Debugging
nav_order: 6
---

# Debugging microservices in Kubernetes

In the previous tutorials, we leveraged the service mesh from a developers' perspective to build resilience and observability into our service calls. This helps us both mitigate issues on the network between services as well as surface the appropriate signals to identify where and when there are issues. Ultimately, the service mesh can help save us from a cascading failure by improving resilience, but we need to couple this with debugging and fixing the issues altogether. 

In this part of the tutorial, we'll explore some mesh-complementary tools to debug a remote running microservice within your IDE and identify problems. We'll also explore how to do this in a more "production" like environment where you may not (most likely!) be able to connect up directly to the service running in production. We'll use a tool called `loop` that will allow us to record and reply requests from an environment. Let's get started.

## Debugging your services with Squash debugger

[Squash debugger](https://squash.solo.io) is a multi-language, distributed debugger that allows you to use your own IDE and other familiar tooling to debug microservices running in Kubernetes. Squash takes care of the plumbing to expose the debugging ports and remote debugger which should allow you to focus on "squashing" bugs instead of fumbling around with Kubernetes and `inode`s and so forth.

![](images/debugging/squashlogo.png)

Squash is available as both a `squashctl` CLI tool as well as a plugin for VSCode. A plugin for IntelliJ/IDEA is also available, though a bit outdated. If there is interest in using this for IDEA, please let us know. For this tutorial, we'll use Squash through the VSCode extensions.

### Pick a pod to debug

From within the VSCode web-based IDE, hit "CTRL+Shift+P" to bring up the plugin dialog and start typing "Squash". You should see the squash plugin pop up. Hit Enter on it:

![](images/debugging/vscode-plugin-search.png)

Under the covers, the VSCode plugin leverages the `squashctl` cli tool, so you can use that alternatively if you wish. Squash will query the current Kubernetes context and navigate you through picking which service/pod you want to debug. Start by selecting the namespace to debug, in our case `default`:

![](images/debugging/squash-select-namespace.png)

Next, select a pod to debug. In this case we are going to debug the payment service, so start typing "pay..." into the search box to filter down the pods.

![](images/debugging/squash-select-pod.png)

Next we need to pick which container to debug. Since we are deployed in a service mesh, we will need to decide between debugging the service-mesh sidecar or the application itself. In our case, we'll select the `payment` container.

![](images/debugging/squash-select-container.png)

Lastly, since Squash is a multi-language debugger, we need to decide which debugger to use. The source code for this service is written in Go, so we'll pick the `dlv` debugger. In future versions of squash, we're hoping to be able to auto-detect the language and default to a debugger. 

![](images/debugging/squash-select-debugger.png)


If this is your first time running Squash, give it a few moments to download the appropriate debugger image and get it started. For example, if you wish to see how it's doing, you can go to the terminal and check the `squash-debugger` namespace:

```shell
kubectl get po -n squash-debugger
NAME         READY   STATUS              RESTARTS   AGE
plankt48h2   0/1     ContainerCreating   0          2m33s
```

If it takes too long for the image to download and connect up with your IDE, delete the pod and try again:

```shell
kubectl delete po -n squash-debugger --all
```

If all goes well, you should be taken to the debug perspective:

![](images/debugging/debug-perspective.png)


### Setting break points

From here you can set a break point on the `payment-service` source code. The code that gets executed on a request is in the `handler.go` source. Place a break point on a location in that code and try exercise the `web` service which ultimately calls into the payment service:

![](images/debugging/set-breakpoint.png)

Now if you exercise your service, you should hit the break point:

![](images/debugging/hit-breakpoint.png)


From here, you're in the VSCode debugger -- there's nothing special about squash here. You can step-by-step debug into the rest of the source code with full access to the stack and variables. You can use this approach to continue debugging for where things might be incorrect in your service. 

## Debugging in production with Loop

continue here....


* Explain why sometimes it is necessary to debug in production
* Show how Loop can be used to capture errored requests and replayed
* Replay errors and debug with Squash