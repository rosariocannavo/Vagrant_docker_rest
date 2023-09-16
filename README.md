# Vagrant Docker REST


This repository contains a simple project that combines Vagrant and Docker to create a development environment for a RESTful API. It helps you quickly set up and run a REST API server within a Vagrant-managed virtual machine using Docker containers.
## Technologies used

<!-- -->
<img src="https://www.vectorlogo.zone/logos/vagrantup/vagrantup-icon.svg" width="90" height="90"> **Vagrant** is an open-source tool and a powerful command-line utility for managing and provisioning virtualized development environments. It is designed to simplify the process of creating, configuring, and sharing reproducible development environments across different platforms and virtualization providers.
<!-- -->
<img src="https://www.vectorlogo.zone/logos/docker/docker-official.svg" width="100" height="100"> **Docker** is a platform and set of tools that simplify the process of developing, shipping, and running applications in lightweight, portable containers. Containers are standalone, executable packages that include everything needed to run a piece of software, including the code, runtime, libraries, and system tools. 
<!-- -->
<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="100" height="130"> **Gin** is a web framework for the Go programming language that simplifies the development of web applications and APIs. It is known for its speed, minimal memory footprint, and ease of use, making it a popular choice for building high-performance web services in Go. 
<!-- -->
<img src="https://www.vectorlogo.zone/logos/expressjs/expressjs-icon.svg" width="100" height="130"> **Express** is a popular and minimalistic web application framework for Node.js, a runtime environment that executes JavaScript on the server side.


## Prerequisites

Before you begin, ensure you have met the following requirements:

- [Vagrant](https://www.vagrantup.com/) installed on your local machine.
- [VirtualBox](https://www.virtualbox.org/) (or another supported provider) installed.
- Basic knowledge of Vagrant and Docker.

## Getting Started

To get started with this project, follow these steps:

1. Clone this repository to your local machine:
   
   `git clone https://github.com/rosariocannavo/Vagrant_docker_rest.git`

3. Run the following command to execute Vagrant:
    `vagrant up` 

4. Run the following to attach a terminal to the VM:
    `vagrant ssh` 

6. Run this command inside the machine to check the containers:
    `docker ps`
  

8. You can check the api root at:
   
    `localhost:8080           //node express`
   
    `localhost:9090/hello     //go gin`

5. To stop the machine run:
    `vagrant halt`

6. To destroy the machine run:
    `vagrant destroy --force`

