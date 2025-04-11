FROM golang:1.24.2-bookworm AS builder
WORKDIR /app
RUN git clone https://github.com/fujiwara/awslim.git .
RUN rm -f gen.yaml

ENTRYPOINT ["./build-in-docker.sh"]
