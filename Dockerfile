# Stage 1: Build the Go application
FROM golang:1.19 AS builder

## create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app

# Set the working directory inside the container
#

ADD . /app

WORKDIR /app

RUN go build -o main .
EXPOSE 8080

CMD ["/app/main"]


