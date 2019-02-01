#PUB SUB Demo

The purpose of this project is to demonstrate ability to create an application with a Developer Operations mindset.

The scope of this project will include modifying a simple publish subscribe microservice written in golang, a client interface written in golang, incorporating the elastic stack, utilizing single sign on to create private interactions, containerization via Docker, and application packaging using Kubernetes/Helm.

I intend to have a chart consist of a variable number of pub-sub applications and apm servers, a single client interface and apm server, a redis cache to act as a buffer between packetbeat agents and logstash instances, an nginx proxy with a packetbeat agent, a variable number of logstash instances, a single elasticsearch container, a single kibana container, and a vault instance.

