ARG GO_VERSION

FROM docker.io/golang:${GO_VERSION}-alpine3.17 as builder
RUN apk --no-cache add make bash
WORKDIR /app  
COPY . /app
RUN make service

FROM docker.io/alpine:3.17.2
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/cmd/policycraft/policycraft /usr/bin/policycraft
CMD ["policycraft"]
