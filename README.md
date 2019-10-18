# Service Mesh for the Developer Workflow

## Abstract

Service mesh is often presented as a solution for network engineering and system operability, increasing security, reliability, and observability. However, service mesh is also an incredibly useful tool for developers, and understanding how to leverage this technology can dramatically simplify your day to day workflow.

By leveraging free and open-source tools and a scenario-based approach, we will illustrate how a service mesh can help with application resilience, observability, and debugging.

By the end of this workshop you will understand:
How to use metrics and distributed tracing effectively
Reliability patterns like retries, timeouts, and circuit breaking
How to leverage Canary deployments
How you can effectively debug distributed systems

The cloud-native, open-source technology used in this tutorial include:
Envoy
Prometheus
Gloo shot
Consul Service Mesh
Loop
Squash
Open Census

## Requirements

Ideally it will be possible to bundle these tools into a Docker Container to account for for different environments and minimum setup for the user.

* Docker
* Kubectl
* VSCode / Vim
* Squash
* Gloo Shot
* Loop
* Consul (CLI)

## Lesson Plan (total 1hr30 mins)
* What is a Service Mesh (10 mins)
* Setting up the environment (10 mins)
* Deploy your first application (10 mins)
* Investigating metrics and distributed tracing (20 mins)
* Implementing reliability with Service Mesh (20 mins)
  - Retries
  - Timeouts
  - Circuit breaking
* Debugging services (20 mins)
