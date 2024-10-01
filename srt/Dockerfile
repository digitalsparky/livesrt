# Stage 1: Build Stage
FROM ubuntu:24.04 AS builder

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    libtool \
    pkg-config \
    autoconf \
    automake \
    libssl-dev \
    git \
    wget \
    libavcodec-dev \
    libavformat-dev \
    libavutil-dev \
    libswresample-dev \
    libavfilter-dev \
    libopencv-dev \
    tclsh \
    golang

# Build SRT
RUN git clone https://github.com/BELABOX/srt.git /opt/srt && \
    cd /opt/srt && \
    chmod +x ./configure && \
    ./configure && \
    make

# Build SRTLA
RUN git clone https://github.com/BELABOX/srtla.git /opt/srtla && \
    cd /opt/srtla && \
    make && \
    ls -laR /opt/srtla

RUN wget -O /tmp/upx.tar.xz https://github.com/upx/upx/releases/download/v4.2.4/upx-4.2.4-amd64_linux.tar.xz && \
    cd /tmp && \
    tar -xvf ./upx.tar.xz --strip 1 && \
    mv upx /usr/local/bin/

RUN wget -O /tmp/noalbs.tgz https://github.com/NOALBS/nginx-obs-automatic-low-bitrate-switching/releases/download/v2.11.2/noalbs-v2.11.2-x86_64-unknown-linux-musl.tar.gz && \
  cd /tmp && \
  tar -zxvf /tmp/noalbs.tgz --strip=1

COPY ./stats /opt/stats

RUN cd /opt/stats && \
    go mod tidy && \
    go build -ldflags="-s -w" -o /opt/stats/stats main.go && \
    upx /opt/stats/stats

# Stage 2: Runtime Stage
FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    libssl-dev \
    libavcodec-dev \
    libavformat-dev \
    libavutil-dev \
    libswresample-dev \
    libavfilter-dev \
    libopencv-dev \
    supervisor

# Copy the built binaries from the builder stage
COPY --from=builder /usr/local/lib/libsrt* /usr/local/lib/
COPY --from=builder /opt/srt/srt-live-transmit /opt/livesrt/bin/
COPY --from=builder /opt/srtla/srtla_rec /opt/livesrt/bin/
COPY --from=builder /opt/srtla/srtla_send /opt/livesrt/bin/
COPY --from=builder /opt/stats/stats /opt/livesrt/bin/
COPY --from=builder /tmp/noalbs /opt/livesrt/bin/
COPY srt.sh /opt/srt.sh
COPY supervisord.conf /etc/supervisor/supervisord.conf

RUN mv /opt/srt.sh /opt/livesrt/bin/ && \
    chmod a+x /opt/livesrt/bin/* && \
    ldconfig

# Expose the default SRT port
EXPOSE 7654/udp
EXPOSE 5002/udp
EXPOSE 8080

# Set the command to run when the container starts
CMD ["/bin/supervisord"]
