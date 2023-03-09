FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
RUN /bin/bash build.sh
