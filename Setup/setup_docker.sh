#update packages
sudo apt-get update && sudo apt-get upgrade
#install docker
curl -sSL https://get.docker.com | sh
#add permission to current user to run docker command
sudo usermod -aG docker pi
sudo reboot
#check version
docker --version
#check if installed properly
docker run armhf/hello-world

#Launch App through docker
#create local folder
mkdir dk_test
cd dk_test
touch Dockerfile app.py requirements.txt
#fill Dockerfile with
#==================
# Use an official Python runtime as a parent image
FROM python:2.7-slim

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed packages specified in requirements.txt
RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Make port 80 available to the world outside this container
EXPOSE 80

# Define environment variable
ENV NAME World

# Run app.py when the container launches
CMD ["python", "app.py"]
#===========
#fill requirements with
#==========
Flask
Redis
#=========
#fill app.py with
#=========
from flask import Flask
from redis import Redis, RedisError
import os
import socket

# Connect to Redis
redis = Redis(host="redis", db=0, socket_connect_timeout=2, socket_timeout=2)

app = Flask(__name__)

@app.route("/")
def hello():
    try:
        visits = redis.incr("counter")
    except RedisError:
        visits = "<i>cannot connect to Redis, counter disabled</i>"

    html = "<h3>Hello {name}!</h3>" \
           "<b>Hostname:</b> {hostname}<br/>" \
           "<b>Visits:</b> {visits}"
    return html.format(name=os.getenv("NAME", "world"), hostname=socket.gethostname(), visits=visits)

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=80)
#=========
#build docker image
docker build --tag=friendlyhello:v0.0.1 .
#run the app
docker run -p 4000:80 friendlyhello:v0.0.1
#check the contents
curl http://localhost:4000
#run the app in background
docker run -d -p 4000:80 friendlyhello
#list all images
docker container ls
#stop image with container ID
docker container stop 1fa4ab2cf395
#remove image
docker image rm -f docker_test:v1

#share the image
docker login
docker tag friendlyhello kevin92sdocker/get-started:part2 
docker push kevin92sdocker/get-started:part2
#pull and run image
docker pull kevin92sdocker/get-started
docker run -p 4000:80 kevin92sdocker/get-started:part2