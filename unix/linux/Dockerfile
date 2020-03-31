FROM ubuntu:19.04

# Dependencies to get the git sources and go binaries
RUN apt-get update && apt-get install -y  --no-install-recommends \
        ca-certificates \
        curl \
        git \
        rsync \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Get the git sources. If not cached, this takes O(5 minutes).
WORKDIR /git
RUN git config --global advice.detachedHead false
# Linux Kernel: Released 29 Mar 2020
RUN git clone --branch v5.6 --depth 1 https://kernel.googlesource.com/pub/scm/linux/kernel/git/torvalds/linux
# GNU C library: Released 01 Feb 2020 (we should try to get a secure way to clone this)
RUN git clone --branch release/2.31/master --depth 1 git://sourceware.org/git/glibc.git

# Get Go
# TODO: update to Go 1.14 once there is a minor release with https://golang.org/issue/37622 fixed.
ENV GOLANG_VERSION 1.13.9
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 f4ad8180dd0aaf7d7cda7e2b0a2bf27e84131320896d376549a7d849ecf237d7

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
    && echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
    && tar -C /usr/local -xzf golang.tar.gz \
    && rm golang.tar.gz

ENV PATH /usr/local/go/bin:$PATH

# Linux and Glibc build dependencies and emulator
RUN apt-get update && apt-get install -y  --no-install-recommends \
        bison gawk make python3 \
        gcc gcc-multilib \
        gettext texinfo \
        qemu-user \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
# Cross compilers (install recommended packages to get cross libc-dev)
RUN apt-get update && apt-get install -y \
        gcc-aarch64-linux-gnu       gcc-arm-linux-gnueabi     \
        gcc-mips-linux-gnu          gcc-mips64-linux-gnuabi64 \
        gcc-mips64el-linux-gnuabi64 gcc-mipsel-linux-gnu      \
        gcc-powerpc64-linux-gnu     gcc-powerpc64le-linux-gnu \
	gcc-riscv64-linux-gnu                                 \
        gcc-s390x-linux-gnu         gcc-sparc64-linux-gnu     \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Let the scripts know they are in the docker environment
ENV GOLANG_SYS_BUILD docker
WORKDIR /build
ENTRYPOINT ["go", "run", "linux/mkall.go", "/git/linux", "/git/glibc"]
