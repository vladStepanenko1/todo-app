FROM golang:alpine as builder

# Move to working directory
WORKDIR /go/src/todo-app

# Copy the source code into the container
COPY . .

# Download dependecy using go.mod
RUN go mod download

# Build the application
RUN go build -o ./bin/todo-app ./cmd/main.go

# Build a small image
FROM alpine:3.13
WORKDIR /usr/bin

# Copy binary file from builder image to runtime image 
COPY --from=builder /go/src/todo-app/bin /go/bin

# Copy config file to runtime image
COPY --from=builder /go/src/todo-app/configs /usr/bin/configs

# Export neccesary port
EXPOSE 8000

# Command to run when starting the container
CMD [ "/go/bin/todo-app" ]