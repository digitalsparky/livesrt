# Stage 1: Build Stage
FROM ubuntu:24.04 AS builder

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get upgrade -yq && apt-get install -yq wget tar

RUN wget -O /tmp/noalbs.tgz https://github.com/NOALBS/nginx-obs-automatic-low-bitrate-switching/releases/download/v2.11.2/noalbs-v2.11.2-x86_64-unknown-linux-musl.tar.gz && \
  cd /tmp && \
  tar -zxvf /tmp/noalbs.tgz --strip=1

# Stage 2: Runtime Stage
FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive

COPY --from=builder /tmp/noalbs /opt/noalbs.bin

RUN chmod a+x /opt/noalbs.bin

WORKDIR /opt/noalbs
# Set the command to run when the container starts
CMD ["/opt/noalbs.bin"]
