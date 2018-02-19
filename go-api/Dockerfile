FROM golang:1.10 AS build

ENV APP_PATH /go/src/go-api
# Set working directory.
WORKDIR $APP_PATH
# Install dep.
RUN go get -u github.com/golang/dep/cmd/dep
# Copy Gopkg files for dep.
COPY ./Gopkg.toml ./Gopkg.lock $APP_PATH/
# Install dependencies with dep.
RUN dep ensure -vendor-only
# Add src.
ADD . $APP_PATH
# Build the source.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/main


# # Use a minimal alpine image.
# FROM alpine:3.7
# # Add ca-certificates.
# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
# # Set working directory.
# WORKDIR /root
# # Copy the binary from build.
# COPY --from=build /go/src/go-api/main .
# # Run the binary.
# CMD ["./main"]


# Use a minimal alpine image.
FROM scratch
# Copy the binary from build.
COPY --from=build /go/src/go-api/bin/main bin/main
# Run the binary.
CMD ["./bin/main"]
