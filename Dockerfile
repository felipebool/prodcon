FROM golang:1.14-alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /go/src/github.com/felipebool/prodcon

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o /go/bin/prodcon cmd/server/main.go

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["prodcon"]