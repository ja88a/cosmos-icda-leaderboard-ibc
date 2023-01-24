FROM --platform=linux ubuntu:22.04
ARG BUILDARCH

## Tools versions
ENV GO_VERSION=1.18.3
ENV IGNITE_VERSION=0.22.1
ENV NODE_VERSION=18.x

## Local dirs & preps
ENV LOCAL=/usr/local
ENV GOROOT=$LOCAL/go
ENV HOME=/root
ENV GOPATH=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/bin

ENV PACKAGES curl gcc jq make unzip
RUN apt-get update
RUN apt-get install -y $PACKAGES

## Install Go
RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-$BUILDARCH.tar.gz | tar -C $LOCAL -xzf -

## Install Ignite
RUN curl -L https://get.ignite.com/cli@v${IGNITE_VERSION}! | bash

## Install Node
RUN curl -fsSL https://deb.nodesource.com/setup_${NODE_VERSION} | bash -
RUN apt-get install -y nodejs

EXPOSE 1317 3000 4500 5000 26657

ENV WORK_DIR=leaderboard
WORKDIR /${WORK_DIR}

## Integrate the project's Go dependencies
## Uncomment ONLY once the target repo/local working dir has been created (via `ignite scaffold`)
## And you use throwable docker containers (i.e. not a peristent one)
# COPY go.mod /${WORK_DIR}/go.mod
# RUN go mod download
# RUN rm /${WORK_DIR}/go.mod

## Add a new user "john" with user id 8877
# RUN useradd -u 8877 john
## Change to non-root privilege
# USER john