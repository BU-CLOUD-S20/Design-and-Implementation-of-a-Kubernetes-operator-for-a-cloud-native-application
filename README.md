# Design and Implementation of a Kubernetes operator for a cloud-native application

Team: Lin Ma, Brendan Slabe, Ganghao "Eric" Li, Chenhui "Elvis" Zhu, Yuan Wei
Mentors:
* Mandana Vaziri ​mvaziri@us.ibm.com​
* Tamar Eilam eilamt@us.ibm.com


## Project mentors: 

* Mandana Vaziri ​mvaziri@us.ibm.com​
* Tamar Eilam eilamt@us.ibm.com

## Fundamental technologies 

* Go
* Kubernetes

## Project Overview:

In this project, students will select a sample cloud-native application (consisting of at least 2microservices). They can either select from samples that exist, or build a small application oftheir own, or modify an application that already exists. They will then build an operator(Kubernetes Custom Resource and controller) to manage the life cycle of this application as awhole when deployed on Kubernetes. This requires designing the schema for the operator, andimplementing a controller that given instances of this schema knows how todeploy/undeploy/update all the various components of the application. Since controllersemploy optimistic concurrency control, special attention needs to be given to correctness andtesting

## User Stories

