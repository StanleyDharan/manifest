#!/bin/bash

# Download dependencies
go mod download

# Build the Lambda function
GOOS=linux GOARCH=amd64 go build -o main main.go

# Create a deployment package
zip function.zip main

echo "Build complete. Deployment package created as function.zip" 