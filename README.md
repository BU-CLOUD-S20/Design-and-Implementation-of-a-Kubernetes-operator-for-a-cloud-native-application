# Design and Implementation of a Kubernetes operator for a cloud-native application

Team: Lin Ma, Brendan Slabe, Ganghao "Eric" Li, Chenhui "Elvis" Zhu, Yuan Wei   
Mentors:  
Mandana Vaziri ​mvaziri@us.ibm.com​  
Tamar Eilam eilamt@us.ibm.com  

# Project Proposal 

## 1.Vision and Goals Of The Project:
The goal of this project is to build an operator for a cloud native application for developers to deploy microservices on Kubernetes. 

The high level definition of this project is to: 
* Ease the deployment of an application in Kubernetes.
* Extending Kubernetes and it’s CLI’s ability to understand the function, performance and potential of this new resource  


## 2. Users/Personas Of The Project:
The personas we are targeting is an infrastructure operator for entrepreneurs to deploy any applications without knowledge of the dependencies of the application.


## 3. Scope and Features Of The Project:
The scope of the project is to develop an operator (CRD + controller) for a sample application.

Features A:
* Automatically set up running dependencies for sample applications  
* Feedback for custom resource usage 

Feature B:
* Design CRD (which includes the schema)
* Design and implement the CDR controller
* Testing and the controller and refine it 

## 4. Installing:



## 5. Acceptance Criteria:
This section discusses the minimum acceptance criteria at the end of the project and stretch goals.

* Develop an operator for one sample application on Kubernetes 
* Implement test case for the operator 
* Design CDR wich schema 
*  Implement the CDR for the operator and corresponding test cases 
* Testing overall performance on the operator on Kubernetes with multiple different applications 

## 7. Demo PPT
[Demo](https://docs.google.com/presentation/d/1BU5wfGXK9S8Pf8SZImy87siaeywIGy-xiimfU3lXoaY/edit)  


[Demo2](https://docs.google.com/presentation/d/1gmJYY3QYA_xBIzChOqLy_ovGRANse0GVDBQ747lgFTE/edit#slide=id.g70d0994122_1_2956)


[Demo3](https://docs.google.com/presentation/d/10IKEp8qYGuKD5has8A5orP4RJ6Knr88RqPgRaFwaDdI/edit#slide=id.p)


[Demo4](https://docs.google.com/presentation/d/10--GophiCxISNXilWHVruWOu4IiQ06m_fROMDC1EthA/edit#slide=id.g732f1a43a8_0_107) 


[Final](https://docs.google.com/presentation/d/14BEiWadCBx_q0WPgXZw_FoNf3Y3B75TwLAfUIig8MqM/edit#slide=id.g800f20a85e_3_2123)


