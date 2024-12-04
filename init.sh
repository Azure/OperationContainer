#!/bin/bash -x

# For the api.
cd api
go mod init github.com/Azure/OperationContainer/api
go mod edit -require github.com/Azure/aks-middleware@v0.0.23
go get google.golang.org/genproto@latest
cd ..
cd api
cd v1 
mkdir mock 
make service
if [ $? -ne 0 ]
then
    echo "Make service failed."
    exit 1
fi
cd ..
go build ./...
if [ $? -ne 0 ]
then
    echo "Building the api module failed."
    exit 1
fi
go test ./...
if [ $? -ne 0 ]
then
    echo "Testing the api module failed."
    exit 1
fi
gofmt -s -w .
cd ..
