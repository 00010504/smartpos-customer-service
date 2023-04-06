FROM golang:1.16 as builder

#
RUN mkdir -p $GOPATH/src/gitlab.7i.uz/invan/invan_customer_service 
WORKDIR $GOPATH/src/gitlab.7i.uz/invan/invan_customer_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
  export GOOS=linux && \
  go mod vendor && \
  make build && \
  mv ./bin/invan_customer_service /

FROM alpine
COPY --from=builder invan_customer_service .
ENTRYPOINT ["/invan_customer_service"]
