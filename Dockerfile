FROM golang:latest

RUN mkdir /app/
ADD src/ /app/
WORKDIR /app

RUN make
