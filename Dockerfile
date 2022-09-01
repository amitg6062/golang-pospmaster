FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY ./ ./

RUN go mod tidy
RUN go get database/sql

# Build
RUN go build -o /posp_api_go_v2

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 9000

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Run
CMD [ "/posp_api_go_v2" ]
