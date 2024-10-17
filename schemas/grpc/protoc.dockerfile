FROM ubuntu:22.04 AS base

# Target architecture is set by docker automatically
ARG TARGETARCH
ARG NODE_VERSION="22"
ARG PROTOC_VERSION="28.2"
ARG PROTBUF_TS_PLUGIN_VERSION="2.9.4"
ARG DEBIAN_FRONTEND="noninteractive"

RUN set -eu \
    && apt -y update \
    && apt -y install --no-install-recommends \
    ca-certificates \
    curl \
    sudo \
    unzip

# Install node
RUN curl -fsSL https://deb.nodesource.com/setup_${NODE_VERSION}.x | sudo -E bash - && \
    sudo apt-get install --no-install-recommends -y nodejs

# Install protoc and typescript plugin
RUN if [ "$TARGETARCH" = "amd64" ]; then ARCH="x86_64"; elif [ "$TARGETARCH" = "arm64" ]; then ARCH="aarch_64"; fi \
    && curl -L "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-${ARCH}.zip" -o protoc.zip \
    && sudo unzip protoc.zip -d /usr/local \
    && chmod +x /usr/local/bin/protoc \
    && sudo npm install --global @protobuf-ts/plugin@${PROTOBUF_TS_PLUGIN_VERSION}

CMD ["bash"]
