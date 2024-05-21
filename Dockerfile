FROM golang:1.22.3-bookworm AS builder
WORKDIR /app
RUN git clone https://github.com/fujiwara/aws-sdk-client-go.git .
RUN rm -f gen.yaml
ENTRYPOINT ["./build-in-docker.sh"]
