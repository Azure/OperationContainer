FROM mcr.microsoft.com/openjdk/jdk:17-ubuntu


# Install required packages
RUN apt-get update && apt-get upgrade -y && apt-get install -y \
    wget \
    default-jre \
    python3-pip \
    npm 

#Download both platform's versions of golang
RUN wget https://golang.org/dl/go1.23.1.linux-arm64.tar.gz && \
    tar -C /usr/local -xzf go1.23.1.linux-arm64.tar.gz && \
    mv /usr/local/go /usr/local/go-arm64

RUN wget https://golang.org/dl/go1.23.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz && \
    mv /usr/local/go /usr/local/go-amd64

WORKDIR /app

ARG TARGETPLATFORM

#Copy over version of go depending on platform
RUN if [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
    cp -r /usr/local/go-amd64 /usr/local/go; \
    else \
    cp -r /usr/local/go-arm64 /usr/local/go; \
    fi && \
    rm -rf /usr/local/go-amd64 && \
    rm -rf /usr/local/go-arm64

#Set required environment variables
ENV GO111MODULE=on
ENV GOBIN=/usr/local/go/bin
ENV GOPATH=/usr/local/go
ENV PATH=$PATH:/usr/local/go/bin

#Install necessary go packages.
RUN go install github.com/bufbuild/buf/cmd/buf@v1.29.0
RUN go install go.uber.org/mock/mockgen@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Install npm packages
RUN npm install -g swagger-nodegen-cli

# Start a shell session (optional)
CMD ["/bin/bash"]
