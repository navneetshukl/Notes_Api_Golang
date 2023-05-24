# Stage 1: Build the Go application
FROM golang:1.19 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go mod and sum files
COPY go.mod go.sum ./

# Enable Go modules and set GOPROXY
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

# Download the Go module dependencies
RUN go mod download

# Copy the source code to the container's working directory
COPY ./cmd/main.go ./cmd/
COPY ./database/database.go ./database/
COPY ./routes/routes.go ./routes/
COPY ./models/login.go ./models/
COPY ./models/notes.go ./models/
COPY ./models/signup.go ./models/
COPY ./models/studentnotes.go ./models/
COPY ./models/students.go ./models/

# Build the Go application inside the container
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/

# Stage 2: Create the final image with PostgreSQL
FROM postgres:13

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go application from the previous stage
COPY --from=builder /app/myapp .

# Copy the database scripts or other required files
COPY ./database ./database

# Set environment variables for PostgreSQL configuration
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=mydatabase

# Expose the PostgreSQL default port
EXPOSE 5432

# Run the PostgreSQL server and start the Go application
CMD service postgresql start && ./myapp
