# Download and Run the GO Project Using Docker

Follow the steps below to download the Docker image of the project and run it on your local machine.

## Prerequisites

Make sure you have the following installed on your system:

1. **Docker**: [Install Docker](https://docs.docker.com/get-docker/) if it is not already installed.

2. **Git**: You will need Git to clone the repository. If Git is not installed, you can [install Git here](https://git-scm.com/downloads).

## Steps to Download and Run the Docker Image

Step 1: Clone the Repository

git clone https://github.com/stalindavdvas/ProyectoGo.git

cd ProyectoGo

Step 2: Build the Docker Image

docker build -t proyectogo .

Alternatively, pull it from Docker Hub:

docker pull stalinvasco2024/proyectogo:latest

Step 3: Run the Docker Container

Use Docker Desktop to run the container from the image.

Step 4: Access the Application

Go to:

http://localhost:8080

Step 5: Test the Deployed Project

If you want to test the project without downloading the image, you can visit the link:

https://proyectogo-production.up.railway.app/

The application is deployed on Railway cloud.


